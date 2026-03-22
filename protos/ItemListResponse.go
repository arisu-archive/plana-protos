package protos

type ItemListResponse struct {
	ResponsePacket
	ItemDBs       []ItemDB
	ExpiryItemDBs []ItemDB
}
