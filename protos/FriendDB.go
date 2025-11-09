package protos

type FriendDB struct {
	AccountId                   int64               `json:",omitempty,omitzero"`
	Level                       int32               `json:",omitempty,omitzero"`
	Nickname                    string              `json:",omitempty,omitzero"`
	LastConnectTime             MxTime              `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId  int64               `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId int64               `json:",omitempty,omitzero"`
	ComfortValue                int64               `json:",omitempty,omitzero"`
	FriendCount                 int64               `json:",omitempty,omitzero"`
	AttachmentDB                AccountAttachmentDB `json:",omitempty,omitzero"`
}
