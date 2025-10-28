package protos

type GetArenaTeamCheatResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Opponent ArenaUserDB `json:",omitempty,omitzero"`
}
