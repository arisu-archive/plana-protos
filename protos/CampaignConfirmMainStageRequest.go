package protos

type CampaignConfirmMainStageRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
