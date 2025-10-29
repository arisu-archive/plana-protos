package protos

type RaidDetailRequest struct {
	RequestPacket
	RaidServerId int64 `json:",omitempty,omitzero"`
	RaidUniqueId int64 `json:",omitempty,omitzero"`
}
