//-- Package Declaration -----------------------------------------------------------------------------------------------
package authorization

//-- Imports -----------------------------------------------------------------------------------------------------------
import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/allegro/bigcache"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//-- Constants ---------------------------------------------------------------------------------------------------------
var (
	secret = []byte(`secret`)

	users = map[string]string{
		`user@example.com`: `password`,
	}
) //TODO: I should not exist! Kill me!

var (
	jwtExpire   = 5 * time.Minute
	jtiExpire   = 24 * time.Hour
	jwtStore, _ = bigcache.NewBigCache(bigcache.DefaultConfig(jtiExpire))

	AuthenticationMiddleware = middleware.JWT(secret)
)

//-- Structs -----------------------------------------------------------------------------------------------------------
type claims struct {
	Scopes     []string `json:"scopes,omitempty"`                 
	Authorized bool     `json:"admin,omitempty"`

	jwt.StandardClaims
}

type token struct {
	Token string `json:"token"`
}

type authorization struct {
	User     string `json:"email"formdata:"spaghetti"`
	Password string `json:"password"`
}

//-- Exported Functions ------------------------------------------------------------------------------------------------
func Create(context echo.Context) error {
	//-- Parse ----------
	var auth = new(authorization)
	if err := context.Bind(auth); err != nil {
		return err
	}

	//-- Authentication ----------
	if password, found := users[auth.User]; found && password == auth.Password {
		//Create Claims
		var claims = &claims{
			Authorized: true,
			StandardClaims: jwt.StandardClaims{
				Id:        uuid.New().String(),
				Subject:   auth.User,
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(jwtExpire).Unix(),
			},
		}

		// Encode, sign and respond with token
		if output, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret); err != nil {
			return err
		} else if encoded, err := json.Marshal(claims); err != nil {
			return err
		} else if err := jwtStore.Set(claims.Id, encoded); err != nil {
			return err
		} else {
			return context.JSON(http.StatusOK, token{output})
		}
	}

	//-- Return Default Response ----------
	return echo.ErrUnauthorized
}

func Renew(context echo.Context) error {
	//-- Shim ----------
	//jwt.context.Request().Header.Get(echo.HeaderAuthorization)
	var jti = context.Get(`user`).(*jwt.Token).Claims.(jwt.MapClaims)[`jti`].(string)

	//-- Authentication ----------
	var claims = &claims{}

	if encoded, err := jwtStore.Get(jti); err != nil {
		return err
	} else if err := json.Unmarshal([]byte(encoded), claims); err != nil {
		return err
	} else {
		//Update Claims
		claims.ExpiresAt = time.Now().Add(jwtExpire).Unix()

		// Encode, sign and respond with token
		if output, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret); err != nil {
			return err
		} else if encoded, err := json.Marshal(claims); err != nil {
			return err
		} else if err := jwtStore.Set(claims.Id, encoded); err != nil {
			return err
		} else {
			return context.JSON(http.StatusOK, token{output})
		}
	}

	//-- Return Default Response ----------
	return echo.ErrUnauthorized
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
