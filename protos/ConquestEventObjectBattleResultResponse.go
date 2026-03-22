package protos

type ConquestEventObjectBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               ParcelResultDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	ConquestInfoDB               ConquestInfoDB
	ConquestTileDB               ConquestTileDB
	DisplayInfos                 []ConquestDisplayInfo
}
