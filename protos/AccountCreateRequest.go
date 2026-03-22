package protos

type AccountCreateRequest struct {
	RequestPacket
	DevId           string
	Version         int64 `json:",omitempty,omitzero"`
	IMEI            int64 `json:",omitempty,omitzero"`
	AccessIP        string
	MarketId        string
	UserType        string
	AdvertisementId string
	OSType          string
	OSVersion       string
}
