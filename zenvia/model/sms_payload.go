package model

type SMSPatyload struct {
	From     string       `json:"from"`
	To       string       `json:"to"`
	Contents []SMSContent `json:"contents"`
}

func CreateSingleSMSPayload(from string, to string, content SMSContent) SMSPatyload {
	return SMSPatyload{
		From: from,
		To:   to,
		Contents: []SMSContent{
			content,
		},
	}
}
