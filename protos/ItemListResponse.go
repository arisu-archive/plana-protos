package protos

type ItemListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ItemDBs []ItemDB `json:",omitempty,omitzero"`
	ExpiryItemDBs []ItemDB `json:",omitempty,omitzero"`
}
