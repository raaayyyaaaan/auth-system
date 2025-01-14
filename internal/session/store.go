package session

import (
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("your-secret-key")) // In production, this should come from env vars
