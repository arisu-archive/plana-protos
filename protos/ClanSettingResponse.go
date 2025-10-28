package protos

type ClanSettingResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanDB ClanDB `json:",omitempty,omitzero"`
}
