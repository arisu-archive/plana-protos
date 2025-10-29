package protos

type ArenaSettingChangeRequest struct {
	RequestPacket
	MapId int64 `json:",omitempty,omitzero"`
}
