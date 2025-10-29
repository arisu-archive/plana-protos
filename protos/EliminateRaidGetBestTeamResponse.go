package protos

type EliminateRaidGetBestTeamResponse struct {
	ResponsePacket
	RaidTeamSettingDBsDict map[string][]RaidTeamSettingDB `json:",omitempty,omitzero"`
}
