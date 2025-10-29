package protos

type CampaignEnterMainStageStrategySkipResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
