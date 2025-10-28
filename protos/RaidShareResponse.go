package protos

type RaidShareResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidDB RaidDB `json:",omitempty,omitzero"`
}
