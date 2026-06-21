package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

const defaultNameSpaceName = "example.com"

var defaultNameSpace = uuid.NameSpaceDNS

type UUIDHandlerFunc func(req *http.Request) (uuid.UUID, error)

func (generator UUIDHandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id, err := generator(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := id.MarshalText()
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, bytes.NewReader(b))
}

func nameBasedUUID(req *http.Request, newFunc func(ns uuid.UUID, data []byte) uuid.UUID) (uuid.UUID, error) {
	ns := defaultNameSpace
	if v := req.URL.Query().Get("namespace"); v != "" {
		parsed, err := uuid.Parse(v)
		if err != nil {
			return uuid.UUID{}, err
		}
		ns = parsed
	}
	name := req.URL.Query().Get("name")
	if name == "" {
		name = defaultNameSpaceName
	}
	return newFunc(ns, []byte(name)), nil
}

func dceUUID(req *http.Request) (uuid.UUID, error) {
	domain := uuid.Person
	if v := req.URL.Query().Get("domain"); v != "" {
		switch v {
		case "person":
			domain = uuid.Person
		case "group":
			domain = uuid.Group
		case "org":
			domain = uuid.Org
		default:
			return uuid.UUID{}, fmt.Errorf("invalid domain: %s (must be person, group, or org)", v)
		}
	}

	id := uint32(os.Getuid())
	if v := req.URL.Query().Get("id"); v != "" {
		parsed, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("invalid id: %s", v)
		}
		id = uint32(parsed)
	}

	return uuid.NewDCESecurity(domain, id)
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", UUIDHandlerFunc(func(req *http.Request) (uuid.UUID, error) {
		return uuid.NewV7()
	}))
	mux.Handle("/v1", UUIDHandlerFunc(func(req *http.Request) (uuid.UUID, error) {
		return uuid.NewUUID()
	}))
	mux.Handle("/v2", UUIDHandlerFunc(dceUUID))
	mux.Handle("/v3", UUIDHandlerFunc(func(req *http.Request) (uuid.UUID, error) {
		return nameBasedUUID(req, uuid.NewMD5)
	}))
	mux.Handle("/v4", UUIDHandlerFunc(func(req *http.Request) (uuid.UUID, error) {
		return uuid.NewRandom()
	}))
	mux.Handle("/v5", UUIDHandlerFunc(func(req *http.Request) (uuid.UUID, error) {
		return nameBasedUUID(req, uuid.NewSHA1)
	}))
	mux.Handle("/v6", UUIDHandlerFunc(func(req *http.Request) (uuid.UUID, error) {
		return uuid.NewV6()
	}))
	mux.Handle("/v7", UUIDHandlerFunc(func(req *http.Request) (uuid.UUID, error) {
		return uuid.NewV7()
	}))
	return mux
}

func main() {
	fmt.Println("listening on :7100")
	http.ListenAndServe(":7100", NewHandler())
}
