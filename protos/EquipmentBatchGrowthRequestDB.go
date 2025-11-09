package protos

type EquipmentBatchGrowthRequestDB struct {
	TargetServerId    int64                     `json:",omitempty,omitzero"`
	ConsumeRequestDBs []ConsumeRequestDB        `json:",omitempty,omitzero"`
	AfterTier         int64                     `json:",omitempty,omitzero"`
	AfterLevel        int64                     `json:",omitempty,omitzero"`
	AfterExp          int64                     `json:",omitempty,omitzero"`
	ReplaceInfos      []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
