package idSDK

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type InnaticalID struct {
	clientID     string
	clientSecret string
	redirectURL  string

	_idURL       string
	_idServerURL string
}

var Client *InnaticalID

func New(clientID, clientSecret, redirectURL string) {
	Client = &InnaticalID{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
		_idURL:       "https://innatical.id",
		_idServerURL: "https://api.innatical.id",
	}
}

func (i *InnaticalID) SetIDURL(idURL string) {
	i._idURL = idURL
}

func (i *InnaticalID) SetIDServerURL(idServerURL string) {
	i._idServerURL = idServerURL
}

func CreateURL(scope string, state string) string {
	return fmt.Sprintf(
		"%s/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s&prompt=consent",
		Client._idURL,
		Client.clientID,
		url.QueryEscape(Client.redirectURL),
		scope,
		state,
	)
}

type tokenRequest struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	UserID       string `json:"user_id"`
}

func GetToken(code string) (*TokenResponse, error) {
	if Client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	data := tokenRequest{
		Code:         code,
		ClientID:     Client.clientID,
		ClientSecret: Client.clientSecret,
		RedirectURI:  Client.redirectURL,
		GrantType:    "authorization_code",
	}

	postBody, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(fmt.Sprintf("%s/oauth2/token", Client._idServerURL), "application/json", responseBody)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var tokenResponse TokenResponse

	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Verified bool     `json:"verified"`
	Flags    []string `json:"flags"`
}

func GetCurrentUser(token string) (*User, error) {
	if Client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}

	url := fmt.Sprintf("%s/users/@me", Client._idServerURL)
	preparedToken := fmt.Sprintf("Bearer %s", token)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request : ", err)
		return nil, err
	}

	req.Header.Add("Authorization", preparedToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Failure : ", err)
		return nil, err
	}

	var user User

	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("Failure : ", err)
		return nil, err
	}

	return &user, nil
}
