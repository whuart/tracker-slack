package main

import (
	"os"
	"strings"
)

// resolveChannel returns the Slack channel for an event type and environment.
// It looks up TRACKER_<TYPE>_CHANNEL_<ENV> first and falls back to
// TRACKER_<TYPE>_CHANNEL when the environment-specific variable is unset or empty.
// eventType is the variable infix, e.g. "DEPLOYMENT", "OPERATION", "RPA_USAGE".
func resolveChannel(eventType, env string) string {
	base := "TRACKER_" + strings.ToUpper(eventType) + "_CHANNEL"
	if env = strings.ToUpper(strings.TrimSpace(env)); env != "" {
		if channel := os.Getenv(base + "_" + env); channel != "" {
			return channel
		}
	}
	return os.Getenv(base)
}
