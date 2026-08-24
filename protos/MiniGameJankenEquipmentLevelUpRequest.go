package protos

type MiniGameJankenEquipmentLevelUpRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
