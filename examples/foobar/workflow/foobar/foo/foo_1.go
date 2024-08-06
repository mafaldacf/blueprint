package foo

import (
    strconvpkg_1 "strconv"
)

func Int64ToString(i int64) string {
    return strconvpkg_1.FormatInt(i, 10)
}
