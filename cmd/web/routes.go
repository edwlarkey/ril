package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	// Home
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	// Articles
	mux.Get("/article/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createArticleForm))
	mux.Post("/article/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createArticle))
	mux.Get("/article/:id", dynamicMiddleware.ThenFunc(app.showArticle))

	// User
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	// Ping
	mux.Get("/ping", http.HandlerFunc(ping))

	// Static assets
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)

}
