package middleware

import "github.com/labstack/echo/v4"

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		_, _, _ = c.Get("pk"), c.Get("msg"), c.Get("sig")

		// verify.

		return nil
	}
}
