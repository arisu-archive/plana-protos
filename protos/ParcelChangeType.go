package protos

type ParcelChangeType int32

const (
	ParcelChangeType_NoChange   ParcelChangeType = 0
	ParcelChangeType_Terminated ParcelChangeType = 1
	ParcelChangeType_MailSend   ParcelChangeType = 2
	ParcelChangeType_Converted  ParcelChangeType = 3
)
