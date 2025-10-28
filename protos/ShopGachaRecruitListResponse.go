package protos

type ShopGachaRecruitListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopRecruits []ShopRecruitDB `json:",omitempty,omitzero"`
	ShopFreeRecruitHistoryDBs []ShopFreeRecruitHistoryDB `json:",omitempty,omitzero"`
}
