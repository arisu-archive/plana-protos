package protos

type EventContentWithdrawEchelonResponse struct {
	ResponsePacket
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	WithdrawEchelonDBs []EchelonDB `json:",omitempty,omitzero"`
}
