package protos

type CampaignEnterTacticResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
