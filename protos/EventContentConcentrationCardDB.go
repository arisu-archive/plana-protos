package protos

type EventContentConcentrationCardDB struct {
	Index     int32 `json:",omitempty,omitzero"`
	CardId    int64 `json:",omitempty,omitzero"`
	IsMatched bool  `json:",omitempty,omitzero"`
}
