package protos

type EventContentChangeDB struct {
	EventContentId        int64  `json:",omitempty,omitzero"`
	UseAmount             int64  `json:",omitempty,omitzero"`
	ChangeCount           int64  `json:",omitempty,omitzero"`
	AccumulateChangeCount int64  `json:",omitempty,omitzero"`
	LastUpdateDate        MxTime `json:",omitempty,omitzero"`
	ChangeFlag            bool   `json:",omitempty,omitzero"`
}
