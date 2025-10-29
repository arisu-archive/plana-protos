package protos

type ShopBeforehandGachaSaveResponse struct {
	ResponsePacket
	SelectGachaSnapshot BeforehandGachaSnapshotDB `json:",omitempty,omitzero"`
}
