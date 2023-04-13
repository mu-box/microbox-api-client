
package client

import (
	"net/url"
	"time"
)

type (

	// User represents a microbox user
	User struct {
		AuthenticationToken string    `json:"authentication_token"` //
		CreatedAt           time.Time `json:"created_at"`           //
		Email               string    `json:"email"`                //
		ID                  string    `json:"id"`                   //
		UpdatedAt           time.Time `json:"updated_at"`           //
		Username            string    `json:"username"`             //
	}
)

// GetAuthToken takes a userSlug and password to return a user's authentication
// token
func GetAuthToken(userSlug, password string) (*User, error) {

	//
	v := url.Values{}
	v.Set("id", userSlug)
	v.Add("password", password)

	// this path is used (vs restful) to avoid sending emails as part of the path
	path := APIURL + "/" + APIVersion + "/user_auth_token?" + v.Encode()

	var user User
	return &user, DoRawRequest(&user, "GET", path, nil, nil)
}
