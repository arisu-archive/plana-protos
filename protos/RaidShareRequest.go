package protos

type RaidShareRequest struct {
	RequestPacket
	RaidServerId int64 `json:",omitempty,omitzero"`
}
