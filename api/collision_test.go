package api

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNoPathCollisionWithDashboard asserts that no api route
// paths collide with the agent-browser dashboard's paths
func TestNoPathCollisionWithDashboard(t *testing.T) {
	dashboardPrefixes := []string{
		"/api/",
		"/_next/",
		"/favicon",
	}

	spec, err := GetSpec()
	require.NoError(t, err)

	for path := range spec.Paths.Map() {
		for _, dashboardPrefix := range dashboardPrefixes {
			require.Falsef(
				t,
				strings.HasPrefix(path, dashboardPrefix),
				"api path %q collides with dashboard path %q",
				path, dashboardPrefix,
			)
		}
	}
}
