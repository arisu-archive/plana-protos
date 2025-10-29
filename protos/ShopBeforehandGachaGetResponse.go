package protos

type ShopBeforehandGachaGetResponse struct {
	ResponsePacket
	AlreadyPicked bool `json:",omitempty,omitzero"`
	BeforehandGachaSnapshot BeforehandGachaSnapshotDB `json:",omitempty,omitzero"`
}
