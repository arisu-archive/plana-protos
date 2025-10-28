package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type MissionInfo struct {
	IMissionConstraint
	Id int64 `json:",omitempty,omitzero"`
	Category flatdata.MissionCategory `json:",omitempty,omitzero"`
	ResetType flatdata.MissionResetType `json:",omitempty,omitzero"`
	ToastDisplayType flatdata.MissionToastDisplayConditionType `json:",omitempty,omitzero"`
	Description uint32 `json:",omitempty,omitzero"`
	IsVisible bool `json:",omitempty,omitzero"`
	IsLimited bool `json:",omitempty,omitzero"`
	StartDate time.Time `json:",omitempty,omitzero"`
	StartableEndDate time.Time `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
	EndDday int64 `json:",omitempty,omitzero"`
	AccountState flatdata.AccountState `json:",omitempty,omitzero"`
	AccountLevel int64 `json:",omitempty,omitzero"`
	PreMissionIds []int64 `json:",omitempty,omitzero"`
	NextMissionId int64 `json:",omitempty,omitzero"`
	SuddenMissionContentTypes []flatdata.SuddenMissionContentType `json:",omitempty,omitzero"`
	CompleteConditionType flatdata.MissionCompleteConditionType `json:",omitempty,omitzero"`
	CompleteConditionCount int64 `json:",omitempty,omitzero"`
	CompleteConditionParameters []int64 `json:",omitempty,omitzero"`
	Tags []flatdata.Tag `json:",omitempty,omitzero"`
	CompleteConditionMissionIds []int64 `json:",omitempty,omitzero"`
	CompleteConditionMissionCount int64 `json:",omitempty,omitzero"`
	CompleteConditionRewards []ParcelInfo `json:",omitempty,omitzero"`
	RewardIcon string `json:",omitempty,omitzero"`
	Rewards []ParcelInfo `json:",omitempty,omitzero"`
	DateAutoRefer flatdata.ContentType `json:",omitempty,omitzero"`
	ToastImagePath string `json:",omitempty,omitzero"`
	DisplayOrder int64 `json:",omitempty,omitzero"`
	HasFollowingMission bool `json:",omitempty,omitzero"`
	Shortcuts []string `json:",omitempty,omitzero"`
	ChallengeStageId int64 `json:",omitempty,omitzero"`
}
