package protos

type ArenaBattleResultRequest struct {
	RequestPacket
	ArenaBattleDB ArenaBattleDB `json:",omitempty,omitzero"`
}
