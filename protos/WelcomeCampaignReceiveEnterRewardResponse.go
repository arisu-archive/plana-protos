package protos

type WelcomeCampaignReceiveEnterRewardResponse struct {
	ResponsePacket
	WelcomeCampaignInfo *WelcomeCampaignDB `json:",omitempty,omitzero"`
	ParcelResult        *ParcelResultDB    `json:",omitempty,omitzero"`
}
