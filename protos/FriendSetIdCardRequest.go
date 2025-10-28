package protos

type FriendSetIdCardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Comment string `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId int64 `json:",omitempty,omitzero"`
	EmblemId int64 `json:",omitempty,omitzero"`
	SearchPermission bool `json:",omitempty,omitzero"`
	AutoAcceptFriendRequest bool `json:",omitempty,omitzero"`
	ShowAccountLevel bool `json:",omitempty,omitzero"`
	ShowFriendCode bool `json:",omitempty,omitzero"`
	ShowRaidRanking bool `json:",omitempty,omitzero"`
	ShowArenaRanking bool `json:",omitempty,omitzero"`
	ShowEliminateRaidRanking bool `json:",omitempty,omitzero"`
	BackgroundId int64 `json:",omitempty,omitzero"`
}
