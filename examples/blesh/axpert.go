package main

import (
	"fmt"

	"github.com/go-ble/ble"
	"github.com/gwaylib/errors"
	"github.com/urfave/cli"
)

const (
	CH_UUID_AXPERT_VERSION                  = "2A01"
	CH_UUID_AXPERT_SN                       = "2A02"
	CH_UUID_AXPERT_WORKING_1                = "2A03"
	CH_UUID_AXPERT_WORKING_2                = "2A04"
	CH_UUID_AXPERT_DEVICE_RATING_INFO       = "2A05"
	CH_UUID_AXPERT_DEVICE_CHARGING_1        = "2A06"
	CH_UUID_AXPERT_DEVICE_CHARGING_2        = "2A07"
	CH_UUID_AXPERT_DEVICE_UTILITY_1         = "2A08"
	CH_UUID_AXPERT_DEVICE_UTILITY_2         = "2A09"
	CH_UUID_AXPERT_DEVICE_DEFAULT_SETTING_1 = "2A0A"
	CH_UUID_AXPERT_DEVICE_DEFAULT_SETTING_2 = "2A0B"
	CH_UUID_AXPERT_DEVICE_SETTING_1         = "2A0C"
	CH_UUID_AXPERT_DEVICE_SETTING_2         = "2A0D"
	CH_UUID_AXPERT_DEVICE_SETTING_3         = "2A0E"
	CH_UUID_AXPERT_PV_INFO_1                = "2A11"
	CH_UUID_AXPERT_PV_INFO_2                = "2A12"
	CH_UUID_AXPERT_PV_INFO_3                = "2A13"
	CH_UUID_AXPERT_PV_INFO_4                = "2A14"
	CH_UUID_AXPERT_RESERVED_1               = "2A21"
	CH_UUID_AXPERT_RESERVED_2               = "2A22"
)

func readAxpert(c *cli.Context) error {
	u, err := ble.Parse("")
	if err != nil {
		return errInvalidUUID
	}
	_ = u
	if u := curr.profile.Find(ble.NewCharacteristic(curr.uuid)); u != nil {
		b, err := curr.client.ReadCharacteristic(u.(*ble.Characteristic))
		if err != nil {
			return errors.As(err, "can't read characteristic")
		}
		fmt.Printf("    Value         %x | %q\n", b, b)
		return nil
	}
	return nil
}
