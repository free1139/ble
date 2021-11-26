package main

import (
	"encoding/binary"

	"github.com/gwaylib/errors"
)

// read/write
type AxpertDeviceSettingParam struct {
	data  [20]byte // 0x2A0C
	data1 [20]byte // 0x2A0D
	data2 [20]byte // 0x2A0E
}

func ParseAxpertDeviceSettingParam(data, data1, data2 []byte) (*AxpertDeviceSettingParam, error) {
	if len(data) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	if len(data1) < 20 {
		return nil, errors.New("need data1 len >= 20")
	}
	if len(data2) < 20 {
		return nil, errors.New("need data2 len >= 20")
	}
	a := &AxpertDeviceSettingParam{}
	copy(a.data[:], data)
	copy(a.data1[:], data1)
	copy(a.data2[:], data2)
	return a, nil
}

// unit 1V
// HV:208/220/230/240
// LV:110/120
func (a *AxpertDeviceSettingParam) GetACOutputVolatge() uint16 {
	return binary.LittleEndian.Uint16(a.data[:2])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetACOutputVolatge(val uint16) error {
	switch val {
	case 208, 220, 230, 240:
	case 110, 120:
	default:
		return errors.New("unsupport value")
	}
	binary.LittleEndian.PutUint16(a.data[0:2], val)
	return nil
}

// unit 0.1Hz
// 50.0/60.0Hz
func (a *AxpertDeviceSettingParam) GetACOutputFrequency() uint16 {
	return binary.LittleEndian.Uint16(a.data[2:4])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetACOutputFrequency(val uint16) error {
	switch val {
	case 500, 600:
	default:
		return errors.New("unsupport value")
	}
	binary.LittleEndian.PutUint16(a.data[2:4], val)
	return nil
}

// unit A
// range refet to uuid 0x2A06~0x2A07
func (a *AxpertDeviceSettingParam) GetMaxChargingCurrents() byte {
	return a.data[4]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetMaxChargingCurrents(val byte) error {
	// TODO: check range
	//switch val {
	//case 500, 600:
	//default:
	//	return errors.New("unsupport value")
	//}
	a.data[4] = val
	return nil
}

// unit A
// range refet to uuid 0x2A08~0x2A09
func (a *AxpertDeviceSettingParam) GetMaxUtilityChargingCurrents() byte {
	return a.data[5]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetMaxUtilityChargingCurrents(val byte) error {
	// TODO: check range
	//switch val {
	//case 500, 600:
	//default:
	//	return errors.New("unsupport value")
	//}
	a.data[5] = val
	return nil
}

// unit 0.1V
func (a *AxpertDeviceSettingParam) GetChargingFloatVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[6:8])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetChargingFloatVoltage(val uint16) error {
	binary.LittleEndian.PutUint16(a.data[6:8], val)
	return nil
}

// unit 0.1V
func (a *AxpertDeviceSettingParam) GetChargingBulkVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[8:10])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetChargingBulkVoltage(val uint16) error {
	binary.LittleEndian.PutUint16(a.data[8:10], val)
	return nil
}

// unit 0.1V
func (a *AxpertDeviceSettingParam) GetBatteryUnderVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[10:12])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryUnderVoltage(val uint16) error {
	binary.LittleEndian.PutUint16(a.data[10:12], val)
	return nil
}

// unit 0.1V
func (a *AxpertDeviceSettingParam) GetBatteryRechargeVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[12:14])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryRechargeVoltage(val uint16) error {
	binary.LittleEndian.PutUint16(a.data[12:14], val)
	return nil
}

// unit 0.1V
func (a *AxpertDeviceSettingParam) GetBatteryReDischargeVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data[14:16])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryReDischargeVoltage(val uint16) error {
	binary.LittleEndian.PutUint16(a.data[14:16], val)
	return nil
}

// unit Hex
// 0x00 for appliance;
// 0x01 for ups
func (a *AxpertDeviceSettingParam) GetACInputVoltageRange() byte {
	return a.data[16]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetACInputVoltageRange(val byte) error {
	switch val {
	case 0x00, 0x01:
	default:
		return errors.New("unsupport value")
	}
	a.data[16] = val
	return nil
}

// unit Hex
// 0x00 for utility first;
// 0x01 for solar first;
// 0x02 for SBU priority;
func (a *AxpertDeviceSettingParam) GetOutputSourcePriority() byte {
	return a.data[17]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetOutputSourcePriority(val byte) error {
	switch val {
	case 0x00, 0x01, 0x02:
	default:
		return errors.New("unsupport value")
	}
	a.data[17] = val
	return nil
}

// unit Hex
// 0x00 for utility first;
// 0x01 for solar first;
// 0x02 for solar and utility;
// 0x03 for only solar charging;
func (a *AxpertDeviceSettingParam) GetChargerSourcePriority() byte {
	return a.data[18]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetChargeSourcePriority(val byte) error {
	switch val {
	case 0x00, 0x01, 0x02, 0x03:
	default:
		return errors.New("unsupport value")
	}
	a.data[18] = val
	return nil
}

// unit Hex
// 0x00 for AGM;
// 0x01 for Flooded;
// 0x02 for User define;
// 0x03 for PYL;
// 0x05 for WECO;
// 0x06 for Soltao;
// Not in the above range is other lithium batteries
func (a *AxpertDeviceSettingParam) GetBatteryType() byte {
	return a.data[19]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryType(val byte) error {
	a.data[19] = val
	return nil
}

func (a *AxpertDeviceSettingParam) GetFaultCodeRecordOn() bool {
	onBit := byte(1 << 0)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetFaultCodeRecordOn(on bool) error {
	onBit := byte(1 << 0)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}
func (a *AxpertDeviceSettingParam) GetPrimaryInterruptAlarmOn() bool {
	onBit := byte(1 << 1)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetPrimaryInterruptAlarmOn(on bool) error {
	onBit := byte(1 << 1)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}
func (a *AxpertDeviceSettingParam) GetBacklightOn() bool {
	onBit := byte(1 << 2)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetBacklightOn(on bool) error {
	onBit := byte(1 << 2)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}
func (a *AxpertDeviceSettingParam) GetOverTemperatureRestartOn() bool {
	onBit := byte(1 << 3)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetOverTemperatureRestartOn(on bool) error {
	onBit := byte(1 << 3)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}
func (a *AxpertDeviceSettingParam) GetOverloadRestartOn() bool {
	onBit := byte(1 << 4)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetOverOverloadRestartOn(on bool) error {
	onBit := byte(1 << 4)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}

// 1 minute timeout
func (a *AxpertDeviceSettingParam) GetLCDDisplayTimeoutOn() bool {
	onBit := byte(1 << 5)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetLCDDisplayTimeoutOn(on bool) error {
	onBit := byte(1 << 5)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}

func (a *AxpertDeviceSettingParam) GetSolarFeedToGridOn() bool {
	onBit := byte(1 << 6)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetSolarFeedToGridOn(on bool) error {
	onBit := byte(1 << 6)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}
func (a *AxpertDeviceSettingParam) GetOverloadByPassFnOn() bool {
	onBit := byte(1 << 7)
	return a.data1[0]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetOverloadByPassFnOn(on bool) error {
	onBit := byte(1 << 7)
	if !on {
		a.data1[0] = (a.data1[0] & (^onBit))
	} else {
		a.data1[0] = (a.data1[0] | onBit)
	}
	return nil
}
func (a *AxpertDeviceSettingParam) GetSilenceOrOpenBuzzerOn() bool {
	onBit := byte(1 << 0)
	return a.data1[1]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetSilenceOrOpenBuzzerOn(on bool) error {
	onBit := byte(1 << 0)
	if !on {
		a.data1[1] = (a.data1[1] & (^onBit))
	} else {
		a.data1[1] = (a.data1[1] | onBit)
	}
	return nil
}
func (a *AxpertDeviceSettingParam) GetBatteryEqualizationOn() bool {
	onBit := byte(1 << 1)
	return a.data1[1]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetBatteryEqualizationOn(on bool) error {
	onBit := byte(1 << 1)
	if !on {
		a.data1[1] = (a.data1[1] & (^onBit))
	} else {
		a.data1[1] = (a.data1[1] | onBit)
	}
	return nil
}

// f15~f11 Reseved
func (a *AxpertDeviceSettingParam) GetOverloadByPassFn2On() bool {
	onBit := byte(1 << 2)
	return a.data1[1]&onBit == onBit
}
func (a *AxpertDeviceSettingParam) SetOverloadByPassFn2On(on bool) error {
	onBit := byte(1 << 2)
	if !on {
		a.data1[1] = (a.data1[1] & (^onBit))
	} else {
		a.data1[1] = (a.data1[1] | onBit)
	}
	return nil
}

// uint hex
// 0x00: single machine output
// 0x01: parallel output
// 0x02: Phase 1 of 3 Phase output
// 0x03: Phase 2 of 3 Phase output
// 0x04: Phase 3 of 3 Phase output
func (a *AxpertDeviceSettingParam) GetOutputMode() byte {
	return a.data1[2]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetOutputMode(val byte) error {
	switch val {
	case 0x00, 0x01, 0x02, 0x03, 0x04:
	default:
		return errors.New("unsupport value")
	}
	a.data1[2] = val
	return nil
}

// uint hex
// 0x00: As long as one unit of inverters has connected PV, parallel system will consider PV OK;
// 0x01: Only all of inverters have connected PV, parallel system will consider PV OK;
func (a *AxpertDeviceSettingParam) GetPVOkConditionOfParallel() byte {
	return a.data1[3]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetPVOkConditionOfParallel(val byte) error {
	switch val {
	case 0x00, 0x01:
	default:
		return errors.New("unsupport value")
	}
	a.data1[3] = val
	return nil
}

// uint hex
// 0x00:PV input max current will be the max charged current
// 0x01: PV input max power will be the sum of the max charged power and loads power.
func (a *AxpertDeviceSettingParam) GetPVPowerBalance() byte {
	return a.data1[4]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetPVPowerBalance(val byte) error {
	switch val {
	case 0x00, 0x01:
	default:
		return errors.New("unsupport value")
	}
	a.data1[4] = val
	return nil
}

// uint hex
// 0x00: Do nothing
// 0x01: Default parameters
func (a *AxpertDeviceSettingParam) GetControlParameterToDefaultValue() byte {
	return a.data1[5]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetControlParameterToDefaultValue(val byte) error {
	switch val {
	case 0x00, 0x01:
	default:
		return errors.New("unsupport value")
	}
	a.data1[5] = val
	return nil
}

// uint 5Min
func (a *AxpertDeviceSettingParam) GetBatteryEqualizationTime() uint16 {
	return binary.LittleEndian.Uint16(a.data1[6:8])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryEqualizationTime(val uint16) error {
	binary.LittleEndian.PutUint16(a.data1[6:8], val)
	return nil
}

// uint 1Day
func (a *AxpertDeviceSettingParam) GetBatteryEqualizationPeriod() uint16 {
	return binary.LittleEndian.Uint16(a.data1[8:10])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryEqualizationPeriod(val uint16) error {
	binary.LittleEndian.PutUint16(a.data1[8:10], val)
	return nil
}

// uint 0.01V
func (a *AxpertDeviceSettingParam) GetBatteryEqualizationVoltage() uint16 {
	return binary.LittleEndian.Uint16(a.data1[10:12])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryEqualizationVoltage(val uint16) error {
	binary.LittleEndian.PutUint16(a.data1[10:12], val)
	return nil
}

// uint 0.01V
func (a *AxpertDeviceSettingParam) GetBatteryEqualizationOverTime() uint16 {
	return binary.LittleEndian.Uint16(a.data1[12:14])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryEqualizationOverTime(val uint16) error {
	binary.LittleEndian.PutUint16(a.data1[12:14], val)
	return nil
}

// unix hex
// 0x00: Inactive
// 0x01: Active
func (a *AxpertDeviceSettingParam) GetBatteryEqualizationActive() byte {
	return a.data1[14]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetBatteryEqualizationActive(val byte) error {
	a.data1[14] = val
	return nil
}

// a.data1[15] for reserve

// uint 5Min
func (a *AxpertDeviceSettingParam) GetCVChargingTime() uint16 {
	return binary.LittleEndian.Uint16(a.data1[16:18])
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetCVChargingTime(val uint16) error {
	binary.LittleEndian.PutUint16(a.data1[16:18], val)
	return nil
}

// a.data1[18:20] for reserve

// unix hex
// 0x00: Automatically
// 0x01: Online
// 0x02: ECO
func (a *AxpertDeviceSettingParam) GetOperationLogic() byte {
	return a.data2[0]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetOperationLogic(val byte) error {
	a.data2[0] = val
	return nil
}

// unix A
func (a *AxpertDeviceSettingParam) GetMaxDischargingCurrents() byte {
	return a.data2[1]
}

// warn need init first
func (a *AxpertDeviceSettingParam) SetMaxDischargingCurrents(val byte) error {
	a.data2[0] = val
	return nil
}
