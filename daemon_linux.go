// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.

// Package daemon linux version
package daemon

import (
	"os"
)

// Get the daemon properly
func newDaemon(config *Config) (Daemon, error) {
	// newer subsystem must be checked first
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return &systemDRecord{config}, nil
	}
	if _, err := os.Stat("/sbin/initctl"); err == nil {
		return &upstartRecord{config}, nil
	}
	return &systemVRecord{config}, nil
}

// Get executable path
func execPath() (string, error) {
	return os.Readlink("/proc/self/exe")
}
