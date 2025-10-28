package protos

type ConquestManageBaseResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClearParcels [][]ParcelInfo `json:",omitempty,omitzero"`
	ConquerBonusParcels [][]ParcelInfo `json:",omitempty,omitzero"`
	BonusParcels []ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConquestInfoDB ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB `json:",omitempty,omitzero"`
	DisplayInfos []ConquestDisplayInfo `json:",omitempty,omitzero"`
}
