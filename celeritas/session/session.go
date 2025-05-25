package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60 // Default to 60 minutes if conversion fails
	}

	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	}

	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.Secure = secure
	session.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(c.SessionType) {
	case "redis":
	case "mysql, mariadb":
	case "postgresql, postgres":
	default:
	}

	return session
}
