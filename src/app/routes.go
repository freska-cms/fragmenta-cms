package app

import (
	"github.com/freska-cms/mux"
	"github.com/freska-cms/server/log"

	// Resource Actions
	"github.com/freska-cms/freska-cms/src/images/actions"
	"github.com/freska-cms/freska-cms/src/lib/session"
	"github.com/freska-cms/freska-cms/src/pages/actions"
	"github.com/freska-cms/freska-cms/src/posts/actions"
	"github.com/freska-cms/freska-cms/src/tags/actions"
	"github.com/freska-cms/freska-cms/src/users/actions"
)

// SetupRoutes creates a new router and adds the routes for this app to it.
func SetupRoutes() *mux.Mux {

	router := mux.New()
	mux.SetDefault(router)

	// Set the default file handler
	router.FileHandler = fileHandler
	router.ErrorHandler = errHandler

	// Add middleware
	router.AddMiddleware(log.Middleware)
	router.AddMiddleware(session.Middleware)

	// Add the home page route
	router.Get("/", pageactions.HandleShowHome)

	// Add a route to handle static files
	router.Add("/favicon.ico", fileHandler)
	router.Add("/files/{path:.*}", fileHandler)
	router.Add("/assets/{path:.*}", fileHandler)

	// Resource Routes

	router.Get("/pages", pageactions.HandleIndex)
	router.Get("/pages/create", pageactions.HandleCreateShow)
	router.Post("/pages/create", pageactions.HandleCreate)
	router.Get("/pages/{id:[0-9]+}/update", pageactions.HandleUpdateShow)
	router.Post("/pages/{id:[0-9]+}/update", pageactions.HandleUpdate)
	router.Post("/pages/{id:[0-9]+}/destroy", pageactions.HandleDestroy)
	router.Get("/pages/{id:[0-9]+}", pageactions.HandleShow)
	router.Get("/freska/setup", pageactions.HandleSetupShow)
	router.Post("/freska/setup", pageactions.HandleSetup)

	router.Get("/images", imageactions.HandleIndex)
	router.Get("/images/create", imageactions.HandleCreateShow)
	router.Post("/images/create", imageactions.HandleCreate)
	router.Get("/images/{id:[0-9]+}/update", imageactions.HandleUpdateShow)
	router.Post("/images/{id:[0-9]+}/update", imageactions.HandleUpdate)
	router.Post("/images/{id:[0-9]+}/destroy", imageactions.HandleDestroy)
	router.Get("/images/{id:[0-9]+}", imageactions.HandleShow)

	router.Get("/posts", postactions.HandleIndex)
	router.Get("/posts/create", postactions.HandleCreateShow)
	router.Post("/posts/create", postactions.HandleCreate)
	router.Get("/posts/{id:[0-9]+}/update", postactions.HandleUpdateShow)
	router.Post("/posts/{id:[0-9]+}/update", postactions.HandleUpdate)
	router.Post("/posts/{id:[0-9]+}/destroy", postactions.HandleDestroy)
	router.Get("/posts/{id:[0-9]+}", postactions.HandleShow)
	router.Get("/blog", postactions.HandleShowBlog)
	router.Get("/blog/{id:[0-9]+}", postactions.HandleShow)

	router.Get("/tags", tagactions.HandleIndex)
	router.Get("/tags/create", tagactions.HandleCreateShow)
	router.Post("/tags/create", tagactions.HandleCreate)
	router.Get("/tags/{id:[0-9]+}/update", tagactions.HandleUpdateShow)
	router.Post("/tags/{id:[0-9]+}/update", tagactions.HandleUpdate)
	router.Post("/tags/{id:[0-9]+}/destroy", tagactions.HandleDestroy)
	router.Get("/tags/{id:[0-9]+}", tagactions.HandleShow)

	router.Get("/users", useractions.HandleIndex)
	router.Get("/users/create", useractions.HandleCreateShow)
	router.Post("/users/create", useractions.HandleCreate)
	router.Get("/users/login", useractions.HandleLoginShow)
	router.Post("/users/login", useractions.HandleLogin)
	router.Post("/users/logout", useractions.HandleLogout)
	router.Get("/users/{id:[0-9]+}/update", useractions.HandleUpdateShow)
	router.Post("/users/{id:[0-9]+}/update", useractions.HandleUpdate)
	router.Post("/users/{id:[0-9]+}/destroy", useractions.HandleDestroy)
	router.Get("/users/{id:[0-9]+}", useractions.HandleShow)

	// Add catch-all for custom page routes - this must be evaluated last.
	router.Get("/{path:.+}", pageactions.HandleShowPath)

	return router
}
