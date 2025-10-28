package protos

type ShopGachaRecruitListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
