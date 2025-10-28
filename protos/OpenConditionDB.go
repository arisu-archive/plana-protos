package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type OpenConditionDB struct {
	ContentType flatdata.OpenConditionContent `json:",omitempty,omitzero"`
	HideWhenLocked bool `json:",omitempty,omitzero"`
	AccountLevel int64 `json:",omitempty,omitzero"`
	ScenarioModeId int64 `json:",omitempty,omitzero"`
	CampaignStageUniqueId int64 `json:",omitempty,omitzero"`
	MultipleConditionCheckType flatdata.MultipleConditionCheckType `json:",omitempty,omitzero"`
	OpenDayOfWeek flatdata.WeekDay `json:",omitempty,omitzero"`
	OpenHour int64 `json:",omitempty,omitzero"`
	CloseDayOfWeek flatdata.WeekDay `json:",omitempty,omitzero"`
	CloseHour int64 `json:",omitempty,omitzero"`
	CafeIdForCafeRank int64 `json:",omitempty,omitzero"`
	CafeRank int64 `json:",omitempty,omitzero"`
	OpenedCafeId int64 `json:",omitempty,omitzero"`
}
