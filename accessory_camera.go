package homekit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/alpr777/homekit/ffmpeg"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/rtp"
	"github.com/brutella/hc/service"
	"github.com/brutella/hc/tlv8"
	"github.com/nfnt/resize"
	"github.com/radovskyb/watcher"
)

//AccessoryCamera provides RTP video streaming.
type AccessoryCamera struct {
	*accessory.Accessory
	Control           *service.CameraControl
	StreamManagement1 *service.CameraRTPStreamManagement
	StreamManagement2 *service.CameraRTPStreamManagement
}

//NewAccessoryCamera returns an IP camera accessory.
func NewAccessoryCamera(info accessory.Info, args ...interface{}) *AccessoryCamera {
	acc := AccessoryCamera{}
	acc.Accessory = accessory.New(info, accessory.TypeIPCamera)
	acc.Control = service.NewCameraControl()
	acc.AddService(acc.Control.Service)

	// TODO (mah) a camera must support at least 2 rtp streams
	acc.StreamManagement1 = service.NewCameraRTPStreamManagement()
	acc.StreamManagement2 = service.NewCameraRTPStreamManagement()
	acc.AddService(acc.StreamManagement1.Service)
	// acc.AddService(acc.StreamManagement2.Service)

	return &acc
}

//SetupFFMPEGStreaming configures a camera to use ffmpeg to stream video.
//The returned handle can be used to interact with the camera (start, stop, take snapshot…).
func (cam *AccessoryCamera) SetupFFMPEGStreaming(cfg ffmpeg.Config) ffmpeg.FFMPEG {
	ff := ffmpeg.New(cfg)

	setupStreamManagement(cam.StreamManagement1, ff, cfg.MultiStream)
	setupStreamManagement(cam.StreamManagement2, ff, cfg.MultiStream)

	return ff
}

func first(ips []net.IP, filter func(net.IP) bool) net.IP {
	for _, ip := range ips {
		if filter(ip) == true {
			return ip
		}
	}
	return nil
}

func setupStreamManagement(m *service.CameraRTPStreamManagement, ff ffmpeg.FFMPEG, multiStream bool) {
	status := rtp.StreamingStatus{Status: rtp.StreamingStatusAvailable}
	setTLV8Payload(m.StreamingStatus.Bytes, status)
	setTLV8Payload(m.SupportedRTPConfiguration.Bytes, rtp.NewConfiguration(rtp.CryptoSuite_AES_CM_128_HMAC_SHA1_80))
	setTLV8Payload(m.SupportedVideoStreamConfiguration.Bytes, rtp.DefaultVideoStreamConfiguration())
	setTLV8Payload(m.SupportedAudioStreamConfiguration.Bytes, rtp.DefaultAudioStreamConfiguration())

	m.SelectedRTPStreamConfiguration.OnValueRemoteUpdate(func(buf []byte) {
		var cfg rtp.StreamConfiguration
		err := tlv8.Unmarshal(buf, &cfg)
		if err != nil {
			log.Debug.Fatalf("SelectedRTPStreamConfiguration: Could not unmarshal tlv8 data: %s\n", err)
		}

		log.Debug.Printf("%+v\n", cfg)

		id := ffmpeg.StreamID(cfg.Command.Identifier)
		switch cfg.Command.Type {
		case rtp.SessionControlCommandTypeEnd:
			ff.Stop(id)

			if ff.ActiveStreams() == 0 {
				// Update stream status when no streams are currently active
				setTLV8Payload(m.StreamingStatus.Bytes, rtp.StreamingStatus{Status: rtp.StreamingStatusAvailable})
			}

		case rtp.SessionControlCommandTypeStart:
			ff.Start(id, cfg.Video, cfg.Audio)

			if multiStream == false {
				// If only one video stream is suppported, set the status to busy.
				// This way HomeKit knows that nobody is allowed to connect anymore.
				// If multiple streams are supported, the status is always availabe.
				setTLV8Payload(m.StreamingStatus.Bytes, rtp.StreamingStatus{Status: rtp.StreamingStatusBusy})
			}
		case rtp.SessionControlCommandTypeSuspend:
			ff.Suspend(id)
		case rtp.SessionControlCommandTypeResume:
			ff.Resume(id)
		case rtp.SessionControlCommandTypeReconfigure:
			ff.Reconfigure(id, cfg.Video, cfg.Audio)
		default:
			log.Debug.Printf("Unknown command type %d", cfg.Command.Type)
		}
	})

	m.SetupEndpoints.OnValueUpdateFromConn(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		buf := m.SetupEndpoints.GetValue()
		var req rtp.SetupEndpoints
		err := tlv8.Unmarshal(buf, &req)
		if err != nil {
			log.Debug.Fatalf("SetupEndpoints: Could not unmarshal tlv8 data: %s\n", err)
		}

		log.Debug.Printf("%+v\n", req)

		iface, err := ifaceOfConnection(conn)
		if err != nil {
			log.Debug.Println(err)
			return
		}
		ip, err := ipAtInterface(*iface, req.ControllerAddr.IPVersion)
		if err != nil {
			log.Debug.Println(err)
			return
		}

		// TODO ssrc is different for every stream
		ssrcVideo := int32(1)
		ssrcAudio := int32(2)

		resp := rtp.SetupEndpointsResponse{
			SessionId: req.SessionId,
			Status:    rtp.SessionStatusSuccess,
			AccessoryAddr: rtp.Addr{
				IPVersion:    req.ControllerAddr.IPVersion,
				IPAddr:       ip.String(),
				VideoRtpPort: req.ControllerAddr.VideoRtpPort,
				AudioRtpPort: req.ControllerAddr.AudioRtpPort,
			},
			Video:     req.Video,
			Audio:     req.Audio,
			SsrcVideo: ssrcVideo,
			SsrcAudio: ssrcAudio,
		}

		ff.PrepareNewStream(req, resp)

		log.Debug.Printf("%+v\n", resp)

		// After a write, the characteristic should contain a response
		setTLV8Payload(m.SetupEndpoints.Bytes, resp)
	})
}

// ipAtInterface returns the ip at iface with a specific version.
// version is either `rtp.IPAddrVersionv4` or `rtp.IPAddrVersionv6`.
func ipAtInterface(iface net.Interface, version uint8) (net.IP, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		log.Debug.Println(err)
		return nil, err
	}

	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			log.Debug.Println(err)
			continue
		}

		switch version {
		case rtp.IPAddrVersionv4:
			if ip.To4() != nil {
				return ip, nil
			}
		case rtp.IPAddrVersionv6:
			if ip.To16() != nil {
				return ip, nil
			}
		default:
			break
		}
	}

	return nil, fmt.Errorf("%s: No ip address found for version %d", iface.Name, version)
}

// ifaceOfConnection returns the network interface at which the connection was established.
func ifaceOfConnection(conn net.Conn) (*net.Interface, error) {
	host, _, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(host)
	// 2019-06-04 (mah) ip might be nil if `host` contains the network interface name
	// I couldn't find any documentation why v6 ip address contains the interface name
	if ip == nil {
		// get the interface name from the host string
		// ex. host = "fe80::e627:bec4:30b9:cb12%wlan0"
		comps := strings.Split(host, "%")
		if len(comps) == 2 {
			name := comps[1]
			log.Debug.Printf("querying interface with name %s\n", name)
			return net.InterfaceByName(name)
		}

		return nil, fmt.Errorf("unable to parse ip %s", host)
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			addrIP, _, err := net.ParseCIDR(addr.String())
			if err != nil {
				return nil, err
			}

			if reflect.DeepEqual(addrIP, ip) {
				return &iface, nil
			}
		}
	}

	return nil, fmt.Errorf("Could not find interface for connection")
}

func setTLV8Payload(c *characteristic.Bytes, v interface{}) {
	if tlv8, err := tlv8.Marshal(v); err == nil {
		c.SetValue(tlv8)
	} else {
		log.Debug.Fatal(err)
	}
}

const (
	// TypeAssets is the uuid of the Assets characteristic
	TypeAssets        = "ACD9DFE7-948D-43D0-A205-D2F6F368541D"
	TypeTakeSnapshot  = "E8AEE54F-6E4B-46D8-85B2-FECE188FDB08"
	TypeCameraControl = "19BDAD9E-6102-48D5-B413-3F11253706AE"

	// TypeDeleteAssets is the uuid of the DeleteAssets characteristic
	TypeDeleteAssets = "3982EB69-1ECE-463E-96C6-E5A7DF2FA1CD"

	TypeGetAsset = "6A6C39F5-67F0-4BE1-BA9D-E56BD27C9606"
)

// RefDate represents the reference date used to generate asset ids.
// Short ids are prefered and therefore we use 1st April 2019 as the reference date.
var RefDate = time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)

// Assets contains a list of assets encoded as JSON.
// A valid JSON looks like this. `{"assets":[{"id":"1.jpg", "date":"2019-04-01T10:00:00+00:00"}]}`
// Writing to this characteristic is discouraged.
type Assets struct {
	*characteristic.Bytes
}

type AssetsMetadataResponse struct {
	Assets []CameraAssetMetadata `json:"assets"`
}

type CameraAssetMetadata struct {
	ID   string `json:"id"`
	Date string `json:"date"`
}

func NewAssets() *Assets {
	b := characteristic.NewBytes(TypeAssets)
	b.Perms = []string{characteristic.PermRead, characteristic.PermEvents}

	b.Value = []byte{}

	return &Assets{b}
}

type CameraControl struct {
	TakeSnapshot *TakeSnapshot
	Assets       *Assets
	GetAsset     *GetAsset
	DeleteAssets *DeleteAssets

	CameraSnapshotReq func(width, height uint) (*image.Image, error)

	snapshots []*snapshot
	w         *watcher.Watcher
}

func NewCameraControl() *CameraControl {
	cc := CameraControl{}

	cc.TakeSnapshot = NewTakeSnapshot()
	cc.Assets = NewAssets()
	cc.GetAsset = NewGetAsset()
	cc.DeleteAssets = NewDeleteAssets()

	return &cc
}

func (cc *CameraControl) SetupWithDir(dir string) {
	r := regexp.MustCompile(".*\\.jpg")

	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Info.Println(err)
	}

	for _, f := range fs {
		if r.MatchString(f.Name()) == false {
			continue
		}
		path := filepath.Join(dir, f.Name())
		b, err := ioutil.ReadFile(path)
		if err != nil {
			log.Info.Println(f, err)
		} else {
			s := snapshot{
				ID:    f.Name(),
				Date:  f.ModTime().Format(time.RFC3339),
				Bytes: b,
				Path:  path,
			}
			cc.add(&s)
		}
	}
	cc.updateAssetsCharacteristic()

	go cc.watch(dir, r)

	cc.GetAsset.OnValueRemoteUpdate(func(buf []byte) {
		var req GetAssetRequest
		err := json.Unmarshal(buf, &req)
		if err != nil {
			log.Debug.Fatalln("GetAssetRequest:", err)
		}

		for _, s := range cc.snapshots {
			if s.ID == req.ID {
				r := bytes.NewReader(s.Bytes)
				img, err := jpeg.Decode(r)
				if err != nil {
					log.Info.Printf("jpeg.Decode() %v", err)
					cc.GetAsset.SetValue([]byte{})
					return
				}

				scaled := resize.Resize(req.Width, req.Height, img, resize.Lanczos3)
				imgBuf := new(bytes.Buffer)
				if err := jpeg.Encode(imgBuf, scaled, nil); err != nil {
					log.Info.Printf("jpeg.Encode() %v", err)
					cc.GetAsset.SetValue([]byte{})
					return
				}

				cc.GetAsset.SetValue(imgBuf.Bytes())
				return
			}
		}
	})

	cc.DeleteAssets.OnValueRemoteUpdate(func(buf []byte) {
		var req DeleteAssetsRequest
		err := json.Unmarshal(buf, &req)
		if err != nil {
			log.Debug.Fatalln("GetAssetRequest:", err)
			return
		}

		for _, id := range req.IDs {
			err = cc.deleteWithID(id)
			if err != nil {
				log.Debug.Println("delete:", err)
			}
		}
	})

	cc.TakeSnapshot.OnValueRemoteUpdate(func(v bool) {
		if v == true {
			img, err := cc.CameraSnapshotReq(1920, 1080)
			if err != nil {
				log.Info.Println(err)
			} else {
				name := fmt.Sprintf("%.0f.jpg", time.Now().Sub(RefDate).Seconds())
				path := filepath.Join(dir, name)

				buf := new(bytes.Buffer)
				if err := jpeg.Encode(buf, *img, nil); err != nil {
					log.Debug.Printf("jpeg.Encode() %v", err)
				} else {
					ioutil.WriteFile(path, buf.Bytes(), os.ModePerm)
				}
			}

			// Disable shutter after some timeout
			go func() {
				<-time.After(1 * time.Second)
				cc.TakeSnapshot.SetValue(false)
			}()
		}
	})
}

func (cc *CameraControl) add(s *snapshot) {
	log.Debug.Println("add:", s.ID)
	cc.snapshots = append(cc.snapshots, s)
}

func (cc *CameraControl) deleteWithID(id string) error {
	log.Debug.Println("del:", id)
	for _, s := range cc.snapshots {
		if s.ID == id {
			return os.Remove(s.Path)
		}
	}

	return fmt.Errorf("File with id %s not found", id)
}

func (cc *CameraControl) removeWithID(id string) {
	log.Debug.Println("rmv:", id)
	for i, s := range cc.snapshots {
		if s.ID == id {
			cc.snapshots = append(cc.snapshots[:i], cc.snapshots[i+1:]...)
			return
		}
	}
}

func (cc *CameraControl) updateAssetsCharacteristic() {
	assets := []CameraAssetMetadata{}
	for _, s := range cc.snapshots {
		asset := CameraAssetMetadata{
			ID:   s.ID,
			Date: s.Date,
		}
		assets = append(assets, asset)
	}

	p := AssetsMetadataResponse{
		Assets: assets,
	}
	if b, err := json.Marshal(p); err != nil {
		log.Info.Println(err)
	} else {
		log.Debug.Println(string(b))
		cc.Assets.SetValue(b)
	}
}

func (cc *CameraControl) watch(dir string, r *regexp.Regexp) {
	w := watcher.New()
	w.FilterOps(watcher.Create, watcher.Remove)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				switch event.Op {
				case watcher.Create:
					b, err := ioutil.ReadFile(event.Path)
					if err != nil {
						log.Info.Println(event.Path, err)
					} else {
						s := snapshot{
							ID:    event.Name(),
							Date:  event.ModTime().Format(time.RFC3339),
							Bytes: b,
							Path:  event.Path,
						}
						cc.add(&s)
					}
				case watcher.Remove:
					cc.removeWithID(event.Name())
				default:
					break
				}

				cc.updateAssetsCharacteristic()

			case err := <-w.Error:
				log.Info.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(dir); err != nil {
		log.Info.Fatalln(err)
	}

	if err := w.Start(time.Second * 1); err != nil {
		log.Info.Fatalln(err)
	}
}

type snapshot struct {
	ID    string
	Date  string
	Bytes []byte
	Path  string
}

// DeleteAssets is used to handle request to delete assets.
// A valid JSON looks like this. `{"ids":["1.jpg"]}`
// Reading the value of this characteristic is discouraged.
type DeleteAssets struct {
	*characteristic.Bytes
}

func NewDeleteAssets() *DeleteAssets {
	b := characteristic.NewBytes(TypeDeleteAssets)
	b.Perms = []string{characteristic.PermRead, characteristic.PermWrite}
	b.Value = []byte{}

	return &DeleteAssets{b}
}

type DeleteAssetsRequest struct {
	IDs []string `json:"ids"`
}

// GetAsset is used to get the raw data of an asset.
// After writing a valid JSON to this characteristic,
// the characteristic value will be the raw data of the requested asset.
// A valid JSON looks like this. `{"id":"1.jpg","width":320,"height":240}`
type GetAsset struct {
	*characteristic.Bytes
}

func NewGetAsset() *GetAsset {
	b := characteristic.NewBytes(TypeGetAsset)
	b.Perms = []string{characteristic.PermRead, characteristic.PermWrite}
	b.Value = []byte{}

	return &GetAsset{b}
}

type GetAssetRequest struct {
	ID     string `json:"id"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

// TakeSnapshot is used to take a snapshot.
// After writing `true` to this characteristic,
// a snapshot is taked and persisted on disk.
type TakeSnapshot struct {
	*characteristic.Bool
}

func NewTakeSnapshot() *TakeSnapshot {
	b := characteristic.NewBool(TypeTakeSnapshot)
	b.Description = "Take Snapshot"
	b.Perms = []string{characteristic.PermWrite}

	return &TakeSnapshot{b}
}
