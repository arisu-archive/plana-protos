package protos

type ConquestErosionBattleStartResponse struct {
	ResponsePacket
	ParcelResultDB      ParcelResultDB      `json:",omitempty,omitzero"`
	ConquestStageSaveDB ConquestStageSaveDB `json:",omitempty,omitzero"`
}
