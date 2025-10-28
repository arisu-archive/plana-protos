package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MinigameJudgeRecord struct {
	NoteIndex int32 `json:",omitempty,omitzero"`
	TimingError float32 `json:",omitempty,omitzero"`
	CurrentCombo int32 `json:",omitempty,omitzero"`
	JudgeGradeOfThisNote flatdata.JudgeGrade `json:",omitempty,omitzero"`
	IsFeverOn bool `json:",omitempty,omitzero"`
}
