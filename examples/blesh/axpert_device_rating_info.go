package main

import (
	"encoding/binary"

	"github.com/gwaylib/errors"
)

type AxpertDeviceRatingInfo struct {
	data [20]byte
}

func ParseAxpertDeviceRatingInfo(data []byte) (*AxpertDeviceRatingInfo, error) {
	if len(data) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	a := &AxpertDeviceRatingInfo{}
	copy(a.data[:], data)
	return a, nil
}

// unit 0.1V
func (a *AxpertDeviceRatingInfo) GridRatingVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[:2])
}

// uint 0.1A
func (a *AxpertDeviceRatingInfo) GridRatingCurrent() uint16 {
	return binary.LittleEndian.Uint16(a.data[2:4])
}

// uint 0.1V
func (a *AxpertDeviceRatingInfo) ACOutputRatingVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[4:6])
}

// uint 0.1Hz
func (a *AxpertDeviceRatingInfo) ACOutputRatingFrequency() uint16 {
	return binary.LittleEndian.Uint16(a.data[6:8])
}

// uint 0.1A
func (a *AxpertDeviceRatingInfo) ACOutputRatingCurrent() uint16 {
	return binary.LittleEndian.Uint16(a.data[8:10])
}

// unit VA
func (a *AxpertDeviceRatingInfo) ACOutputRatingApparentPower() uint16 {
	return binary.LittleEndian.Uint16(a.data[10:12])
}

// unit W
func (a *AxpertDeviceRatingInfo) ACOutputRatingActivePower() uint16 {
	return binary.LittleEndian.Uint16(a.data[12:14])
}

// unit 0.1V
func (a *AxpertDeviceRatingInfo) BatteryRatingVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[14:16])
}

// 0x00:Grid tie; 0x01:Off Grid; 0x10:Hybrid
func (a *AxpertDeviceRatingInfo) MachineType() byte {
	return a.data[16]
}

// 0x00:Transformerless; 0x01:Transformer.
func (a *AxpertDeviceRatingInfo) Topology() byte {
	return a.data[17]
}

// 0x00:signle machine output;
// 0x01:parallel output
// 0x02:Phase 1 of 3 Phase output
// 0x03:Phase 2 of 3 Phase output
// 0x04:Phase 3 of 3 Phase output
func (a *AxpertDeviceRatingInfo) OutputMode() byte {
	return a.data[18]
}

// 0x00:HV model; 0x01:LV model
func (a *AxpertDeviceRatingInfo) HVOrLVModel() byte {
	return a.data[19]
}
