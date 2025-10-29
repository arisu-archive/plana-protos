package protos

type ClanSearchResponse struct {
	ResponsePacket
	ClanDBs []ClanDB `json:",omitempty,omitzero"`
}
