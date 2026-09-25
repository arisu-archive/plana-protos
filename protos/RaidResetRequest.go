package protos

type RaidResetRequest struct {
	RequestPacket
	RaidServerId int64 `json:",omitempty,omitzero"`
}
