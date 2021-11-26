package main

import (
	"encoding/binary"

	"github.com/gwaylib/errors"
)

type AxpertDeviceDefaultSetting struct {
	data  [20]byte
	data1 [20]byte
}

func ParseAxpertDeviceDefaultSetting(data, data1 []byte) (*AxpertDeviceDefaultSetting, error) {
	if len(data) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	if len(data1) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	a := &AxpertDeviceDefaultSetting{}
	copy(a.data[:], data)
	copy(a.data1[:], data1)
	return a, nil
}

// unit 1V
func (a *AxpertDeviceDefaultSetting) ACDefaultOutputVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[:2])
}

// 0.1Hz
func (a *AxpertDeviceDefaultSetting) ACDefaultOutputFrequency() uint16 {
	return binary.LittleEndian.Uint16(a.data[2:4])
}

// uint A
func (a *AxpertDeviceDefaultSetting) DefaultMaxChargingCurrents() byte {
	return a.data[4]
}

// uint A
func (a *AxpertDeviceDefaultSetting) DefaultMaxUtilityChargingCurrents() byte {
	return a.data[5]
}

// 0.1V
func (a *AxpertDeviceDefaultSetting) DefaultChargingFloatVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[6:8])
}

// 0.1V
func (a *AxpertDeviceDefaultSetting) DefaultChargingBulkVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[8:10])
}

// 0.1V
func (a *AxpertDeviceDefaultSetting) DefaultBatteryUnderVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[10:12])
}

// 0.1V
func (a *AxpertDeviceDefaultSetting) DefaultBatteryRechargeVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[12:14])
}

// 0.1V
func (a *AxpertDeviceDefaultSetting) DefaultBatteryReDischargeVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[14:16])
}
func (a *AxpertDeviceDefaultSetting) DefaultACInputVoltageRange() byte {
	return a.data[16]
}
func (a *AxpertDeviceDefaultSetting) DefaultOutputSourcePriority() byte {
	return a.data[17]
}
func (a *AxpertDeviceDefaultSetting) DefaultChargerSourcePriority() byte {
	return a.data[18]
}
func (a *AxpertDeviceDefaultSetting) DefaultBatteryType() byte {
	return a.data[19]
}
func (a *AxpertDeviceDefaultSetting) FlagFaultCodeRecordOn() bool {
	onBit := byte(1 << 0)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagPrimaryInterruptAlarmOn() bool {
	onBit := byte(1 << 1)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagBacklightOn() bool {
	onBit := byte(1 << 2)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagOverTemperatureRestartOn() bool {
	onBit := byte(1 << 3)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagOverloadRestartOn() bool {
	onBit := byte(1 << 4)
	return a.data1[0]&onBit == onBit
}

// 1 minute timeout
func (a *AxpertDeviceDefaultSetting) FlagLCDDisplayTimeoutOn() bool {
	onBit := byte(1 << 5)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagSolarFeedToGridOn() bool {
	onBit := byte(1 << 6)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagOverloadByPassFnOn() bool {
	onBit := byte(1 << 7)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagSilenceOrOpenBuzzerOn() bool {
	onBit := byte(1 << 0)
	return a.data1[1]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) FlagBatteryEqualizationOn() bool {
	onBit := byte(1 << 1)
	return a.data1[1]&onBit == onBit
}

// f15~f11 Reseved
func (a *AxpertDeviceDefaultSetting) FlagOverloadByPassFn2On() bool {
	onBit := byte(1 << 2)
	return a.data1[1]&onBit == onBit
}

// hex
func (a *AxpertDeviceDefaultSetting) DefaultOutputMode() byte {
	return a.data1[2]
}

// hex
func (a *AxpertDeviceDefaultSetting) DefaultPVParallelCondition() byte {
	return a.data1[3]
}

// hex
func (a *AxpertDeviceDefaultSetting) DefaultPVPowerBalance() byte {
	return a.data1[4]
}

// data1[6:10] for reserved
func (a *AxpertDeviceDefaultSetting) DefaultCVTimeCanBeSet() byte {
	return a.data1[5]
}

// B7~B6 for reserved
func (a *AxpertDeviceDefaultSetting) OutputVoltOn() bool {
	onBit := byte(1 << 0)
	return a.data1[10]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) ParallelSetOn() bool {
	onBit := byte(1 << 1)
	return a.data1[10]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) CVTimeSetOn() bool {
	onBit := byte(1 << 2)
	return a.data1[10]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) ChargeModeSetOn() bool {
	onBit := byte(1 << 3)
	return a.data1[10]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) PowerSaveOn() bool {
	onBit := byte(1 << 4)
	return a.data1[10]&onBit == onBit
}
func (a *AxpertDeviceDefaultSetting) AxpertKingOn() bool {
	onBit := byte(1 << 5)
	return a.data1[10]&onBit == onBit
}
