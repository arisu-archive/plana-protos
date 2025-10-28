package protos

type ContentSaveDiscardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
