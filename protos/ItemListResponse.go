package protos

type ItemListResponse struct {
	ResponsePacket
	ItemDBs []ItemDB `json:",omitempty,omitzero"`
	ExpiryItemDBs []ItemDB `json:",omitempty,omitzero"`
}
