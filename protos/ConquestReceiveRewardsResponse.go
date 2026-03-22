package protos

type ConquestReceiveRewardsResponse struct {
	ResponsePacket
	ParcelResultDB  ParcelResultDB
	ConquestInfoDB  ConquestInfoDB
	ConquestTileDBs []ConquestTileDB
}
