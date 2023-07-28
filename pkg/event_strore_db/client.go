package event_strore_db

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
)

func NewEventStoreDB(connectionString string) (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}
	return esdb.NewClient(settings)
}
