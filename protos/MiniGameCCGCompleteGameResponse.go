package protos

type MiniGameCCGCompleteGameResponse struct {
	ResponsePacket
	OldSaveDB      MiniGameCCGSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB    `json:",omitempty,omitzero"`
	RewardParcels  []ParcelInfo      `json:",omitempty,omitzero"`
}
