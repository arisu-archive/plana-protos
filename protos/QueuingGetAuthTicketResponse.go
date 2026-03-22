package protos

type QueuingGetAuthTicketResponse struct {
	ResponsePacket
	EncryptedKey              string
	SignedKey                 string
	EncryptedIV               string
	SignedIV                  string
	EncryptedSqlCipherKey     string
	EncryptedSqlCipherLicense string
	Birth                     string
	AuthTicket                string
}
