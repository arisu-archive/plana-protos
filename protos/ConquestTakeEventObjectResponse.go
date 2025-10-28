package protos

type ConquestTakeEventObjectResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper ConquestEventObjectDB `json:",omitempty,omitzero"`
}
