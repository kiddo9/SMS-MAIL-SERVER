package structures


type MessageJob struct {
	Name         string
	PendingPrice string
	Course       string
	Phone        string
	Email        string
	Data         string
	Method       []string
	EmailId      int
	SmsId        int
	Admin       AdminStructs
}
