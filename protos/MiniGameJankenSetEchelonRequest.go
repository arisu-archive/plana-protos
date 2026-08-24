package protos

type MiniGameJankenSetEchelonRequest struct {
	RequestPacket
	EventContentId   int64 `json:",omitempty,omitzero"`
	CharacterId      int64 `json:",omitempty,omitzero"`
	EquipmentGroupId int64 `json:",omitempty,omitzero"`
}
