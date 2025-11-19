package protos

type ShopRecruitDB struct {
	Id             int64  `json:",omitempty,omitzero"`
	SalesStartDate MxTime `json:",omitempty,omitzero"`
	SalesEndDate   MxTime `json:",omitempty,omitzero"`
	UpdateDate     MxTime `json:",omitempty,omitzero"`
}
