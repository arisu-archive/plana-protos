package protos

type ConquestEventObjectBattleStartResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestStageSaveDB ConquestStageSaveDB `json:",omitempty,omitzero"`
}
