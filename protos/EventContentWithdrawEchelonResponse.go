package protos

type EventContentWithdrawEchelonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB EventContentMainStageSaveDB `json:",omitempty,omitzero"`
	WithdrawEchelonDBs []EchelonDB `json:",omitempty,omitzero"`
}
