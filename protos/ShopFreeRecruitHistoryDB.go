package protos

type ShopFreeRecruitHistoryDB struct {
	UniqueId       int64  `json:",omitempty,omitzero"`
	RecruitCount   int32  `json:",omitempty,omitzero"`
	LastUpdateDate MxTime `json:",omitempty,omitzero"`
}
