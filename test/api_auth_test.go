/*
Unleash API

Testing AuthAPIService

*/

package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_client_AuthAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test AuthAPIService GetPermissions", func(t *testing.T) {
		if enterpriseEnvironmentAvailable() {
			resp, httpRes, err := apiClient.AuthAPI.GetPermissions(context.Background()).Execute()

			require.Nil(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			fmt.Sprintln(resp.Permissions)
			assert.NotEmpty(t, resp.Permissions)
			assert.NotEmpty(t, resp.Permissions.Root)
			assert.NotEmpty(t, resp.Permissions.Project)
			assert.NotEmpty(t, resp.Permissions.Environments)
		}
	})

}
