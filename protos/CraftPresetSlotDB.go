package protos

type CraftPresetSlotDB struct {
	PresetIndex   int32               `json:",omitempty,omitzero"`
	PresetNodeDBs []CraftPresetNodeDB `json:",omitempty,omitzero"`
	PresetName    string              `json:",omitempty,omitzero"`
}
