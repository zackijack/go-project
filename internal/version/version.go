package version

import (
	"fmt"
	"runtime"
)

// Version is the application semantic version.
var Version string

// GitCommit is the Git commit SHA1.
var GitCommit string

// BuildDate is the date when the application was built.
var BuildDate string

// GoVersion is the version of the Go compiler used.
var GoVersion = runtime.Version()

// OsArch is the version of operating system and architecture.
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
