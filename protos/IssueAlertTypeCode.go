package protos

type IssueAlertTypeCode int32

const (
	IssueAlertTypeCode_All                  IssueAlertTypeCode = 1
	IssueAlertTypeCode_File_Target          IssueAlertTypeCode = 2
	IssueAlertTypeCode_AllButFile_Exception IssueAlertTypeCode = 3
)
