package main

import (
	"fmt"
	"strings"

	"github.com/gwaylib/errors"
)

type AxpertSN struct {
	data [20]byte
}

func ParseAxpertSN(data []byte) (*AxpertSN, error) {
	if len(data) < 20 {
		return nil, errors.New("need data len >= 20")
	}
	a := &AxpertSN{}
	copy(a.data[:], data)
	return a, nil
}

func (a *AxpertSN) SN() string {
	return strings.TrimSpace(string(a.data[:]))
}
func (a *AxpertSN) ModleID() string {
	return fmt.Sprintf("0x%x", a.data[19])
}
