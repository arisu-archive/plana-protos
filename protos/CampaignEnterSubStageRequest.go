package protos

type CampaignEnterSubStageRequest struct {
	RequestPacket
	StageUniqueId               int64 `json:",omitempty,omitzero"`
	LastEnterStageEchelonNumber int64 `json:",omitempty,omitzero"`
}
