package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/oauth1/oauth"
	"github.com/jinzhu/gorm"
	"github.com/kittyguys/hash/api/common"
	"github.com/kittyguys/hash/api/config"
	"github.com/kittyguys/hash/api/model"
	"github.com/labstack/echo"
)

const (
	// session token key
	tokenKey  = "tokenKey"
	secretKey = "secretKey"
	// user info url
	userInfoURI = "https://api.twitter.com/1.1/account/verify_credentials.json"
	// twitter base uri
	baseURI = "https://twitter.com/"
)

// OAuth oauth
type OAuth struct {
	client oauth.Client
	db     *gorm.DB
	config *config.Config
}

// Response singin response
type Response struct {
	Status int    `json:"status"`
	URL    string `json:"url"`
}

// Response singin response
type credentials struct {
	Token  string
	Secret string
}

// NewOAuthHandler new oauth handler
func NewOAuthHandler(config *config.Config, db *gorm.DB) *OAuth {
	// new oauth setting from cnofig.yaml
	return &OAuth{
		client: oauth.Client{
			TemporaryCredentialRequestURI: config.Twitter.RequestURI,
			ResourceOwnerAuthorizationURI: config.Twitter.AuthorizationURI,
			TokenRequestURI:               config.Twitter.TokenRequestURI,
			Credentials: oauth.Credentials{
				Token:  config.Twitter.Token,
				Secret: config.Twitter.Secret,
			},
		},
		db:     db,
		config: config,
	}
}

// TwitterLogin login with twitter
func (o *OAuth) TwitterLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get twitter access token
		credentials, err := o.client.RequestTemporaryCredentials(nil, o.config.Twitter.CallbackURI, nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewError(common.ErrGetCredentials, err))
		}

		// return twitter authp page url
		return c.JSON(http.StatusOK, Response{http.StatusOK, o.client.AuthorizationURL(credentials, nil)})
	}
}

// TwitterCallback twitter callback endpoint
func (o *OAuth) TwitterCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		var credentials oauth.Credentials

		credentials.Token = c.QueryParam("oauth_token")
		credentials.Secret = c.QueryParam("oauth_verifier")

		accessCredentials, _, err := o.client.RequestToken(nil, &credentials, c.QueryParam("oauth_verifier"))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewError(common.ErrGetCredentials, err))
		}

		user, err := o.GetUserInfo(accessCredentials)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewError(common.ErrNotFoundUserInfo, err))
		}

		// TODO: userをDBに保存する処理を書く

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["hashID"] = user.Name
		claims["displayName"] = user.ScreenName
		claims["profileImageURL"] = user.Image
		claims["iat"] = time.Now()
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, _ := token.SignedString([]byte(Key))

		c.Response().Header().Set("Authorization", "Bearer "+tokenString)

		// redirect mypage
		return c.Redirect(http.StatusFound, "http://localhost:3000")
	}
}

// GetUserInfo get twitter user info
func (o OAuth) GetUserInfo(credentials *oauth.Credentials) (model.TwitterUser, error) {
	var user model.TwitterUser

	if err := o.APIGet(
		credentials,
		userInfoURI,
		url.Values{"include_entities": {"true"}},
		&user); err != nil {

		return user, fmt.Errorf("Error getting timeline, %s", err.Error())
	}

	return user, nil
}

// APIGet call get twitter api
func (o OAuth) APIGet(cred *oauth.Credentials, urlStr string, form url.Values, data interface{}) error {
	resp, err := o.client.Get(nil, cred, urlStr, form)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != http.StatusOK {
		p, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("get %s returned status %d, %s", resp.Request.URL, resp.StatusCode, p)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
