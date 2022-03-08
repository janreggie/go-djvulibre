package djvu

import (
	"runtime"
)

const (
	osWindows = "windows"
	osMac     = "darwin"
	osLinux   = "linux"
)

const unsupportedSystem = "unsupported system " + runtime.GOOS
