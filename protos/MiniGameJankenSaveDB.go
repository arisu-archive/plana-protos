package protos

type MiniGameJankenSaveDB struct {
	EventContentId          int64 `json:",omitempty,omitzero"`
	CharacterId             int64 `json:",omitempty,omitzero"`
	EquipmentGroupId        int64 `json:",omitempty,omitzero"`
	EquipmentTier           int32 `json:",omitempty,omitzero"`
	CurrentStageId          int64 `json:",omitempty,omitzero"`
	CumulatedScore          int64 `json:",omitempty,omitzero"`
	LastReceivedRewardScore int64 `json:",omitempty,omitzero"`
}
