package gpu

import (
	"errors"
	"fmt"

	"github.com/gonutz/d3d9"
)

type Type string

const (
	TypeAMD    = "AMD"
	TypeNvidia = "Nvidia"
	TypeIntel  = "Intel"

	TypeUnknown = "unknown"
)

const (
	vendorIdAMD1 = 0x1002
	vendorIdAMD2 = 0x1022
	vendorNv     = 0x10de
)

func DetectGPUType() (Type, error) {
	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	if err != nil {
		return TypeUnknown, fmt.Errorf("init d3d: %w", err)
	}
	defer d3d.Release()
	ident, err := d3d.GetAdapterIdentifier(d3d9.ADAPTER_DEFAULT, 0)
	if err != nil {
		return TypeUnknown, fmt.Errorf("read adapter id: %w", err)
	}

	switch ident.VendorId {
	case vendorIdAMD1: fallthrough
	case vendorIdAMD2:
		return TypeAMD, nil
	}

	return TypeUnknown, errors.New("your GPU in not supported by this tool")
}

