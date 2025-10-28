package protos

type ShopBeforehandGachaSaveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SelectGachaSnapshot BeforehandGachaSnapshotDB `json:",omitempty,omitzero"`
}
