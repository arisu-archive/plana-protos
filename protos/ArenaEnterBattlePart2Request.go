package protos

type ArenaEnterBattlePart2Request struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ArenaBattleDB ArenaBattleDB `json:",omitempty,omitzero"`
}
