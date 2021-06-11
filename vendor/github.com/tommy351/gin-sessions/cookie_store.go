package sessions

import "github.com/gorilla/sessions"

// CookieStore stores sessions using secure cookies.
type CookieStore interface {
	Store
	Options(Options)
}

// NewCookieStore creates a new CookieStore.
func NewCookieStore(keyPairs ...[]byte) CookieStore {
	return &cookieStore{sessions.NewCookieStore(keyPairs...)}
}

type cookieStore struct {
	*sessions.CookieStore
}

func (c *cookieStore) Options(options Options) {
	c.CookieStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
