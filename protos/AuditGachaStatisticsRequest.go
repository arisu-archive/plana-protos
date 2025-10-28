package protos

type AuditGachaStatisticsRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MerchandiseUniqueId int64 `json:",omitempty,omitzero"`
	ShopUniqueId int64 `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
