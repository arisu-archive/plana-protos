package protos

type GachaSimulateCheatResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CharacterIdAndCount map[int64]int32 `json:",omitempty,omitzero"`
	SimulationCount int64 `json:",omitempty,omitzero"`
	GoodsUniqueId int64 `json:",omitempty,omitzero"`
	GoodsDevName string `json:",omitempty,omitzero"`
}
