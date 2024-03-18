// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"filmoteka_server/models"
	"filmoteka_server/restapi/operations/actor"
	"filmoteka_server/restapi/operations/film"
	"filmoteka_server/restapi/operations/user"
)

// NewFilmotekaAPI creates a new Filmoteka instance
func NewFilmotekaAPI(spec *loads.Document) *FilmotekaAPI {
	return &FilmotekaAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ActorAddActorHandler: actor.AddActorHandlerFunc(func(params actor.AddActorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation actor.AddActor has not yet been implemented")
		}),
		FilmAddFilmHandler: film.AddFilmHandlerFunc(func(params film.AddFilmParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation film.AddFilm has not yet been implemented")
		}),
		ActorDeleteActorHandler: actor.DeleteActorHandlerFunc(func(params actor.DeleteActorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation actor.DeleteActor has not yet been implemented")
		}),
		FilmDeleteFilmHandler: film.DeleteFilmHandlerFunc(func(params film.DeleteFilmParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation film.DeleteFilm has not yet been implemented")
		}),
		FilmFindFilmHandler: film.FindFilmHandlerFunc(func(params film.FindFilmParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation film.FindFilm has not yet been implemented")
		}),
		ActorGetAllActorsHandler: actor.GetAllActorsHandlerFunc(func(params actor.GetAllActorsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation actor.GetAllActors has not yet been implemented")
		}),
		FilmGetAllFilmsSortedHandler: film.GetAllFilmsSortedHandlerFunc(func(params film.GetAllFilmsSortedParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation film.GetAllFilmsSorted has not yet been implemented")
		}),
		ActorUpdateActorHandler: actor.UpdateActorHandlerFunc(func(params actor.UpdateActorParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation actor.UpdateActor has not yet been implemented")
		}),
		FilmUpdateFilmHandler: film.UpdateFilmHandlerFunc(func(params film.UpdateFilmParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation film.UpdateFilm has not yet been implemented")
		}),
		UserUserLoginHandler: user.UserLoginHandlerFunc(func(params user.UserLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UserLogin has not yet been implemented")
		}),
		UserUserLogoutHandler: user.UserLogoutHandlerFunc(func(params user.UserLogoutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user.UserLogout has not yet been implemented")
		}),

		// Applies when the "admin-key" header is set
		IsAdminAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (isAdmin) admin-key from header param [admin-key] has not yet been implemented")
		},
		// Applies when the "user-key" header is set
		IsUserAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (isUser) user-key from header param [user-key] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*FilmotekaAPI --- */
type FilmotekaAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// IsAdminAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key admin-key provided in the header
	IsAdminAuth func(string) (*models.Principal, error)

	// IsUserAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key user-key provided in the header
	IsUserAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ActorAddActorHandler sets the operation handler for the add actor operation
	ActorAddActorHandler actor.AddActorHandler
	// FilmAddFilmHandler sets the operation handler for the add film operation
	FilmAddFilmHandler film.AddFilmHandler
	// ActorDeleteActorHandler sets the operation handler for the delete actor operation
	ActorDeleteActorHandler actor.DeleteActorHandler
	// FilmDeleteFilmHandler sets the operation handler for the delete film operation
	FilmDeleteFilmHandler film.DeleteFilmHandler
	// FilmFindFilmHandler sets the operation handler for the find film operation
	FilmFindFilmHandler film.FindFilmHandler
	// ActorGetAllActorsHandler sets the operation handler for the get all actors operation
	ActorGetAllActorsHandler actor.GetAllActorsHandler
	// FilmGetAllFilmsSortedHandler sets the operation handler for the get all films sorted operation
	FilmGetAllFilmsSortedHandler film.GetAllFilmsSortedHandler
	// ActorUpdateActorHandler sets the operation handler for the update actor operation
	ActorUpdateActorHandler actor.UpdateActorHandler
	// FilmUpdateFilmHandler sets the operation handler for the update film operation
	FilmUpdateFilmHandler film.UpdateFilmHandler
	// UserUserLoginHandler sets the operation handler for the user login operation
	UserUserLoginHandler user.UserLoginHandler
	// UserUserLogoutHandler sets the operation handler for the user logout operation
	UserUserLogoutHandler user.UserLogoutHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *FilmotekaAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *FilmotekaAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *FilmotekaAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *FilmotekaAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *FilmotekaAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *FilmotekaAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *FilmotekaAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *FilmotekaAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *FilmotekaAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the FilmotekaAPI
func (o *FilmotekaAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.IsAdminAuth == nil {
		unregistered = append(unregistered, "AdminKeyAuth")
	}
	if o.IsUserAuth == nil {
		unregistered = append(unregistered, "UserKeyAuth")
	}

	if o.ActorAddActorHandler == nil {
		unregistered = append(unregistered, "actor.AddActorHandler")
	}
	if o.FilmAddFilmHandler == nil {
		unregistered = append(unregistered, "film.AddFilmHandler")
	}
	if o.ActorDeleteActorHandler == nil {
		unregistered = append(unregistered, "actor.DeleteActorHandler")
	}
	if o.FilmDeleteFilmHandler == nil {
		unregistered = append(unregistered, "film.DeleteFilmHandler")
	}
	if o.FilmFindFilmHandler == nil {
		unregistered = append(unregistered, "film.FindFilmHandler")
	}
	if o.ActorGetAllActorsHandler == nil {
		unregistered = append(unregistered, "actor.GetAllActorsHandler")
	}
	if o.FilmGetAllFilmsSortedHandler == nil {
		unregistered = append(unregistered, "film.GetAllFilmsSortedHandler")
	}
	if o.ActorUpdateActorHandler == nil {
		unregistered = append(unregistered, "actor.UpdateActorHandler")
	}
	if o.FilmUpdateFilmHandler == nil {
		unregistered = append(unregistered, "film.UpdateFilmHandler")
	}
	if o.UserUserLoginHandler == nil {
		unregistered = append(unregistered, "user.UserLoginHandler")
	}
	if o.UserUserLogoutHandler == nil {
		unregistered = append(unregistered, "user.UserLogoutHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *FilmotekaAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *FilmotekaAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "isAdmin":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.IsAdminAuth(token)
			})

		case "isUser":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.IsUserAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *FilmotekaAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *FilmotekaAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *FilmotekaAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *FilmotekaAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the filmoteka API
func (o *FilmotekaAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *FilmotekaAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/actors"] = actor.NewAddActor(o.context, o.ActorAddActorHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/films"] = film.NewAddFilm(o.context, o.FilmAddFilmHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/actors"] = actor.NewDeleteActor(o.context, o.ActorDeleteActorHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/films"] = film.NewDeleteFilm(o.context, o.FilmDeleteFilmHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/films/find"] = film.NewFindFilm(o.context, o.FilmFindFilmHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/actors"] = actor.NewGetAllActors(o.context, o.ActorGetAllActorsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/films"] = film.NewGetAllFilmsSorted(o.context, o.FilmGetAllFilmsSortedHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/actors"] = actor.NewUpdateActor(o.context, o.ActorUpdateActorHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/films"] = film.NewUpdateFilm(o.context, o.FilmUpdateFilmHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/login"] = user.NewUserLogin(o.context, o.UserUserLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/logout"] = user.NewUserLogout(o.context, o.UserUserLogoutHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *FilmotekaAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *FilmotekaAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *FilmotekaAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *FilmotekaAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *FilmotekaAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
