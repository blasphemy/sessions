package sessions

import (
	"github.com/dancannon/gorethink"
	"github.com/gorilla/sessions"
	"github.com/miquella/rethinkdb_session_store"
)

// RethinkStore is an interface that represents a Cookie based storage
// for Sessions.
type RethinkStore interface {
	// Store is an embedded interface so that RethinkStore can be used
	// as a session store.
	Store
	// Options sets the default options for each session stored in this
	// CookieStore.
	Options(Options)
}

// NewCookieStore returns a new CookieStore.
//
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewRethinkStore(rethinkdbSession *gorethink.Session, db, table string, keyPairs ...[]byte) RethinkStore {
	store := rethinkdb_session_store.NewRethinkDBStore(rethinkdbSession, db, table, keyPairs...)
	return &rethinkStore{store}
}

type rethinkStore struct {
	*rethinkdb_session_store.RethinkDBStore
}

func (c *rethinkStore) Options(options Options) {
	c.RethinkDBStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
