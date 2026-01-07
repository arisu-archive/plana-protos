package protos

type MultiFloorRaidLoginResponse struct {
	ResponsePacket
	LastClearedDifficulty int64 `json:",omitempty,omitzero"`
}
