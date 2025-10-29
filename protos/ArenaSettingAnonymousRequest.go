package protos

type ArenaSettingAnonymousRequest struct {
	RequestPacket
	IsAnonymous bool `json:",omitempty,omitzero"`
}
