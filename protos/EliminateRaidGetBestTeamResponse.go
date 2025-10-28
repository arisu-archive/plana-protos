package protos

type EliminateRaidGetBestTeamResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidTeamSettingDBsDict map[string][]RaidTeamSettingDB `json:",omitempty,omitzero"`
}
