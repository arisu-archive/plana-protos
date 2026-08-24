package protos

type MiniGameJankenSetEchelonResponse struct {
	ResponsePacket
	SaveDB *MiniGameJankenSaveDB `json:",omitempty,omitzero"`
}
