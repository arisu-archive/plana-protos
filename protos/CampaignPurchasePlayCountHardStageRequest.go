package protos

type CampaignPurchasePlayCountHardStageRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
