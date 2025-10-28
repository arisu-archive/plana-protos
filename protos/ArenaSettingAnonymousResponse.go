package protos

type ArenaSettingAnonymousResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
