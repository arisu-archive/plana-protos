package protos

type ConquestEventObjectBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               *ParcelResultDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	ConquestInfoDB               *ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestTileDB               *ConquestTileDB `json:",omitempty,omitzero"`
	DisplayInfos                 []ConquestDisplayInfo
}
