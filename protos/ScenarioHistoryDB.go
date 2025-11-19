package protos

type ScenarioHistoryDB struct {
	ScenarioUniqueId int64  `json:",omitempty,omitzero"`
	ClearDateTime    MxTime `json:",omitempty,omitzero"`
}
