package session

import (
	"net/http"

	"github.com/xialvjun/koa.go/koa"
	uuid "other/github.com/satori/go.uuid"
)

var key string

// todo: wrap mem in a mutex
var mem = make(map[string]interface{})

func Middleware(setkey string) koa.Middleware {
	key = setkey
	return func(w http.ResponseWriter, r *http.Request, next koa.Next) {
		sid, error := r.Cookie(key)
		if error != nil || sid.Value == "" {
			uid := uuid.NewV4().String()
			w.Header().Set("Set-Cookie", "sid="+uid)
			r.AddCookie(&http.Cookie{Name: key, Value: uid})
		}
	}
}

func Set(r *http.Request, value interface{}) error {
	sid, error := r.Cookie(key)
	if error != nil {
		return error
	}
	mem[sid.Value] = value
	return nil
}

func Get(r *http.Request) (interface{}, error) {
	sid, err := r.Cookie(key)
	if err != nil {
		return nil, err
	}
	value, ok := mem[sid.Value]
	if ok {
		return value, nil
	}
	return nil, nil
}
