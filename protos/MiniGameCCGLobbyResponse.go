package protos

type MiniGameCCGLobbyResponse struct {
	ResponsePacket
	CCGSaveDB   MiniGameCCGSaveDB `json:",omitempty,omitzero"`
	Perks       []int64           `json:",omitempty,omitzero"`
	RewardPoint int32             `json:",omitempty,omitzero"`
	CanSweep    bool              `json:",omitempty,omitzero"`
}
