package protos

type ClanKickHistoryCancelRequest struct {
	RequestPacket
	TargetAccountId int64 `json:",omitempty,omitzero"`
}
