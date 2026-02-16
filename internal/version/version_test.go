/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package version

import (
	"regexp"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}
}

func TestVersionFormat(t *testing.T) {
	// Version should follow semantic versioning format: vX.Y.Z
	// where X, Y, Z are non-negative integers
	semverPattern := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)

	if !semverPattern.MatchString(Version) {
		t.Errorf("Version %q does not follow semantic versioning format (vX.Y.Z)", Version)
	}
}

func TestVersionStartsWithV(t *testing.T) {
	if !strings.HasPrefix(Version, "v") {
		t.Errorf("Version %q should start with 'v'", Version)
	}
}

func TestVersionHasTwoDots(t *testing.T) {
	// Semantic version should have exactly two dots (MAJOR.MINOR.PATCH)
	dotCount := strings.Count(Version, ".")
	if dotCount != 2 {
		t.Errorf("Version %q should have exactly 2 dots, got %d", Version, dotCount)
	}
}

func TestVersionComponents(t *testing.T) {
	// Remove 'v' prefix and split into components
	versionWithoutV := strings.TrimPrefix(Version, "v")
	parts := strings.Split(versionWithoutV, ".")

	if len(parts) != 3 {
		t.Fatalf("Version should have 3 components (MAJOR.MINOR.PATCH), got %d", len(parts))
	}

	// Each component should be numeric
	for i, part := range parts {
		if part == "" {
			t.Errorf("Version component %d is empty", i)
		}

		// Check if component contains only digits
		for _, char := range part {
			if char < '0' || char > '9' {
				t.Errorf("Version component %d (%q) contains non-digit character: %c", i, part, char)
			}
		}
	}
}

func TestVersionNoWhitespace(t *testing.T) {
	if strings.TrimSpace(Version) != Version {
		t.Errorf("Version %q should not contain leading or trailing whitespace", Version)
	}

	if strings.Contains(Version, " ") {
		t.Errorf("Version %q should not contain spaces", Version)
	}
}

func TestVersionNoPreReleaseOrBuildMetadata(t *testing.T) {
	// Version should not contain pre-release identifiers (e.g., v1.0.0-alpha)
	// or build metadata (e.g., v1.0.0+20130313144700)
	if strings.Contains(Version, "-") {
		t.Errorf("Version %q should not contain pre-release identifier (dash)", Version)
	}

	if strings.Contains(Version, "+") {
		t.Errorf("Version %q should not contain build metadata (plus sign)", Version)
	}
}

// Benchmark the version string access
func BenchmarkVersionAccess(b *testing.B) {
	var v string
	for b.Loop() {
		v = Version
	}
	_ = v
}
