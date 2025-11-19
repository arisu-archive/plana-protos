package protos

type MiniGameHistoryDB struct {
	EventContentId   int64  `json:",omitempty,omitzero"`
	UniqueId         int64  `json:",omitempty,omitzero"`
	HighScore        int64  `json:",omitempty,omitzero"`
	AccumulatedScore int64  `json:",omitempty,omitzero"`
	ClearDate        MxTime `json:",omitempty,omitzero"`
	IsFullCombo      bool   `json:",omitempty,omitzero"`
}
