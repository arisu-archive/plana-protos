package protos

type RaidOpponentListResponse struct {
	ResponsePacket
	OpponentUserDBs []SingleRaidUserDB `json:",omitempty,omitzero"`
}
