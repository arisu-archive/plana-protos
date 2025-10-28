package protos

type ArenaSettingAnonymousRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IsAnonymous bool `json:",omitempty,omitzero"`
}
