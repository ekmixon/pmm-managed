package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	pmmapitests "github.com/Percona-Lab/pmm-api-tests"
)

func TestVersion(t *testing.T) {
	paths := []string{
		"managed/v1/version",
		"v1/version",
	}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			uri := pmmapitests.BaseURL.ResolveReference(&url.URL{
				Path: path,
			})

			t.Logf("URI: %s", uri)
			resp, err := http.Get(uri.String())
			require.NoError(t, err)
			defer resp.Body.Close() //nolint:errcheck
			b, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode, "response:\n%s", b)

			var res struct {
				Version string
			}
			err = json.Unmarshal(b, &res)
			require.NoError(t, err, "response:\n%s", b)

			assert.True(t, strings.HasPrefix(res.Version, "2.0.0-"), "version = %q should has suffix", res.Version)
		})
	}
}
