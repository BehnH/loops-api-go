/*
Loops OpenAPI Spec

Testing MailingListsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package loops

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/behnh/loops-api-go"
)

func Test_loops_MailingListsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MailingListsAPIService ListsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MailingListsAPI.ListsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
