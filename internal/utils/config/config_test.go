package config

import "testing"

func TestParseConfigJson(t *testing.T) {
	tests := []struct {
		Name    string
		Data    string
		WantErr bool
	}{{
		Name: "valid json",
		Data: `{ 
		"validApiKeys": {
			"abc123": "dev",
			"def456": "prod"
			}
		}`,
		WantErr: false},
		{
			Name: "invalid json",
			Data: `{ 
		"validApiKeys": {
			"abc1
			}
		}`,
			WantErr: true}}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := ParseConfig([]byte(tt.Data))
			if err != nil && !tt.WantErr {
				t.Errorf("Failed to parse json")
			}
			if err == nil && tt.WantErr {
				t.Errorf("Failed to trigger error")
			}

		})

	}
}
