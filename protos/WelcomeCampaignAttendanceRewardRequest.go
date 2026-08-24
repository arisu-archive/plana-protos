package protos

type WelcomeCampaignAttendanceRewardRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
