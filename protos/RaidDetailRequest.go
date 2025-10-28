package protos

type RaidDetailRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidServerId int64 `json:",omitempty,omitzero"`
	RaidUniqueId int64 `json:",omitempty,omitzero"`
}
