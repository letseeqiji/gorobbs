package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// Session is an interface that stores values and configurations for a session.
type Session interface {
	Get(key interface{}) interface{}
	Set(key interface{}, val interface{})
	Delete(key interface{})
	Clear()
	AddFlash(value interface{}, vars ...string)
	Flashes(vars ...string) []interface{}
	Save() error
	Options(Options)
}

// Store is an interface for custom session stores.
type Store interface {
	sessions.Store
}

// Options stores configurations of a session
type Options struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

type session struct {
	name    string
	request *http.Request
	writer  http.ResponseWriter
	store   Store
	session *sessions.Session
}

func (s *session) Get(key interface{}) interface{} {
	return s.Session().Values[key]
}

func (s *session) Set(key interface{}, value interface{}) {
	s.Session().Values[key] = value
}

func (s *session) Delete(key interface{}) {
	delete(s.Session().Values, key)
}

func (s *session) Clear() {
	for key := range s.Session().Values {
		s.Delete(key)
	}
}

func (s *session) AddFlash(value interface{}, vars ...string) {
	s.Session().AddFlash(value, vars...)
}

func (s *session) Flashes(vars ...string) []interface{} {
	return s.Session().Flashes(vars...)
}

func (s *session) Save() error {
	return s.Session().Save(s.request, s.writer)
}

func (s *session) Session() *sessions.Session {
	if s.session == nil {
		if session, err := s.store.Get(s.request, s.name); err != nil {
			panic(err)
		} else {
			s.session = session
		}
	}

	return s.session
}

func (s *session) Options(options Options) {
	s.Session().Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
