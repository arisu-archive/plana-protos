package protos

type RaidResetResponse struct {
	ResponsePacket
	RaidDB *RaidDB `json:",omitempty,omitzero"`
}
