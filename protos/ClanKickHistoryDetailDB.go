package protos

type ClanKickHistoryDetailDB struct {
	AccountId                  int64  `json:",omitempty,omitzero"`
	KickDate                   MxTime `json:",omitempty,omitzero"`
	LastLoginDate              MxTime `json:",omitempty,omitzero"`
	AccountLevel               int64  `json:",omitempty,omitzero"`
	NickName                   string `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId int64  `json:",omitempty,omitzero"`
	EmblemUniqueId             int64  `json:",omitempty,omitzero"`
	StudentFrameUniqueId       int64  `json:",omitempty,omitzero"`
}
