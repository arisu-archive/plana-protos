package protos

type ArenaEnterBattlePart1Response struct {
	ResponsePacket
	ArenaBattleDB ArenaBattleDB `json:",omitempty,omitzero"`
}
