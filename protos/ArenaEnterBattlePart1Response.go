package protos

type ArenaEnterBattlePart1Response struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ArenaBattleDB ArenaBattleDB `json:",omitempty,omitzero"`
}
