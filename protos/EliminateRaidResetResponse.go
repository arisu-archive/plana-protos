package protos

type EliminateRaidResetResponse struct {
	ResponsePacket
	RaidDB *RaidDB `json:",omitempty,omitzero"`
}
