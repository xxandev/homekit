package homekit

import (
	"fmt"
	"regexp"
)

type SlaveAccessory struct {
	ID   uint64 `json:"id"`
	SN   string `json:"sn"`
	Name string `json:"name,omitempty"`
	Type uint8  `json:"type,omitempty"`
}

func (acc *SlaveAccessory) Example() {}

func (acc *SlaveAccessory) Valid() error {
	if !regexp.MustCompile(`^[a-zA-Z0-9\-]+$`).MatchString(acc.SN) || len(acc.SN) < 1 {
		return fmt.Errorf("serial number cannot be empty and should only contain (a-z, A-Z, 0-9, -)")
	}
	// if !regexp.MustCompile(`^[a-zA-Z0-9\s\-]+$`).MatchString(acc.Name) || len(acc.Name) < 1 {
	// 	return fmt.Errorf("name cannot be empty and should only contain (a-z, A-Z, 0-9, -, space)")
	// }
	return nil
}
