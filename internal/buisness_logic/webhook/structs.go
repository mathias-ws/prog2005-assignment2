package webhook

// webHookPost contains the data that is sent when a webhook is triggered.
type webHookPost struct {
	Id      string `json:"webhook_id"`
	Country string `json:"country"`
	Calls   int    `json:"calls"`
}

// generateStruct generates a struct of type webHookPost.
func generateStruct(webhookId string, country string, numberOfCalls int) webHookPost {
	return webHookPost{
		Id:      webhookId,
		Country: country,
		Calls:   numberOfCalls,
	}
}
