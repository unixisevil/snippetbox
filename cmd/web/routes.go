package main

import (
	"net/http"

	fr "github.com/DATA-DOG/fastroute"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	stdMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dyn := alice.New(app.session.Enable, noSurf, app.authenticate)
	auth := dyn.Append(app.requireAuthentication)
	fileServer := http.FileServer(disableDirList{http.Dir("./ui/static/")})
	routes := map[string]fr.Router{
		"GET": fr.Chain(
			fr.New("/", dyn.ThenFunc(app.home)),
			fr.New("/snippet/create", auth.ThenFunc(app.createSnippetForm)),
			fr.New("/snippet/:id", dyn.ThenFunc(app.showSnippet)),
			fr.New("/static/*all", http.StripPrefix("/static", fileServer)),
			fr.New("/user/signup", dyn.ThenFunc(app.signupUserForm)),
			fr.New("/user/login", dyn.ThenFunc(app.loginUserForm)),
			fr.New("/ping", ping),
		),
		"POST": fr.Chain(
			fr.New("/snippet/create", auth.ThenFunc(app.createSnippet)),
			fr.New("/user/signup", dyn.ThenFunc(app.signupUser)),
			fr.New("/user/login", dyn.ThenFunc(app.loginUser)),
			fr.New("/user/logout", auth.ThenFunc(app.logoutUser)),
		),
	}
	return stdMiddleware.Then(fr.RouterFunc(func(req *http.Request) http.Handler {
		return routes[req.Method]
	}))
}
