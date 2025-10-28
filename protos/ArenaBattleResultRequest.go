package protos

type ArenaBattleResultRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ArenaBattleDB ArenaBattleDB `json:",omitempty,omitzero"`
}
