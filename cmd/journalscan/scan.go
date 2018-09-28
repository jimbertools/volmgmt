package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gentlemanautomaton/volmgmt/fileattr"
	"github.com/gentlemanautomaton/volmgmt/usn"
	"github.com/gentlemanautomaton/volmgmt/usnfilter"
	"github.com/gentlemanautomaton/volmgmt/volume"
)

func scan(path string, settings Settings) {
	fmt.Printf("Path: \"%s\"\n", path)

	if settings.Include != nil {
		fmt.Printf("Include: %s\n", settings.Include)
	}

	if settings.Exclude != nil {
		fmt.Printf("Exclude: %s\n", settings.Exclude)
	}

	vol, err := volume.New(path)
	if err != nil {
		fmt.Printf("Unable to read \"%s\": %v\n", err)
		return
	}
	defer vol.Close()

	printVolume(vol)

	journal := vol.Journal()
	defer journal.Close()

	data, err := journal.Query()
	if err != nil {
		fmt.Printf("Unable to access USN Journal: %v\n", err)
		return
	}

	fmt.Printf("USN Journal: Present, ID: %d, Oldest USN: %d, Next USN: %d, Supporting Versions: %d-%d\n", data.JournalID, data.FirstUSN, data.NextUSN, data.MinSupportedMajorVersion, data.MaxSupportedMajorVersion)

	fmt.Printf("Scanning MFT...")
	start := time.Now()
	cache, err := journal.Cache(usnfilter.IsDir, 0, data.FirstUSN)
	end := time.Now()
	duration := end.Sub(start)
	if err != nil {
		fmt.Printf("failed: %v. Ran %s.\n", err, duration)
		return
	}
	fmt.Printf("done. Completed in %s. Found %d directories.\n", duration, cache.Size())

	cacheUpdater := func(record usn.Record) {
		if usnfilter.IsDir(record) {
			cache.Set(record)
		}
	}

	cursor, cursorErr := journal.Cursor(cacheUpdater, settings.Reason, nil, cache.Filer)
	if cursorErr != nil {
		fmt.Printf("Unable to create USN journal cursor: %v\n", cursorErr)
		return
	}
	defer cursor.Close()

	buffer := make([]byte, 262144)
	i := 0
	for {
		records, cursorErr := cursor.Next(buffer)
		if cursorErr != nil {
			if cursorErr != io.EOF {
				fmt.Printf("Unable to retreive USN journal records: %v\n", cursorErr)
			}
			return
		}

		for _, record := range records {
			if settings.Include != nil {
				if record.Path == "" {
					if !settings.Include.MatchString(record.FileName) {
						continue
					}
				} else {
					if !settings.Include.MatchString(record.Path) {
						continue
					}
				}
			}

			if settings.Exclude != nil {
				if record.Path == "" {
					if settings.Exclude.MatchString(record.FileName) {
						continue
					}
				} else {
					if settings.Exclude.MatchString(record.Path) {
						continue
					}
				}
			}

			id := record.FileReferenceNumber.String()
			when := record.TimeStamp.In(settings.Location).Format("2006-01-02 15:04:05.000000 MST")
			attr := record.FileAttributes.Join("", fileattr.FormatCode)
			action := strings.ToUpper(record.Reason.Join("|", usn.ReasonFormatShort))

			fmt.Printf("%s  %d.%d  %-5s  %20s  \"%s\"  %s  %s\n", when, record.MajorVersion, record.MinorVersion, record.SourceInfo, id, record.Path, attr, action)
			i++
		}
	}
}

func printVolume(vol *volume.Volume) {
	label, labelErr := vol.Label()
	name, nameErr := vol.Name()
	devicePath, devicePathErr := vol.DevicePath()

	fmt.Printf("Volume Label: %s\n", strOrErr(label, labelErr))
	fmt.Printf("Volume Name: %s\n", strOrErr(name, nameErr))
	fmt.Printf("NT Namespace Device Path: %s\n", strOrErr(devicePath, devicePathErr))
	fmt.Printf("Device Information: Number %d, Partition %d, Type %d\n", vol.DeviceNumber(), vol.PartitionNumber(), vol.DeviceType())
	fmt.Printf("Device Description: Removable: %t, Vendor: %s, Product: %s, Revision: %s, OS S/N: %s\n", vol.RemovableMedia(), vol.VendorID(), vol.ProductID(), vol.ProductRevision(), vol.SerialNumber())
}

func strOrErr(s string, err error) string {
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return fmt.Sprintf("\"%s\"", s)
}
