package main

import (
        // ... other imports
        "fmt"
        "./gemini" // Import the gemini package
)

func handler(rawEvt interface{}) {
        // ... (existing code)
        case *events.Message:
                // ... (existing code)
                if triggeredByGemini(evt.Message.GetText()) {
                        geminiResponse, err := gemini.QueryGeminiAPI(evt.Message.GetText(), "your_api_key") // Replace with your API key
                        if err != nil {
                                log.Errorf("Error querying Gemini API: %v", err)
                                return
                        }
                        replyMessage := extractReplyFromGeminiResponse(geminiResponse)
                        _, _ = cli.SendMessage(context.Background(), evt.Info.Sender, &waE2E.Message{Conversation: proto.String(replyMessage)})
                }
        }
}
