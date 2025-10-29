package protos

type RaidGetBestTeamResponse struct {
	ResponsePacket
	RaidTeamSettingDBs []RaidTeamSettingDB `json:",omitempty,omitzero"`
}
