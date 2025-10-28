package protos

type ArenaSettingChangeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
