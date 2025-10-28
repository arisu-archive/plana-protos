package protos

type CampaignListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
