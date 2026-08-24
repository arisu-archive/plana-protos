package protos

type WelcomeCampaignCheckResponse struct {
	ResponsePacket
	HasNotReceiveReward bool `json:",omitempty,omitzero"`
	HasCompleteMission  bool `json:",omitempty,omitzero"`
}
