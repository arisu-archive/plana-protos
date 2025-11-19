package protos

type EventContentConcentrationSaveDB struct {
	FlipCount int32                             `json:",omitempty,omitzero"`
	Round     int32                             `json:",omitempty,omitzero"`
	CardDBs   []EventContentConcentrationCardDB `json:",omitempty,omitzero"`
}
