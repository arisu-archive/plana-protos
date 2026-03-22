package protos

type CraftUpdateNodeLevelResponse struct {
	ResponsePacket
	CraftInfoDB       CraftInfoDB
	CraftNodeDB       CraftNodeDB
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
}
