package protos

type MiniGameJankenScoreRewardResponse struct {
	ResponsePacket
	SaveDB         *MiniGameJankenSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB *ParcelResultDB       `json:",omitempty,omitzero"`
}
