package protos

type AuditGachaStatisticsRequest struct {
	RequestPacket
	MerchandiseUniqueId int64 `json:",omitempty,omitzero"`
	ShopUniqueId        int64 `json:",omitempty,omitzero"`
	Count               int64 `json:",omitempty,omitzero"`
}
