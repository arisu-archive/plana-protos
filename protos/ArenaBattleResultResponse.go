package protos

type ArenaBattleResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
