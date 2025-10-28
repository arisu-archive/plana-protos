package protos

type ShopBeforehandGachaRunResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SelectGachaSnapshot BeforehandGachaSnapshotDB `json:",omitempty,omitzero"`
}
