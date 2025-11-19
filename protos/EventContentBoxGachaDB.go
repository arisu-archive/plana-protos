package protos

type EventContentBoxGachaDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	Seed           int64 `json:",omitempty,omitzero"`
	Round          int64 `json:",omitempty,omitzero"`
	PurchaseCount  int32 `json:",omitempty,omitzero"`
}
