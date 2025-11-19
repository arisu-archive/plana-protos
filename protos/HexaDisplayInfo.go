package protos

type HexaDisplayInfo struct {
	Type            HexaDisplayType         `json:",omitempty,omitzero"`
	EntityId        int64                   `json:",omitempty,omitzero"`
	UniqueId        int64                   `json:",omitempty,omitzero"`
	Location        HexLocation             `json:",omitempty,omitzero"`
	Parameter       int64                   `json:",omitempty,omitzero"`
	StageRewardInfo StrategyClearRewardInfo `json:",omitempty,omitzero"`
}
