package datastore

import (
	"os"
    "context"
    "go.mercari.io/datastore"
    "go.mercari.io/datastore/clouddatastore"
)

var client *datastore.Client

func Init() {
	ProjectId := os.Getenv("GCP_PROJECT")
    ctx := context.Background()
    _client, err := clouddatastore.FromContext(
        ctx,
        datastore.WithProjectID(ProjectId),
    )
    if err != nil {
        panic(err)
    }
    client = &_client
}

func GetConnection() datastore.Client {
	return *client
}
