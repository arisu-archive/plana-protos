package protos

type WelcomeCampaignCheckRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
