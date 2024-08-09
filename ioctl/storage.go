package ioctl

import (
	"github.com/jimbertools/volmgmt/ioctlcode"
	"github.com/jimbertools/volmgmt/ioctltype"
)

// I/O control codes for mass storage devices.
var (
	StorageCheckVerify       = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0200, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_STORAGE_CHECK_VERIFY
	StorageCheckVerify2      = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0200, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_CHECK_VERIFY2
	StorageEjectMedia        = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0202, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_STORAGE_EJECT_MEDIA
	StorageEjectionControl   = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0250, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_EJECTION_CONTROL
	StorageFindNewDevices    = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0206, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_STORAGE_FIND_NEW_DEVICES
	StorageGetDeviceNumber   = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0420, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_GET_DEVICE_NUMBER
	StorageMediaSerialNumber = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0304, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_GET_MEDIA_SERIAL_NUMBER
	StorageGetMediaTypes     = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0300, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_GET_MEDIA_TYPES
	StorageGetMediaTypesEx   = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0301, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_GET_MEDIA_TYPES_EX
	StorageLoadMedia         = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0203, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_STORAGE_LOAD_MEDIA
	StorageLoadMedia2        = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0203, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_LOAD_MEDIA2
	StorageMCNControl        = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0251, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_MCN_CONTROL
	StorageMediaRemoval      = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0201, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_STORAGE_MEDIA_REMOVAL
	StoragePredictFailure    = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0440, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_PREDICT_FAILURE
	StorageQueryProperty     = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0500, ioctlcode.MethodBuffered, ioctlcode.AccessAny)       // IOCTL_STORAGE_QUERY_PROPERTY
	StorageRelease           = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0205, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_STORAGE_RELEASE
	StorageReserve           = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0204, ioctlcode.MethodBuffered, ioctlcode.AccessRead)      // IOCTL_STORAGE_RESERVE
	StorageResetBus          = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0400, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_STORAGE_RESET_BUS
	StorageResetDevice       = ioctlcode.New(ioctltype.DeviceMassStorage, 0x0401, ioctlcode.MethodBuffered, ioctlcode.AccessReadWrite) // IOCTL_STORAGE_RESET_DEVICE
)
