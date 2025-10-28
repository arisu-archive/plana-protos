package protos

type MailReceiveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MailServerIds []int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	BattlePassInfoDBs []BattlePassInfoDB `json:",omitempty,omitzero"`
}
