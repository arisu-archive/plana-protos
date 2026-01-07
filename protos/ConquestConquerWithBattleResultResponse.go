package protos

type ConquestConquerWithBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               ParcelResultDB          `json:",omitempty,omitzero"`
	ConquestTileDB               ConquestTileDB          `json:",omitempty,omitzero"`
	ConquestInfoDB               ConquestInfoDB          `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB `json:",omitempty,omitzero"`
	DisplayInfos                 []ConquestDisplayInfo   `json:",omitempty,omitzero"`
	StepAfterBattle              int32                   `json:",omitempty,omitzero"`
	DisplayParcelByRewardTag     map[string][]ParcelInfo `json:",omitempty,omitzero"`
}
