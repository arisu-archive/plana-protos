package protos

type RaidGetBestTeamResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidTeamSettingDBs []RaidTeamSettingDB `json:",omitempty,omitzero"`
}
