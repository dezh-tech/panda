package middleware

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/dezh-tech/panda/pkg"
	"github.com/dezh-tech/panda/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/nbd-wtf/go-nostr"
)

const oneMinute = 60 * time.Second

func Auth(url string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Response().Header().Get("Authorization")
			if len(token) <= 6 {
				return echo.NewHTTPError(http.StatusUnauthorized, pkg.ResponseDto{
					Success: false,
					Error:   validator.Varror{Error: echo.ErrUnauthorized.Error()},
				})
			}

			data, err := base64.RawStdEncoding.DecodeString(token[:6])
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, pkg.ResponseDto{
					Success: false,
					Error:   validator.Varror{Error: echo.ErrUnauthorized.Error()},
				})
			}

			event := new(nostr.Event)
			if err := event.UnmarshalJSON(data); err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, pkg.ResponseDto{
					Success: false,
					Error:   validator.Varror{Error: echo.ErrUnauthorized.Error()},
				})
			}

			if !CheckAuthEvent(event, url) {
				return echo.NewHTTPError(http.StatusUnauthorized, pkg.ResponseDto{
					Success: false,
					Error:   validator.Varror{Error: echo.ErrUnauthorized.Error()},
				})
			}

			c.Set("pubkey", event.PubKey)

			return next(c)
		}
	}
}

func CheckAuthEvent(e *nostr.Event, url string) bool {
	if len(e.Tags) != 2 {
		return false
	}

	if isValid, err := e.CheckSignature(); !isValid || err != nil {
		return false
	}

	diff := time.Until(e.CreatedAt.Time())
	if !(diff <= oneMinute && diff >= -oneMinute) {
		return false
	}

	expirationStr := e.Tags.GetFirst([]string{"expiration"}).Value()
	expirationInt, err := strconv.ParseInt(expirationStr, 10, 64)
	if err != nil {
		return false
	}

	expiration := time.Unix(expirationInt, 0)
	if !(expiration.After(time.Now().UTC()) && time.Until(expiration) >= 10*time.Minute) {
		return false
	}

	return e.Tags.GetFirst([]string{"u"}).Value() == url
}
