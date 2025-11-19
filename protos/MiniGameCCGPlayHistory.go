package protos

type MiniGameCCGPlayHistory struct {
	LevelId     int64                  `json:",omitempty,omitzero"`
	NodeId      int64                  `json:",omitempty,omitzero"`
	StagePlayDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
}
