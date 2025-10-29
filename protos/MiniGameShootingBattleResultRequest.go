package protos

type MiniGameShootingBattleResultRequest struct {
	RequestPacket
	Summary MiniGameShootingSummary `json:",omitempty,omitzero"`
}
