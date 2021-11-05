package zenvia

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/booscaaa/zenvia-go-sdk/zenvia/model"
)

func init() {
	Client = &http.Client{}
}

type zenviaAccess struct {
	access ZenviaAccess
}

func Instance(access ZenviaAccess) ZenviaAccessRepository {
	return &zenviaAccess{
		access: access,
	}
}

func (zenvia zenviaAccess) SendSMS(message string) {
	urlString := zenvia.access.api + "/v2/channels/sms/messages"

	content := model.CreateSMSContent(message)
	smsPayload := model.CreateSingleSMSPayload(zenvia.access.from, zenvia.access.to, content)

	payload, _ := json.Marshal(smsPayload)

	r, _ := http.NewRequest(http.MethodPost, urlString, strings.NewReader(string(payload))) // URL-encoded payload
	r.Header.Add("X-API-TOKEN", zenvia.access.apiToken)
	r.Header.Add("Content-Type", "application/json")

	Client.Do(r)

	// if err != nil {
	// 	return nil, err
	// }

	// defer resp.Body.Close()

	// if resp.StatusCode == 200 {
	// 	zenviaAccessAuth, err := model.FromJsonZenviaAccessAuth(resp.Body)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return zenviaAccessAuth, nil
	// } else {
	// 	defaultError := errors.ParseDefaultError(resp.Body)

	// 	return nil, defaultError
	// }
}
