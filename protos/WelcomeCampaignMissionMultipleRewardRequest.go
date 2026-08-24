package protos

type WelcomeCampaignMissionMultipleRewardRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
