package ddd

import "encoding/json"

type Metadata struct {
	AggregateName    string `json:"aggregate_name"`
	AggregateID      string `json:"aggregate_id"`
	AggregateVersion int64  `json:"aggregate_version"`
}

func (m Metadata) configureEvent(e *event) error {

	metadataBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	e.metadata = metadataBytes

	return nil
}
