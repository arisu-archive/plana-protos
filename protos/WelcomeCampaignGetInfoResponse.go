package protos

type WelcomeCampaignGetInfoResponse struct {
	ResponsePacket
	WelcomeCampaignInfo *WelcomeCampaignDB `json:",omitempty,omitzero"`
}
