package protos

type EliminateRaidOpponentListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	OpponentUserDBs []EliminateRaidUserDB `json:",omitempty,omitzero"`
}
