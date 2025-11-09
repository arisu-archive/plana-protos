package protos

type ItemSelectTicketRequest struct {
	RequestPacket
	TicketItemServerId int64 `json:",omitempty,omitzero"`
	SelectItemUniqueId int64 `json:",omitempty,omitzero"`
	ConsumeCount       int32 `json:",omitempty,omitzero"`
}
