package protos

type MiniGameCCGEndStageEventResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
	SaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
}
