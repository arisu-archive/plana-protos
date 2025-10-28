package protos

type ClanSearchResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDBs []ClanDB `json:",omitempty,omitzero"`
}
