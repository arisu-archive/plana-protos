package protos

type WelcomeCampaignMissionListRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
