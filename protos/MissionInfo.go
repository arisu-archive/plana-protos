package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MissionInfo struct {
	Id                            int64                                     `json:",omitempty,omitzero"`
	Category                      flatdata.MissionCategory                  `json:",omitempty,omitzero"`
	ResetType                     flatdata.MissionResetType                 `json:",omitempty,omitzero"`
	ToastDisplayType              flatdata.MissionToastDisplayConditionType `json:",omitempty,omitzero"`
	Description                   uint32                                    `json:",omitempty,omitzero"`
	IsVisible                     bool                                      `json:",omitempty,omitzero"`
	IsLimited                     bool                                      `json:",omitempty,omitzero"`
	StartDate                     MxTime                                    `json:",omitempty,omitzero"`
	StartableEndDate              MxTime                                    `json:",omitempty,omitzero"`
	EndDate                       MxTime                                    `json:",omitempty,omitzero"`
	EndDday                       int64                                     `json:",omitempty,omitzero"`
	TargetGroup                   flatdata.TargetGroup                      `json:",omitempty,omitzero"`
	AccountLevel                  int64                                     `json:",omitempty,omitzero"`
	PreMissionIds                 []int64
	NextMissionId                 int64 `json:",omitempty,omitzero"`
	SuddenMissionContentTypes     []flatdata.SuddenMissionContentType
	CompleteConditionType         flatdata.MissionCompleteConditionType `json:",omitempty,omitzero"`
	CompleteConditionCount        int64                                 `json:",omitempty,omitzero"`
	CompleteConditionParameters   []int64
	Tags                          []flatdata.Tag
	CompleteConditionMissionIds   []int64
	CompleteConditionMissionCount int64 `json:",omitempty,omitzero"`
	CompleteConditionRewards      []ParcelInfo
	RewardIcon                    string
	Rewards                       []ParcelInfo
	DateAutoRefer                 flatdata.ContentType
	ToastImagePath                string
	DisplayOrder                  int64 `json:",omitempty,omitzero"`
	HasFollowingMission           bool  `json:",omitempty,omitzero"`
	Shortcuts                     []string
	ChallengeStageId              int64 `json:",omitempty,omitzero"`
}
