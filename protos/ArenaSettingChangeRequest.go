package protos

type ArenaSettingChangeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MapId int64 `json:",omitempty,omitzero"`
}
