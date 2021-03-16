package session

import (
	"alwayslate/security"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

const authCookieKey = "auth"

type Authenticator struct {
	store sessions.Store
}

func New() *Authenticator {
	authKey := securecookie.GenerateRandomKey(64)
	encryptionKey :=  securecookie.GenerateRandomKey(32)

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options = &sessions.Options{
		MaxAge: 60 * 20, // Seconds
		HttpOnly: true,
	}

	return &Authenticator{store: store}
}

func (a *Authenticator) Authenticate(r *http.Request, w http.ResponseWriter, u security.User) error {
	session, _ := a.store.Get(r, authCookieKey)
	session.Values["user"] = u

	if err := session.Save(r, w); err != nil {
		return security.ErrUnauthenticated
	}

	return nil
}

func (a *Authenticator) Logout(r *http.Request, w http.ResponseWriter, id string) error {
	session, _ := a.store.Get(r, authCookieKey)
	session.Values["user"] = nil

	if err := session.Save(r, w); err != nil {
		return err
	}

	return nil
}

func (a *Authenticator) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := a.store.Get(r, authCookieKey)
			value, ok := session.Values["user"]
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			user, ok := value.(security.User)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}

			ctx := security.WithUser(r.Context(), user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}