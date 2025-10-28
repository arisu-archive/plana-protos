package protos

type RaidSearchResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidDBs []RaidDB `json:",omitempty,omitzero"`
}
