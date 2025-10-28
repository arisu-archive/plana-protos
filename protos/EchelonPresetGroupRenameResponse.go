package protos

type EchelonPresetGroupRenameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetGroupDB EchelonPresetGroupDB `json:",omitempty,omitzero"`
}
