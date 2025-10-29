package protos

type CampaignEnterMainStageRequest struct {
	RequestPacket
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
