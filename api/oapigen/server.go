// Package oapigen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package oapigen

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
	// Creates a new task
	// (POST /v1/tasks)
	CreateTask(w http.ResponseWriter, r *http.Request, params CreateTaskParams)
	// Marks a task for deletion
	// (DELETE /v1/tasks/{name})
	DeleteTaskByName(w http.ResponseWriter, r *http.Request, name string)
	// Gets a task by name
	// (GET /v1/tasks/{name})
	GetTaskByName(w http.ResponseWriter, r *http.Request, name string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateTaskParams

	// ------------- Optional query parameter "run" -------------
	if paramValue := r.URL.Query().Get("run"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "run", r.URL.Query(), &params.Run)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter run: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTask(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteTaskByName operation middleware
func (siw *ServerInterfaceWrapper) DeleteTaskByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", chi.URLParam(r, "name"), &name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTaskByName(w, r, name)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetTaskByName operation middleware
func (siw *ServerInterfaceWrapper) GetTaskByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", chi.URLParam(r, "name"), &name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTaskByName(w, r, name)
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
		r.Post(options.BaseURL+"/v1/tasks", wrapper.CreateTask)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/v1/tasks/{name}", wrapper.DeleteTaskByName)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/tasks/{name}", wrapper.GetTaskByName)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
