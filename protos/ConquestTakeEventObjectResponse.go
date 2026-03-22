package protos

type ConquestTakeEventObjectResponse struct {
	ResponsePacket
	ParcelResultDB               ParcelResultDB
	ConquestEventObjectDBWrapper ConquestEventObjectDB
}
