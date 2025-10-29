package protos

type ShopBeforehandGachaRunResponse struct {
	ResponsePacket
	SelectGachaSnapshot BeforehandGachaSnapshotDB `json:",omitempty,omitzero"`
}
