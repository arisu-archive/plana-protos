package protos

type MiniGameJankenEnterBattleResponse struct {
	ResponsePacket
	SaveDB *MiniGameJankenSaveDB `json:",omitempty,omitzero"`
}
