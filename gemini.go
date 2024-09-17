package gemini

import (
        "github.com/google/generative-ai-go"
        "bytes"
        "encoding/json"
        "fmt"
        "net/http"
)

// QueryGeminiAPI function to interact with the Gemini API
func QueryGeminiAPI(text string, apiKey string) (string, error) {
        // API endpoint
        url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent"

        // Prepare request body
        requestBody := map[string]interface{}{
                "prompt": text,
                // Add other parameters as needed based on the API
        }
        jsonBody, err := json.Marshal(requestBody)
        if err != nil {
                return "", fmt.Errorf("error marshalling request body: %w", err)
        }

        // Create HTTP request
        req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
        if err != nil {
                return "", fmt.Errorf("error creating HTTP request: %w", err)
        }

        // Set headers
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", "Bearer "+apiKey) // Example for API key authentication

        // Send request
        client := http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                return "", fmt.Errorf("error sending HTTP request: %w", err)
        }
        defer resp.Body.Close()

        // Parse response
        var responseBody map[string]interface{}
        err = json.NewDecoder(resp.Body).Decode(&responseBody)
        if err != nil {
                return "", fmt.Errorf("error decoding response body: %w", err)
        }

        // Extract desired information from response
        generatedText, ok := responseBody["generated_text"].(string)
        if !ok {
                return "", fmt.Errorf("response does not contain generated text")
        }

        return generatedText, nil
}
