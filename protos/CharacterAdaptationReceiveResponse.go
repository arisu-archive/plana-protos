package protos

type CharacterAdaptationReceiveResponse struct {
	ResponsePacket
	CharacterAdaptationDB *CharacterAdaptationDB `json:",omitempty,omitzero"`
	ParcelResult          *ParcelResultDB        `json:",omitempty,omitzero"`
	ReturnParcels         []*ParcelInfo
}
