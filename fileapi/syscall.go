package fileapi

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/jimbertools/volmgmt/fileref"
	"golang.org/x/sys/windows"
)

var (
	// ErrEmptyBuffer is returned when a nil or zero-sized buffer is provided
	// to a system call.
	ErrEmptyBuffer = errors.New("nil or empty buffer provided")
)

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procOpenFileByID                 = modkernel32.NewProc("OpenFileById")
	procReOpenFile                   = modkernel32.NewProc("ReOpenFile")
	procGetFileInformationByHandle   = modkernel32.NewProc("GetFileInformationByHandle")
	procGetFileInformationByHandleEx = modkernel32.NewProc("GetFileInformationByHandleEx")
	procSetFileInformationByHandle   = modkernel32.NewProc("SetFileInformationByHandle")
)

// OpenFileByID opens a file by its file ID. The file will be opened with the
// given access, share mode and flags.
//
// The handle provided can be to any file or on the volume, or to the volume
// itself.
func OpenFileByID(peer syscall.Handle, id fileref.ID, access, shareMode, flags uint32) (handle syscall.Handle, err error) {
	d := id.Descriptor()

	r0, _, e := syscall.SyscallN(
		procOpenFileByID.Addr(),
		uintptr(peer),
		uintptr(unsafe.Pointer(&d)),
		uintptr(access),
		uintptr(shareMode),
		0,
		uintptr(flags))
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e != 0 {
			err = syscall.Errno(e)
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

// ReOpenFile opens a file by its file ID. The file will be opened with the
// given access, share mode and flags.
//
// The handle provided can be to any file or on the volume, or to the volume
// itself.
func ReOpenFile(original syscall.Handle, access, shareMode, flags uint32) (handle syscall.Handle, err error) {
	r0, _, e := syscall.SyscallN(
		procReOpenFile.Addr(),
		uintptr(original),
		uintptr(access),
		uintptr(shareMode),
		uintptr(flags))
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e != 0 {
			err = e
		} else {
			err = syscall.EINVAL
		}
	}

	return
}

// GetFileInformationByHandle retrieves standard information about the file
// represented by the given system handle.
func GetFileInformationByHandle(handle syscall.Handle) (info syscall.ByHandleFileInformation, err error) {
	r0, _, e := syscall.SyscallN(
		procGetFileInformationByHandle.Addr(),
		uintptr(handle),
		uintptr(unsafe.Pointer(&info)))
	if r0 == 0 {
		if e != 0 {
			err = e
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// GetFileInformationByHandleEx retrieves information about the file
// represented by the given system handle. The type of information returned
// is determined by class.
func GetFileInformationByHandleEx(handle syscall.Handle, info FileInfoUnmarshaler) (err error) {
	buffer := make([]byte, info.Size())
	if len(buffer) == 0 {
		return ErrEmptyBuffer
	}

	r0, _, e := syscall.SyscallN(
		procGetFileInformationByHandleEx.Addr(),
		uintptr(handle),
		uintptr(info.Class()),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)))
	if r0 == 0 {
		if e != 0 {
			return e
		}
		return syscall.EINVAL
	}

	return info.UnmarshalBinary(buffer)
}

// SetFileInformationByHandle updates the file identified by the given system
// handle. The type of information file information set is determined by
// class.
func SetFileInformationByHandle(handle syscall.Handle, info FileInfoMarshaler) (err error) {
	data, err := info.MarshalBinary()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return ErrEmptyBuffer
	}

	r0, _, e := syscall.SyscallN(
		procSetFileInformationByHandle.Addr(),
		uintptr(handle),
		uintptr(info.Class()),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)))
	if r0 == 0 {
		if e != 0 {
			return e
		}
		return syscall.EINVAL
	}
	return nil
}
