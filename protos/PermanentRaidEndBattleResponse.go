package protos

type PermanentRaidEndBattleResponse struct {
	ResponsePacket
	ScoreInfo RaidScoreInfo `json:",omitempty,omitzero"`
}
