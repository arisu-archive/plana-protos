package protos

type AccountRestrictionsDB struct {
	NicknameRestriction bool `json:",omitempty,omitzero"`
	CommentRestriction  bool `json:",omitempty,omitzero"`
	CallnameRestriction bool `json:",omitempty,omitzero"`
}
