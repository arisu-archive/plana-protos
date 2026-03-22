package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type AttendanceBookReward struct {
	UniqueId          int64                   `json:",omitempty,omitzero"`
	Type              flatdata.AttendanceType `json:",omitempty,omitzero"`
	TargetGroup       flatdata.TargetGroup    `json:",omitempty,omitzero"`
	DisplayOrder      int64                   `json:",omitempty,omitzero"`
	AccountLevelLimit int64                   `json:",omitempty,omitzero"`
	Title             string
	TitleImagePath    string
	CountRule         flatdata.AttendanceCountRule `json:",omitempty,omitzero"`
	CountReset        flatdata.AttendanceResetType `json:",omitempty,omitzero"`
	BookSize          int64                        `json:",omitempty,omitzero"`
	StartDate         MxTime                       `json:",omitempty,omitzero"`
	StartableEndDate  MxTime                       `json:",omitempty,omitzero"`
	EndDate           MxTime                       `json:",omitempty,omitzero"`
	ExpiryDate        int64                        `json:",omitempty,omitzero"`
	MailType          flatdata.MailType            `json:",omitempty,omitzero"`
	DailyRewardIcons  *mapx.OrderedMap[int64, string]
	DailyRewards      *mapx.OrderedMap[int64, []ParcelInfo]
}
