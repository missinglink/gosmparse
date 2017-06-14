package gosmparse

import "os"

// FeatureEnabled - return true if feature is enabled, else false
func FeatureEnabled(flag string) bool {
	if ff := os.Getenv(flag); ff != "" {
		return true
	}
	return false
}
