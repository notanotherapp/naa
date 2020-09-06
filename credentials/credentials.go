package credentials

/*
Get the client credentials for Secret Manager.
*/

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"encoding/json"
	"fmt"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

// --------------------------------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------------------------------
type Credentials struct {
	PubSubTopic string     `json:"pubsub_topic"`
	Clientid    int        `json:"clientid"`
	Siteid      int        `json:"locationid"`
	Providers   []Provider `json:"providers"`
}

// --------------------------------------------------------------------------------------------------------------
type Provider struct {
	Provider string `json:"provider"`
	Code     string `json:"cashless_code"`
	Url      string `json:"url"`
	Uname    string `json:"username"`
	Password string `json:"password"`
	APIKey   string `json:"apikey"`
}

// --------------------------------------------------------------------------------------------------------------

func GetCredentials(projectid string, clientid string, version int) (Credentials, error) {

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Println(fmt.Errorf("failed to create secretmanager client: %v", err))
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/%d", projectid, clientid, version),
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Println(fmt.Errorf("failed to access secret version: %v", err))
	}

	dbc := Credentials{}

	err = json.Unmarshal(result.Payload.Data, &dbc)
	return dbc, err

}

// --------------------------------------------------------------------------------------------------------------
