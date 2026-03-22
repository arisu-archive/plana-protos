package protos

type ConquestErosionBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               ParcelResultDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	ConquestInfoDB               ConquestInfoDB
	DisplayInfos                 []ConquestDisplayInfo
}
