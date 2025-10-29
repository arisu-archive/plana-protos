package protos

type RaidShareResponse struct {
	ResponsePacket
	RaidDB RaidDB `json:",omitempty,omitzero"`
}
