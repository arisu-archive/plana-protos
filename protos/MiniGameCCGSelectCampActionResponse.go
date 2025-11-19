package protos

type MiniGameCCGSelectCampActionResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
	SaveDB  MiniGameCCGSaveDB      `json:",omitempty,omitzero"`
}
