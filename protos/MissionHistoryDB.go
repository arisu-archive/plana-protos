package protos

type MissionHistoryDB struct {
	MissionUniqueId int64  `json:",omitempty,omitzero"`
	CompleteTime    MxTime `json:",omitempty,omitzero"`
	Expired         bool   `json:",omitempty,omitzero"`
}
