package protos

type EliminateRaidResetRequest struct {
	RequestPacket
	RaidServerId int64 `json:",omitempty,omitzero"`
}
