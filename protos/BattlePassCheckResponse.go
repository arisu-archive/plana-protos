package protos

type BattlePassCheckResponse struct {
	ResponsePacket
	HasNotReceiveReward bool `json:",omitempty,omitzero"`
	HasCompleteMission bool `json:",omitempty,omitzero"`
}
