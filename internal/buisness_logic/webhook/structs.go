package webhook

import "assignment-2/internal/structs"

// generateStruct generates a struct of type WebHookPost.
func generateStruct(webhookId string, country string, numberOfCalls int) structs.WebHookPost {
	return structs.WebHookPost{
		Id:      webhookId,
		Country: country,
		Calls:   numberOfCalls,
	}
}
