package protos

type AccountLimitedGachaDB struct {
	Id        int64  `json:",omitempty,omitzero"`
	StartDate MxTime `json:",omitempty,omitzero"`
	EndDate   MxTime `json:",omitempty,omitzero"`
}
