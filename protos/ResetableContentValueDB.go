package protos

type ResetableContentValueDB struct {
	ResetableContentId ResetableContentId
	ContentValue       int64  `json:",omitempty,omitzero"`
	LastUpdateTime     MxTime `json:",omitempty,omitzero"`
}
