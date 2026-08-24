package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WelcomeCampaignDB struct {
	SeasonId                      int64                `json:",omitempty,omitzero"`
	AttendanceCheckType           flatdata.TargetGroup `json:",omitempty,omitzero"`
	StartDate                     MxTime               `json:",omitempty,omitzero"`
	EndDate                       MxTime               `json:",omitempty,omitzero"`
	CompletedDate                 *MxTime              `json:",omitempty,omitzero"`
	WelcomeRewardReceivedDate     *MxTime              `json:",omitempty,omitzero"`
	CumulativeAttendanceDays      int16                `json:",omitempty,omitzero"`
	CumulativeRewardReceivedDay   int16                `json:",omitempty,omitzero"`
	LastAttendanceDate            *MxTime              `json:",omitempty,omitzero"`
	ConsecutiveAttendanceDays     int16                `json:",omitempty,omitzero"`
	ConsecutiveRewardReceivedDate *MxTime              `json:",omitempty,omitzero"`
}
