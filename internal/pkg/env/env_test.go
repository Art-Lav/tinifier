package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, "TINYPNG_API_KEY", string(TinyPngAPIKey))
}

func TestEnvVariable_Lookup(t *testing.T) {
	cases := []struct {
		giveEnv envVariable
	}{
		{giveEnv: TinyPngAPIKey},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.giveEnv.String(), func(t *testing.T) {
			assert.NoError(t, os.Unsetenv(tt.giveEnv.String())) // make sure that env is unset for test

			defer func() { assert.NoError(t, os.Unsetenv(tt.giveEnv.String())) }()

			value, exists := tt.giveEnv.Lookup()
			assert.False(t, exists)
			assert.Empty(t, value)

			assert.NoError(t, os.Setenv(tt.giveEnv.String(), "foo"))

			value, exists = tt.giveEnv.Lookup()
			assert.True(t, exists)
			assert.Equal(t, "foo", value)
		})
	}
}
