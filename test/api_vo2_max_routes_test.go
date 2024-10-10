/*
Oura API Documentation

Testing VO2MaxRoutesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package oura

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_oura_VO2MaxRoutesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VO2MaxRoutesAPIService MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VO2MaxRoutesAPI.MultipleVO2MaxDocumentsV2UsercollectionVO2MaxGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VO2MaxRoutesAPIService SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var documentId string

		resp, httpRes, err := apiClient.VO2MaxRoutesAPI.SingleVO2MaxDocumentV2UsercollectionVO2MaxDocumentIdGet(context.Background(), documentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
