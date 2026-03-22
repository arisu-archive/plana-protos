package protos

type CharacterExpGrowthResponse struct {
	ResponsePacket
	CharacterDB       CharacterDB
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
}
