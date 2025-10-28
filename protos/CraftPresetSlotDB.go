package protos

type CraftPresetSlotDB struct {
	PresetNodeDBs []CraftPresetNodeDB `json:",omitempty,omitzero"`
}
