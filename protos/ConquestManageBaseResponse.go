package protos

type ConquestManageBaseResponse struct {
	ResponsePacket
	ClearParcels                 [][]*ParcelInfo
	ConquerBonusParcels          [][]*ParcelInfo
	BonusParcels                 []*ParcelInfo
	ParcelResultDB               *ParcelResultDB `json:",omitempty,omitzero"`
	ConquestInfoDB               *ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []*ConquestDisplayInfo
}
