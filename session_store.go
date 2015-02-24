package sunday

import (
   "net/http"
//    "github.com/gorilla/sessions"
)

//Copied verbatim from http://www.gorillatoolkit.org/pkg/sessions#Store
type SessionStore interface {
    Get(r *http.Request, name string) (Session, error)
    New(r *http.Request, name string) (Session, error)
    Save(r *http.Request, w http.ResponseWriter, s Session) error
}

type sessionStore struct {

}