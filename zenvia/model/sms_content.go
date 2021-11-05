package model

type SMSContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func CreateSMSContent(tx string) SMSContent {
	return SMSContent{
		Type: "text",
		Text: tx,
	}
}
