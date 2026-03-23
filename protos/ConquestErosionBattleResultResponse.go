package protos

type ConquestErosionBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               *ParcelResultDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	ConquestInfoDB               *ConquestInfoDB `json:",omitempty,omitzero"`
	DisplayInfos                 []ConquestDisplayInfo
}
