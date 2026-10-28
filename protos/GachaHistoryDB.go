package protos

type GachaHistoryDB struct {
	RecruitedAt   MxTime `json:",omitempty,omitzero"`
	ShopRecruitId int64  `json:",omitempty,omitzero"`
	RecruitCount  int32  `json:",omitempty,omitzero"`
	SeqInBatch    int32  `json:",omitempty,omitzero"`
	CharacterId   int64  `json:",omitempty,omitzero"`
}
