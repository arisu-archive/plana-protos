package protos

type BattlePassCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	HasNotReceiveReward bool `json:",omitempty,omitzero"`
	HasCompleteMission bool `json:",omitempty,omitzero"`
}
