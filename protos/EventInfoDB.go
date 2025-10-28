package protos

type EventInfoDB struct {
	EventId int64 `json:",omitempty,omitzero"`
	ImageNameHash uint32 `json:",omitempty,omitzero"`
}
