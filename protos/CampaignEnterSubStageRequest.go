package protos

type CampaignEnterSubStageRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	LastEnterStageEchelonNumber int64 `json:",omitempty,omitzero"`
}
