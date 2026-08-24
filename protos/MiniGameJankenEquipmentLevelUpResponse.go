package protos

type MiniGameJankenEquipmentLevelUpResponse struct {
	ResponsePacket
	SaveDB         *MiniGameJankenSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB *ParcelResultDB       `json:",omitempty,omitzero"`
}
