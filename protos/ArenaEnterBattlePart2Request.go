package protos

type ArenaEnterBattlePart2Request struct {
	RequestPacket
	ArenaBattleDB ArenaBattleDB `json:",omitempty,omitzero"`
}
