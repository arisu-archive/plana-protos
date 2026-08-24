package protos

type WelcomeCampaignGetInfoRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
