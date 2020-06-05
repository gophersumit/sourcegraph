package version

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

const devVersion = "0.0.0+dev"                              // version string for unreleased development builds
var devTimestamp = strconv.FormatInt(time.Now().Unix(), 10) // build timestamp for unreleased development builds

// version is configured at build time via ldflags like this:
// -ldflags "-X github.com/sourcegraph/sourcegraph/internal/version.version=1.2.3"
//
// The version may not be semver-compatible, e.g. `insiders` or `65769_2020-06-05_9bd91a3`.
var version = devVersion

// Version returns the version string configured at build time.
func Version() string {
	return version
}

// IsDev reports whether the version string is an unreleased development build.
func IsDev(version string) bool {
	return version == devVersion
}

// Mock is used by tests to mock the result of Version and IsDev.
func Mock(mockVersion string) {
	version = mockVersion
}

// timestamp is the build timestamp configured at build time via ldflags like this:
// -ldflags "-X github.com/sourcegraph/sourcegraph/internal/version.timestamp=$UNIX_SECONDS"
var timestamp = devTimestamp

// HowLongOutOfDate returns a time in months since this build of Sourcegraph was created. It is
// based on a constant baked into the Go binary at build time.
func HowLongOutOfDate(currentVersion string) (int, error) {
	buildUnixTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("unable to parse version build timestamp: %w", err)
	}
	buildTime := time.Unix(buildUnixTimestamp, 0)

	now := time.Now()
	if buildTime.After(now) {
		return 0, errors.New("version build timestamp is in the future")
	}
	daysSinceBuild := now.Sub(buildTime).Hours() / 24

	return monthsFromDays(daysSinceBuild), nil
}

// monthsFromDays roughly determines the number of months given days
func monthsFromDays(days float64) int {
	const daysInAMonth = 30
	months := math.Floor(days / daysInAMonth)
	return int(months)
}
