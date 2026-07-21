//go:build windows

package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type SHQUERYRBINFO struct {
	cbSize      uint32
	i64Size     int64
	i64NumItems int64
}

var (
	shell32                = windows.NewLazySystemDLL("shell32.dll")
	shQueryRecycleBinW     = shell32.NewProc("SHQueryRecycleBinW")
	procSHEmptyRecycleBinW = shell32.NewProc("SHEmptyRecycleBinW")
)

func IsRecycleBinEmpty() (bool, error) {
	info := SHQUERYRBINFO{
		cbSize: uint32(unsafe.Sizeof(SHQUERYRBINFO{})),
	}

	r, _, err := shQueryRecycleBinW.Call(
		0, // all drives
		uintptr(unsafe.Pointer(&info)),
	)

	if r != 0 {
		return false, err
	}

	return info.i64NumItems == 0, nil
}

const (
	SHERB_NOCONFIRMATION = 0x00000001
	SHERB_NOPROGRESSUI   = 0x00000002
	SHERB_NOSOUND        = 0x00000004
)

func emptyBin() {
	// hwnd = 0, pszRootPath = nil => all drives
	_, _, _ = procSHEmptyRecycleBinW.Call(
		0,
		0,
		uintptr(SHERB_NOCONFIRMATION),
	)
	_ = unsafe.Pointer(nil)
}	
