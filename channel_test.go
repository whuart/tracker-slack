package main

import (
	"strings"
	"testing"
)

func TestResolveChannel(t *testing.T) {
	tests := []struct {
		name      string
		env       map[string]string
		eventType string
		trackEnv  string
		want      string
	}{
		{
			name:      "env-specific channel",
			env:       map[string]string{"TRACKER_DEPLOYMENT_CHANNEL": "default", "TRACKER_DEPLOYMENT_CHANNEL_PROD": "prod-chan"},
			eventType: "DEPLOYMENT",
			trackEnv:  "PROD",
			want:      "prod-chan",
		},
		{
			name:      "fallback when env-specific unset",
			env:       map[string]string{"TRACKER_INCIDENT_CHANNEL": "default", "TRACKER_INCIDENT_CHANNEL_PROD": "prod-chan"},
			eventType: "INCIDENT",
			trackEnv:  "DEV",
			want:      "default",
		},
		{
			name:      "fallback when env-specific empty",
			env:       map[string]string{"TRACKER_DRIFT_CHANNEL": "default", "TRACKER_DRIFT_CHANNEL_UAT": ""},
			eventType: "DRIFT",
			trackEnv:  "UAT",
			want:      "default",
		},
		{
			name:      "fallback when environment empty",
			env:       map[string]string{"TRACKER_OPERATION_CHANNEL": "default"},
			eventType: "OPERATION",
			trackEnv:  "",
			want:      "default",
		},
		{
			name:      "case insensitive environment and type with underscore",
			env:       map[string]string{"TRACKER_RPA_USAGE_CHANNEL": "default", "TRACKER_RPA_USAGE_CHANNEL_PREP": "prep-chan"},
			eventType: "rpa_usage",
			trackEnv:  "prep",
			want:      "prep-chan",
		},
		{
			name:      "nothing configured",
			env:       map[string]string{},
			eventType: "DEPLOYMENT",
			trackEnv:  "PROD",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := "TRACKER_" + strings.ToUpper(tt.eventType) + "_CHANNEL"
			for _, suffix := range []string{"", "_PROD", "_PREP", "_UAT", "_DEV"} {
				t.Setenv(base+suffix, "")
			}
			for k, v := range tt.env {
				t.Setenv(k, v)
			}
			if got := resolveChannel(tt.eventType, tt.trackEnv); got != tt.want {
				t.Errorf("resolveChannel(%q, %q) = %q, want %q", tt.eventType, tt.trackEnv, got, tt.want)
			}
		})
	}
}
