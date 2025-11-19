package protos

type ArenaUserDB struct {
	AccountServerId             int64               `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId  int64               `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId int64               `json:",omitempty,omitzero"`
	NickName                    string              `json:",omitempty,omitzero"`
	Rank                        int64               `json:",omitempty,omitzero"`
	Level                       int64               `json:",omitempty,omitzero"`
	Exp                         int64               `json:",omitempty,omitzero"`
	TeamSettingDB               ArenaTeamSettingDB  `json:",omitempty,omitzero"`
	AccountAttachmentDB         AccountAttachmentDB `json:",omitempty,omitzero"`
	IsAnonymous                 bool                `json:",omitempty,omitzero"`
	UserName                    string              `json:",omitempty,omitzero"`
}
