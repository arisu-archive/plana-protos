package protos

type WelcomeCampaignAttendanceRewardResponse struct {
	ResponsePacket
	WelcomeCampaignInfo *WelcomeCampaignDB `json:",omitempty,omitzero"`
	ParcelResult        *ParcelResultDB    `json:",omitempty,omitzero"`
}
