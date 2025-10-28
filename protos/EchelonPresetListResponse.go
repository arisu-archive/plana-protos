package protos

type EchelonPresetListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetGroupDBs []EchelonPresetGroupDB `json:",omitempty,omitzero"`
}
