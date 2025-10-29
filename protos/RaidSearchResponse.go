package protos

type RaidSearchResponse struct {
	ResponsePacket
	RaidDBs []RaidDB `json:",omitempty,omitzero"`
}
