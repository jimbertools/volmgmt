// +build windows

package volumeapi

import "unicode/utf16"

// utf16ToSplitString splits a set of null-separated utf16 characters and
// returns a slice of substrings between those separators.
func utf16ToSplitString(s []uint16) []string {
	var values []string
	cut := 0
	for i, v := range s {
		if v == 0 {
			values = append(values, string(utf16.Decode(s[cut:i])))
			cut = i + 1
		}
	}
	return values
}
