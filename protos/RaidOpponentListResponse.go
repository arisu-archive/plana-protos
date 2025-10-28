package protos

type RaidOpponentListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	OpponentUserDBs []SingleRaidUserDB `json:",omitempty,omitzero"`
}
