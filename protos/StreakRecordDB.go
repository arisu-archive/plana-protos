package protos

type StreakRecordDB struct {
	SeasonId           int64 `json:",omitempty,omitzero"`
	Step               int32 `json:",omitempty,omitzero"`
	Day                int32 `json:",omitempty,omitzero"`
	PeakStreakStep     int32 `json:",omitempty,omitzero"`
	PeakStreakDay      int32 `json:",omitempty,omitzero"`
	FreeRewardStep     int32 `json:",omitempty,omitzero"`
	FreeRewardDay      int32 `json:",omitempty,omitzero"`
	PurchaseRewardStep int32 `json:",omitempty,omitzero"`
	PurchaseRewardDay  int32 `json:",omitempty,omitzero"`
	ProductId          int64 `json:",omitempty,omitzero"`
	AbsentDay          int32 `json:",omitempty,omitzero"`
}
