package protos

type MiniGameDefenseEnterBattleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
