/*
Unleash API

Testing TagsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	openapiclient "github.com/Unleash/unleash-server-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_client_TagsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsApiService AddTagToFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId string

		httpRes, err := apiClient.TagsApi.AddTagToFeatures(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService CreateTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsApi.CreateTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService CreateTagType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsApi.CreateTagType(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService DeleteTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var type_ string
		var value string

		httpRes, err := apiClient.TagsApi.DeleteTag(context.Background(), type_, value).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService DeleteTagType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.TagsApi.DeleteTagType(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var type_ string
		var value string

		resp, httpRes, err := apiClient.TagsApi.GetTag(context.Background(), type_, value).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTagType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.TagsApi.GetTagType(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTagTypes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsApi.GetTagTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsApi.GetTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService GetTagsByType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.TagsApi.GetTagsByType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService UpdateTagType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		httpRes, err := apiClient.TagsApi.UpdateTagType(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsApiService ValidateTagType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagsApi.ValidateTagType(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}