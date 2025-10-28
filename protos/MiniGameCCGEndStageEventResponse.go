package protos

type MiniGameCCGEndStageEventResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
	SaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
}
