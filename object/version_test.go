package object

import (
	"strings"
	"testing"
	"time"
)

func TestGetVersionInfo_Structure(t *testing.T) {
	info := GetVersionInfo()

	if info == nil {
		t.Fatal("expected VersionInfo to be non-nil")
	}

	if info.Version == "" {
		t.Error("expected Version to be non-empty")
	}

	if info.Commit == "" {
		t.Error("expected Commit to be non-empty")
	}

	if info.GoVersion == "" {
		t.Error("expected GoVersion to be non-empty")
	}

	if !strings.HasPrefix(info.GoVersion, "go") {
		t.Errorf("expected GoVersion to start with 'go', got: %s", info.GoVersion)
	}

	if info.OS == "" {
		t.Error("expected OS to be non-empty")
	}

	if info.Arch == "" {
		t.Error("expected Arch to be non-empty")
	}
}

func TestGetVersionInfo_Uptime(t *testing.T) {
	// Ensure startTime is set before calling GetVersionInfo
	time.Sleep(10 * time.Millisecond)

	info := GetVersionInfo()

	if info.Uptime == "" {
		t.Error("expected Uptime to be non-empty")
	}
}

func TestGetVersionInfo_Defaults(t *testing.T) {
	origVersion := BuildVersion
	origCommit := BuildCommit
	defer func() {
		BuildVersion = origVersion
		BuildCommit = origCommit
	}()

	BuildVersion = "v1.2.3"
	BuildCommit = "abc1234"

	info := GetVersionInfo()

	if info.Version != "v1.2.3" {
		t.Errorf("expected Version 'v1.2.3', got: %s", info.Version)
	}

	if info.Commit != "abc1234" {
		t.Errorf("expected Commit 'abc1234', got: %s", info.Commit)
	}
}
