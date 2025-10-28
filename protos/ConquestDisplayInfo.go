package protos

type ConquestDisplayInfo struct {
	TriggerType ConquestTriggerType `json:",omitempty,omitzero"`
	Type ConquestDisplayType `json:",omitempty,omitzero"`
	EntityId int64 `json:",omitempty,omitzero"`
	TileUniqueId int64 `json:",omitempty,omitzero"`
	Parameter string `json:",omitempty,omitzero"`
	DisplayOrder int32 `json:",omitempty,omitzero"`
	DisplayOnce bool `json:",omitempty,omitzero"`
}
