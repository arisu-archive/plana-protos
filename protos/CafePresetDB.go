package protos

type CafePresetDB struct {
	ServerId   int64  `json:",omitempty,omitzero"`
	SlotId     int32  `json:",omitempty,omitzero"`
	PresetName string `json:",omitempty,omitzero"`
	IsEmpty    bool   `json:",omitempty,omitzero"`
}
