package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type AttendanceBookReward struct {
	UniqueId          int64                        `json:",omitempty,omitzero"`
	Type              flatdata.AttendanceType      `json:",omitempty,omitzero"`
	AccountType       flatdata.AccountState        `json:",omitempty,omitzero"`
	DisplayOrder      int64                        `json:",omitempty,omitzero"`
	AccountLevelLimit int64                        `json:",omitempty,omitzero"`
	Title             string                       `json:",omitempty,omitzero"`
	TitleImagePath    string                       `json:",omitempty,omitzero"`
	CountRule         flatdata.AttendanceCountRule `json:",omitempty,omitzero"`
	CountReset        flatdata.AttendanceResetType `json:",omitempty,omitzero"`
	BookSize          int64                        `json:",omitempty,omitzero"`
	StartDate         MxTime                       `json:",omitempty,omitzero"`
	StartableEndDate  MxTime                       `json:",omitempty,omitzero"`
	EndDate           MxTime                       `json:",omitempty,omitzero"`
	ExpiryDate        int64                        `json:",omitempty,omitzero"`
	MailType          flatdata.MailType            `json:",omitempty,omitzero"`
	DailyRewardIcons  map[int64]string             `json:",omitempty,omitzero"`
	DailyRewards      map[int64][]ParcelInfo       `json:",omitempty,omitzero"`
}
