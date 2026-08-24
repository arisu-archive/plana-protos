package protos

type WelcomeCampaignMissionRewardRequest struct {
	RequestPacket
	SeasonId        int64 `json:",omitempty,omitzero"`
	MissionUniqueId int64 `json:",omitempty,omitzero"`
}
