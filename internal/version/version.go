package version

import "fmt"

// The values below can be overridden at build time using -ldflags
// (e.g. go build -ldflags "-X github.com/OneideLuizSchneider/blitzctl/internal/version.Version=v0.1.0").
var (
	Version   = "dev"
	GitCommit = "unknown"
	BuildDate = "unknown"
)

// String returns a human friendly version string with optional metadata.
func String() string {
	base := Version

	switch {
	case GitCommit == "unknown" && BuildDate == "unknown":
		return base
	case BuildDate == "unknown":
		return fmt.Sprintf("%s (commit: %s)", base, GitCommit)
	case GitCommit == "unknown":
		return fmt.Sprintf("%s (built: %s)", base, BuildDate)
	default:
		return fmt.Sprintf("%s (commit: %s, built: %s)", base, GitCommit, BuildDate)
	}
}
