package protos

type ShopGachaRecruitListResponse struct {
	ResponsePacket
	ShopRecruits              []ShopRecruitDB            `json:",omitempty,omitzero"`
	ShopFreeRecruitHistoryDBs []ShopFreeRecruitHistoryDB `json:",omitempty,omitzero"`
}
