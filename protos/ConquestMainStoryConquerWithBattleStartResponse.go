package protos

type ConquestMainStoryConquerWithBattleStartResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestStageSaveDB ConquestStageSaveDB `json:",omitempty,omitzero"`
}
