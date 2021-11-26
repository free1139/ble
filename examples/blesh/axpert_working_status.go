package main

import (
	"encoding/binary"

	"github.com/gwaylib/errors"
)

type AxpertWorkingStatus struct {
	data  [20]byte
	data1 [20]byte
}

func ParseAxpertWorkingStatus(data, data1 []byte) (*AxpertWorkingStatus, error) {
	if len(data) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	if len(data1) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	a := &AxpertWorkingStatus{}
	copy(a.data[:], data)
	copy(a.data1[:], data1)
	return a, nil
}

// unit 0.1V
func (a *AxpertWorkingStatus) GridVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[:2])
}

// unit 0.1Hz
func (a *AxpertWorkingStatus) GridFrequency() uint16 {
	return binary.LittleEndian.Uint16(a.data[2:4])
}

// unit 0.1V
func (a *AxpertWorkingStatus) ACOutputVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[4:6])
}

// unit 0.1Hz
func (a *AxpertWorkingStatus) ACOutputFrequency() uint16 {
	return binary.LittleEndian.Uint16(a.data[6:8])
}

// unit VA
func (a *AxpertWorkingStatus) ACOutputApparentPower() uint16 {
	return binary.LittleEndian.Uint16(a.data[8:10])
}

// unit Watt
func (a *AxpertWorkingStatus) ACOutputPower() uint16 {
	return binary.LittleEndian.Uint16(a.data[10:12])
}

// unit 1%
func (a *AxpertWorkingStatus) ACOutputLoadPercent() uint16 {
	return binary.LittleEndian.Uint16(a.data[12:14])
}

// unit V
func (a *AxpertWorkingStatus) BUSVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[14:16])
}

// unit 0.01V
func (a *AxpertWorkingStatus) BatteryVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[16:18])
}

// unit A
func (a *AxpertWorkingStatus) BatteryCharging() uint16 {
	return binary.LittleEndian.Uint16(a.data[18:20])
}

// uint 1%
func (a *AxpertWorkingStatus) BatteryCapacity() uint16 {
	return binary.LittleEndian.Uint16(a.data1[:2])
}

// unit 'C
func (a *AxpertWorkingStatus) InverterSinkTemperature() uint16 {
	return binary.LittleEndian.Uint16(a.data1[2:4])
}

// unit A
func (a *AxpertWorkingStatus) BatteryDischarge() uint16 {
	return binary.LittleEndian.Uint16(a.data1[4:6])
}

// data[6] low, data[7] for reserved
func (a *AxpertWorkingStatus) ACChargingOn() bool {
	onBit := byte(1 << 0)
	return a.data1[6]&onBit == onBit
}
func (a *AxpertWorkingStatus) SCCChargingOn() bool {
	onBit := byte(1 << 1)
	return a.data1[6]&onBit == onBit
}
func (a *AxpertWorkingStatus) ChargingOn() bool {
	onBit := byte(1 << 2)
	return a.data1[6]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrPVLoss() bool {
	onBit := byte(1 << 0)
	return a.data1[8]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrInverterFault() bool {
	onBit := byte(1 << 1)
	return a.data1[8]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrBusOver() bool {
	onBit := byte(1 << 2)
	return a.data1[8]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrBusUnder() bool {
	onBit := byte(1 << 3)
	return a.data1[8]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrBusSoftFail() bool {
	onBit := byte(1 << 4)
	return a.data1[8]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrLineFail() bool {
	onBit := byte(1 << 5)
	return a.data1[8]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrOutputShort() bool {
	onBit := byte(1 << 6)
	return a.data1[8]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrInvererVoltageTooLow() bool {
	onBit := byte(1 << 7)
	return a.data1[8]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrInvererVoltageTooHight() bool {
	onBit := byte(1 << 0)
	return a.data1[9]&onBit == onBit
}

// Compile with ErrInverterFault, it's fault if InverterFault, otherwise warning
func (a *AxpertWorkingStatus) ErrOverTemperature() bool {
	onBit := byte(1 << 0)
	return a.data1[9]&onBit == onBit
}

// Compile with ErrInverterFault, it's fault if InverterFault, otherwise warning
func (a *AxpertWorkingStatus) ErrFanLocked() bool {
	onBit := byte(1 << 1)
	return a.data1[9]&onBit == onBit
}

// Compile with ErrInverterFault, it's fault if InverterFault, otherwise warning
func (a *AxpertWorkingStatus) ErrBatteryVoltageHigh() bool {
	onBit := byte(1 << 2)
	return a.data1[9]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrBatteryVoltageLow() bool {
	onBit := byte(1 << 3)
	return a.data1[9]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrOverCharge() bool {
	onBit := byte(1 << 4)
	return a.data1[9]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrBatteryUnderShutDown() bool {
	onBit := byte(1 << 5)
	return a.data1[9]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrBatteryDerating() bool {
	onBit := byte(1 << 6)
	return a.data1[9]&onBit == onBit
}

// Compile with ErrInverterFault, it's fault if InverterFault, otherwise warning
func (a *AxpertWorkingStatus) ErrOverLoad() bool {
	onBit := byte(1 << 7)
	return a.data1[9]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrERRPROMFault() bool {
	onBit := byte(1 << 0)
	return a.data1[10]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrInverterOverCurrent() bool {
	onBit := byte(1 << 1)
	return a.data1[10]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrInverterSoftFail() bool {
	onBit := byte(1 << 2)
	return a.data1[10]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrSelfTestFail() bool {
	onBit := byte(1 << 3)
	return a.data1[10]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrOutputDCVoltageOver() bool {
	onBit := byte(1 << 4)
	return a.data1[10]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrBatteryOpen() bool {
	onBit := byte(1 << 5)
	return a.data1[10]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrCurrentSensorFail() bool {
	onBit := byte(1 << 6)
	return a.data1[10]&onBit == onBit
}

// fault
func (a *AxpertWorkingStatus) ErrBatteryShort() bool {
	onBit := byte(1 << 7)
	return a.data1[10]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrPowerLimit() bool {
	onBit := byte(1 << 0)
	return a.data1[11]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrPVVoltageHigh() bool {
	onBit := byte(1 << 1)
	return a.data1[11]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrMPPTOverloadFault() bool {
	onBit := byte(1 << 2)
	return a.data1[11]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrMPPTOverloadWarning() bool {
	onBit := byte(1 << 3)
	return a.data1[11]&onBit == onBit
}

// warn
func (a *AxpertWorkingStatus) ErrBatteryTooLowToCharge() bool {
	onBit := byte(1 << 4)
	return a.data1[11]&onBit == onBit
}

func (a *AxpertWorkingStatus) WorkingMode() string {
	return string(a.data1[13])
}
