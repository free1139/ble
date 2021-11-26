package main

import (
	"strings"

	"github.com/gwaylib/errors"
)

type AxpertVersion struct {
	data [20]byte
}

func ParseAxpertVersion(data []byte) (*AxpertVersion, error) {
	if len(data) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	a := &AxpertVersion{}
	copy(a.data[:], data)
	return a, nil
}
func (a *AxpertVersion) String() string {
	return strings.TrimSpace(string(a.data[:]))
}
func (a *AxpertVersion) ProtoID() string {
	return string(a.data[:2])
}
func (a *AxpertVersion) MainBoard() string {
	return string(a.data[3:10])
}
func (a *AxpertVersion) BLEBoard() string {
	return string(a.data[11:])
}
