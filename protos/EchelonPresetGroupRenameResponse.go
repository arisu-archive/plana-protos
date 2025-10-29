package protos

type EchelonPresetGroupRenameResponse struct {
	ResponsePacket
	PresetGroupDB EchelonPresetGroupDB `json:",omitempty,omitzero"`
}
