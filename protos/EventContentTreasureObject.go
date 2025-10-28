package protos

type EventContentTreasureObject struct {
	ServerId int64 `json:",omitempty,omitzero"`
	RewardId int64 `json:",omitempty,omitzero"`
	Rotation int32 `json:",omitempty,omitzero"`
	IsHiddenImage bool `json:",omitempty,omitzero"`
	Cells []EventContentTreasureCell `json:",omitempty,omitzero"`
}
