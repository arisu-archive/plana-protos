package protos

type EventContentPermanentDB struct {
	EventContentId            int64 `json:",omitempty,omitzero"`
	IsStageAllClear           bool  `json:",omitempty,omitzero"`
	IsReceivedCharacterReward bool  `json:",omitempty,omitzero"`
}
