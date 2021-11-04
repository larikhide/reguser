// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create user
	// (POST /create)
	PostCreate(w http.ResponseWriter, r *http.Request)
	// Delete user
	// (DELETE /delete/{id})
	DeleteDeleteId(w http.ResponseWriter, r *http.Request, id string)
	// Get user
	// (GET /read/{id})
	GetReadId(w http.ResponseWriter, r *http.Request, id string)
	// Search user
	// (GET /search/{q})
	FindUsers(w http.ResponseWriter, r *http.Request, q string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// PostCreate operation middleware
func (siw *ServerInterfaceWrapper) PostCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostCreate(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteDeleteId operation middleware
func (siw *ServerInterfaceWrapper) DeleteDeleteId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteDeleteId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetReadId operation middleware
func (siw *ServerInterfaceWrapper) GetReadId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetReadId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// FindUsers operation middleware
func (siw *ServerInterfaceWrapper) FindUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "q" -------------
	var q string

	err = runtime.BindStyledParameter("simple", false, "q", chi.URLParam(r, "q"), &q)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter q: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FindUsers(w, r, q)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/create", wrapper.PostCreate)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/delete/{id}", wrapper.DeleteDeleteId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/read/{id}", wrapper.GetReadId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/search/{q}", wrapper.FindUsers)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+ySy2vcMBDG/xUzZxNvXxff+gxLDw3Z9lRy0FqTXQXrsTPjwGL0v5eRk5CuXUqhbC+9",
	"GEnfzKfR9/MIXfQpBgzC0I65BhduI7QjWOSOXBIXA7Rw/XHztXp7tYYaxEmP0MLG+dRjxUj3rsMH8R6J",
	"p44XF6uLFeQaYsJgkoMWXpWjGpKRvd4GTUdoBHWZIsv82vdFrwZGguJERpW1hRauIsukQw2EhwFZ3kV7",
	"VJcuBsFQDE1KvetKW3PH6joCd3v0RldyTPqWuL3DTiDnXJ9MoC3VVm2nSxyhhVZowKwHnGJgLI95uVrN",
	"H/Dls0bweknaGls9jK01b5ZqXBCkYPqSMlKFRJFAp+TBe0PHk4xUaSz2KNiMzubJUbdz7w/lfDnbSZu+",
	"a1uQkfEoSAzt99mU9tHF6VbxQg3BeCziLLl6DoCFXNhBzjfLqf49oGcg8jzZQoTQ2CceO1z40S9Rlklc",
	"olyjsf8h/CmEp0QLAUZD3b4ZD79GsCklyxQ+uWC/sQb/GwraXWno1S1FX21x58IyksNZiThBzz8tThjV",
	"jweGyBzn+3/C8DmSMsFUOwU/UA8tNJBv8o8AAAD//1zjgi7GBgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}