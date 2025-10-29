package protos

type ConquestConquerWithBattleStartResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestStageSaveDB ConquestStageSaveDB `json:",omitempty,omitzero"`
}
