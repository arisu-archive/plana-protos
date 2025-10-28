package protos

type EventContentCardShopShuffleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CardShopElementDBs []CardShopElementDB `json:",omitempty,omitzero"`
}
