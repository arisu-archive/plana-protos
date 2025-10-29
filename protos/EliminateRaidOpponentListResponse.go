package protos

type EliminateRaidOpponentListResponse struct {
	ResponsePacket
	OpponentUserDBs []EliminateRaidUserDB `json:",omitempty,omitzero"`
}
