package invoke_fg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// getTokenUserNamePassword retrieves an authentication token using username and password
func getTokenUserNamePassword(client http.Client, userName, userPassword, domainName, authURL, projectID string) (string, error) {
	tokenURI := authURL + "/auth/tokens?v3/auth/tokens?nocatalog=true"

	// Build auth request body
	authBody := map[string]interface{}{
		"auth": map[string]interface{}{
			"identity": map[string]interface{}{
				"methods": []string{"password"},
				"password": map[string]interface{}{
					"user": map[string]interface{}{
						"name":     userName,
						"password": userPassword,
						"domain": map[string]interface{}{
							"name": domainName,
						},
					},
				},
			},
			"scope": map[string]interface{}{
				"domain": map[string]interface{}{
					"name": domainName,
				},
				"project": map[string]interface{}{
					"id": projectID,
				},
			},
		},
	}

	requestBody, err := json.Marshal(authBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling auth body: %w", err)
	}

	req, err := http.NewRequest("POST", tokenURI, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting token: %w", err)
	}
	defer resp.Body.Close()

	token := resp.Header.Get("X-Subject-Token")

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseText, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error getting token: %s\n", string(responseText))
		return token, fmt.Errorf("token request failed with status: %d", resp.StatusCode)
	}

	return token, nil
}
