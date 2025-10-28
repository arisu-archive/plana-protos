package protos

type MiniGameShootingBattleEnterResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
