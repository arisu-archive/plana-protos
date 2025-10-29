package protos

type CraftUpdateNodeLevelResponse struct {
	ResponsePacket
	CraftInfoDB CraftInfoDB `json:",omitempty,omitzero"`
	CraftNodeDB CraftNodeDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
}
