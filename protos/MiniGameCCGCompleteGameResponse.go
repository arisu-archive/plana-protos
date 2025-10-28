package protos

type MiniGameCCGCompleteGameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	OldSaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	RewardParcels []ParcelInfo `json:",omitempty,omitzero"`
}
