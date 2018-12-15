package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cases := []struct {
		name     string
		path     string
		wantData *Configuration
		wantErr  bool
	}{
		{
			name:    "Fail on non-existing file",
			path:    "notExists",
			wantErr: true,
		},
		{
			name:    "Fail on wrong file format",
			path:    "test/config.invalid.yaml",
			wantErr: true,
		},
		{
			name: "Success in loading Configuration",
			path: "test/config.testdata.yaml",
			wantData: &Configuration{
				DB: &Database{
					DSN:        "postgres://postgres:postgres@postgres",
					LogQueries: true,
					Timeout:    20,
				},
				Server: &Server{
					Port:         ":8080",
					Debug:        true,
					ReadTimeout:  15,
					WriteTimeout: 20,
				},
				JWT: &JWT{
					Secret:           "testing",
					Duration:         10,
					RefreshDuration:  10,
					MaxRefresh:       144,
					SigningAlgorithm: "HS384",
				},
				App: &Application{
					MinPasswordStr: 3,
					SwaggerUIPath:  "assets/swagger",
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.path)
			assert.Equal(t, tt.wantData, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
