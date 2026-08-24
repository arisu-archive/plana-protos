package protos

type WelcomeCampaignReceiveEnterRewardRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
