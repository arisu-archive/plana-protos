package protos

type ConquestEventObjectBattleResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB `json:",omitempty,omitzero"`
	ConquestInfoDB ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestTileDB ConquestTileDB `json:",omitempty,omitzero"`
	DisplayInfos []ConquestDisplayInfo `json:",omitempty,omitzero"`
}
