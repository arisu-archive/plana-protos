package protos

type ClanSetAssistResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanAssistSlotDB ClanAssistSlotDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	RewardInfo ClanAssistRewardInfo `json:",omitempty,omitzero"`
}
