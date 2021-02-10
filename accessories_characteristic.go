package homekit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/nfnt/resize"
	"github.com/radovskyb/watcher"
)

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
