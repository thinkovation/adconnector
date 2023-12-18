package adapi

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

type Adapi struct {
	clientSecretCredential *azidentity.ClientSecretCredential
	appClient              *msgraphsdk.GraphServiceClient
}

// NewADHelperWithClientSecret creates a new ADHelper with client secret authentication.
// It assumes that an app has been registered in the tenant with the given client ID and client secret.
// Many of the examples do a call to os.getEnv("AZURE_TENANT_ID") to get the tenant ID etc, within these functions
// but I prefer to pass the values in on initialisation so that the code is more self-contained.

func NewADHelperWithClientSecret(tenantID, clientID, clientSecret string) (*Adapi, error) {
	h := &Adapi{}

	credential, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		return nil, err
	}

	h.clientSecretCredential = credential

	// Create an auth provider using the credential
	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(h.clientSecretCredential, []string{
		"https://graph.microsoft.com/.default",
	})
	if err != nil {
		return nil, err
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
		return nil, err
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)
	h.appClient = client
	return h, nil
}

func (g *Adapi) GetFilteredUsers(filter string) (models.UserCollectionResponseable, error) {
	var topValue int32 = 25
	var filterValue *string
	if filter != "" {
		filterValue = &filter
	}
	query := users.UsersRequestBuilderGetQueryParameters{
		// Only request specific properties
		Select: []string{"displayName", "id", "mail", "jobTitle", "department"},
		// Get at most 25 results
		Top: &topValue,
		// Sort by display name
		//Orderby: []string{"displayName"},
		Filter: filterValue,
	}
	return g.appClient.Users().
		Get(context.Background(),
			&users.UsersRequestBuilderGetRequestConfiguration{
				QueryParameters: &query,
			})
}

func (g *Adapi) GetUsers() (models.UserCollectionResponseable, error) {
	var topValue int32 = 25
	query := users.UsersRequestBuilderGetQueryParameters{
		// Only request specific properties
		Select: []string{"displayName", "id", "mail", "jobTitle", "department"},
		// Get at most 25 results
		Top: &topValue,
		// Sort by display name
		Orderby: []string{"displayName"},
	}

	return g.appClient.Users().
		Get(context.Background(),
			&users.UsersRequestBuilderGetRequestConfiguration{
				QueryParameters: &query,
			})
}
