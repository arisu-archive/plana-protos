package protos

type AcademyGetInfoResponse struct {
	ResponsePacket
	AcademyDB          AcademyDB
	AcademyLocationDBs []AcademyLocationDB
}
