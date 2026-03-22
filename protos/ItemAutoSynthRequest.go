package protos

type ItemAutoSynthRequest struct {
	RequestPacket
	TargetParcels []ParcelKeyPair
}
