package protos

type AcademyLocationDB struct {
	LocationId int64 `json:",omitempty,omitzero"`
	Rank       int64 `json:",omitempty,omitzero"`
	Exp        int64 `json:",omitempty,omitzero"`
}
