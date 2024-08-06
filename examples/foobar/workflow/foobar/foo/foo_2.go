package foo

import (
	strconvpkg_2 "strconv"
)

func StringToInt64(s string) (int64, error) {
	// Parse the string into an int64 with base 10
	num, err := strconvpkg_2.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
