package protos

type ConquestUpgradeBaseResponse struct {
	ResponsePacket
	UpgradeRewards               []ParcelInfo
	ParcelResultDB               ParcelResultDB
	ConquestTileDB               ConquestTileDB
	ConquestInfoDB               ConquestInfoDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []ConquestDisplayInfo
}
