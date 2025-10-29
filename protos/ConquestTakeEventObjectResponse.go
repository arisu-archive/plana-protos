package protos

type ConquestTakeEventObjectResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper ConquestEventObjectDB `json:",omitempty,omitzero"`
}
