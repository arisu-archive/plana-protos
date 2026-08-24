package protos

type ShopRecruitMileageHistoryDB struct {
	MileageGroupId      int64 `json:",omitempty,omitzero"`
	RecruitMileageCount int64 `json:",omitempty,omitzero"`
}
