package client

import (
	"context"
	"testing"

	"github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanupContextField(t *testing.T, apiClient *client.APIClient, name string) {
	_, err := apiClient.ContextAPI.DeleteContextField(context.Background(), name).Execute()
	require.Nil(t, err)
}

func fetchContextField(t *testing.T, apiClient *client.APIClient, name string) *client.ContextFieldSchema {
	response, _, err := apiClient.ContextAPI.GetContextField(context.Background(), name).Execute()
	require.Nil(t, err)
	require.NotNil(t, response)
	return response
}

func Test_client_ContextFieldAPIService(t *testing.T) {
	apiClient := testClient()

	t.Run("Test ContextFieldAPIService CreateContextField", func(t *testing.T) {
		brieLegalValue := *client.NewLegalValueSchemaWithDefaults()
		brieLegalValue.Value = "Brie"
		brieLegalValue.SetDescription("A soft cheese from France.")

		stiltonLegalValue := *client.NewLegalValueSchemaWithDefaults()
		stiltonLegalValue.Value = "Stilton"
		stiltonLegalValue.SetDescription("A blue cheese from England.")

		valuesRequest := []client.LegalValueSchema{brieLegalValue, stiltonLegalValue}

		createContextFieldRequest := *client.NewCreateContextFieldSchemaWithDefaults()
		createContextFieldRequest.Name = "Cheese"
		createContextFieldRequest.SetDescription("Constrain feature toggles on a required cheese type.")
		createContextFieldRequest.SetLegalValues(valuesRequest)

		resp, httpRes, err := apiClient.ContextAPI.CreateContextField(context.Background()).CreateContextFieldSchema(createContextFieldRequest).Execute()
		defer cleanupContextField(t, apiClient, "Cheese")

		if err != nil {
			t.Log(err)
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)
	})

	t.Run("Test ContextFieldAPIService UpdateContextField", func(t *testing.T) {

		targetedContextField := "Cheese"

		createContextFieldRequest := *client.NewCreateContextFieldSchemaWithDefaults()
		createContextFieldRequest.Name = targetedContextField
		createContextFieldRequest.SetDescription("Constrain feature toggles on a required cheese type.")

		var resp, httpRes, err = apiClient.ContextAPI.CreateContextField(context.Background()).CreateContextFieldSchema(createContextFieldRequest).Execute()
		defer cleanupContextField(t, apiClient, targetedContextField)
		if err != nil {
			t.Log(err)
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)

		fieldValue := fetchContextField(t, apiClient, targetedContextField)
		require.NotNil(t, fieldValue)
		assert.Equal(t, "Constrain feature toggles on a required cheese type.", *fieldValue.Description.Get())

		updateContextFieldRequest := *client.NewUpdateContextFieldSchemaWithDefaults()
		updateContextFieldRequest.SetDescription("Bounded by the enum values returned by the CheeseMongerAPI")

		httpRes, err = apiClient.ContextAPI.UpdateContextField(context.Background(), targetedContextField).UpdateContextFieldSchema(updateContextFieldRequest).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		if err != nil {
			t.Log(err)
		}

		updatedFieldValue := fetchContextField(t, apiClient, targetedContextField)
		assert.Equal(t, "Bounded by the enum values returned by the CheeseMongerAPI", *updatedFieldValue.Description.Get())
	})

	t.Run("Test ContextFieldAPIService UpdateLegalValues ", func(t *testing.T) {

		targetedContextField := "PoisonousMushrooms"

		flyAgaricValue := *client.NewLegalValueSchemaWithDefaults()
		flyAgaricValue.Value = "Fly agaric"
		flyAgaricValue.SetDescription("Beautiful, red with white spots.")

		valuesRequest := []client.LegalValueSchema{flyAgaricValue}

		createContextFieldRequest := *client.NewCreateContextFieldSchemaWithDefaults()
		createContextFieldRequest.Name = targetedContextField
		createContextFieldRequest.SetDescription("Mushrooms you probably shouldn't eat..")
		createContextFieldRequest.SetLegalValues(valuesRequest)

		resp, httpRes, err := apiClient.ContextAPI.CreateContextField(context.Background()).CreateContextFieldSchema(createContextFieldRequest).Execute()
		defer cleanupContextField(t, apiClient, targetedContextField)

		if err != nil {
			t.Log(err)
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 201, httpRes.StatusCode)

		deathCapValue := *client.NewLegalValueSchemaWithDefaults()
		deathCapValue.Value = "Death cap"
		deathCapValue.SetDescription("Pale white with a brownish top")

		valuesRequest = []client.LegalValueSchema{deathCapValue}

		updateContextFieldRequest := *client.NewUpdateContextFieldSchemaWithDefaults()
		updateContextFieldRequest.SetLegalValues(valuesRequest)

		httpRes, err = apiClient.ContextAPI.UpdateContextField(context.Background(), targetedContextField).UpdateContextFieldSchema(updateContextFieldRequest).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

		fieldValue := fetchContextField(t, apiClient, targetedContextField)
		require.NotNil(t, fieldValue)
		assert.Equal(t, "Death cap", fieldValue.LegalValues[0].Value)
		assert.Equal(t, "Pale white with a brownish top", *fieldValue.LegalValues[0].Description)
	})
}
