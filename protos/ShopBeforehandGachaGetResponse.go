package protos

type ShopBeforehandGachaGetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AlreadyPicked bool `json:",omitempty,omitzero"`
	BeforehandGachaSnapshot BeforehandGachaSnapshotDB `json:",omitempty,omitzero"`
}
