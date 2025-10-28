package protos

type EventContentBoxGachaElement struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	VariationId int64 `json:",omitempty,omitzero"`
	Round int64 `json:",omitempty,omitzero"`
	GroupId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	IsPrize bool `json:",omitempty,omitzero"`
	Rewards []ParcelInfo `json:",omitempty,omitzero"`
}
