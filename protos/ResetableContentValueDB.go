package protos

type ResetableContentValueDB struct {
	ResetableContentId ResetableContentId `json:",omitempty,omitzero"`
	ContentValue       int64              `json:",omitempty,omitzero"`
	LastUpdateTime     MxTime             `json:",omitempty,omitzero"`
}
