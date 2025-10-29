package protos

type ClanSettingResponse struct {
	ResponsePacket
	ClanDB ClanDB `json:",omitempty,omitzero"`
}
