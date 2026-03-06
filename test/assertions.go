package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"

	api "github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/require"
)

func requireNoError(t *testing.T, err error) {
	t.Helper()
	require.NoError(t, err, formatErrorForTestOutput(err))
}

func formatErrorForTestOutput(err error) string {
	if err == nil {
		return ""
	}

	var apiErr *api.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		status := strings.TrimSpace(apiErr.Error())
		body := strings.TrimSpace(string(apiErr.Body()))
		if body == "" {
			return status
		}

		var payload map[string]interface{}
		if json.Unmarshal(apiErr.Body(), &payload) == nil {
			if message, ok := payload["message"].(string); ok && message != "" {
				return fmt.Sprintf("%s: %s", status, message)
			}
		}

		return fmt.Sprintf("%s: %s", status, body)
	}

	return err.Error()
}
