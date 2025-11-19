package protos

type PickupFirstGetHistoryDB struct {
	ShopRecruitId int64  `json:",omitempty,omitzero"`
	LogDate       MxTime `json:",omitempty,omitzero"`
}
