package protos

type ScenarioGroupHistoryDB struct {
	ScenarioGroupUniqueId int64  `json:",omitempty,omitzero"`
	ScenarioType          int64  `json:",omitempty,omitzero"`
	EventContentId        *int64 `json:",omitempty,omitzero"`
	ClearDateTime         MxTime `json:",omitempty,omitzero"`
	IsReturn              bool   `json:",omitempty,omitzero"`
}
