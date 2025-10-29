package protos

type EchelonPresetListResponse struct {
	ResponsePacket
	PresetGroupDBs []EchelonPresetGroupDB `json:",omitempty,omitzero"`
}
