package protos

type ScenarioCollectionDB struct {
	GroupId     int64  `json:",omitempty,omitzero"`
	UniqueId    int64  `json:",omitempty,omitzero"`
	ReceiveDate MxTime `json:",omitempty,omitzero"`
}
