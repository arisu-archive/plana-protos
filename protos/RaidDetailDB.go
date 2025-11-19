package protos

type RaidDetailDB struct {
	RaidUniqueId int64              `json:",omitempty,omitzero"`
	EndDate      MxTime             `json:",omitempty,omitzero"`
	DamageTable  []RaidPlayerInfoDB `json:",omitempty,omitzero"`
}
