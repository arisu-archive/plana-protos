package protos

type GetArenaTeamCheatResponse struct {
	ResponsePacket
	Opponent ArenaUserDB `json:",omitempty,omitzero"`
}
