package protos

type CafeDB_CafeCharacterDB struct {
	VisitingCharacterDB
	IsSummon         bool   `json:",omitempty,omitzero"`
	LastInteractTime MxTime `json:",omitempty,omitzero"`
}
