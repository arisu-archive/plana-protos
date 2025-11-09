package protos

type ConquestErosionBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               ParcelResultDB          `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB `json:",omitempty,omitzero"`
	ConquestInfoDB               ConquestInfoDB          `json:",omitempty,omitzero"`
	DisplayInfos                 []ConquestDisplayInfo   `json:",omitempty,omitzero"`
}
