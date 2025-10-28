package protos

type MiniGameShootingBattleResultRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Summary MiniGameShootingSummary `json:",omitempty,omitzero"`
}
