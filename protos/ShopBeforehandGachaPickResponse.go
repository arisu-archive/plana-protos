package protos

type ShopBeforehandGachaPickResponse struct {
	ResponsePacket
	GachaResults  []GachaResult
	AcquiredItems []ItemDB
}
