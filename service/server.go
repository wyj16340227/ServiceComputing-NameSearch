package service

import (
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/{command}:{name}:{email}:{phone}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        command := vars["command"]
        name := vars["name"]
        email := vars["email"]
        phone := vars["phone"]
        result := Manager(command, name, email, phone)
        formatter.JSON(w, http.StatusOK, struct{ RESULT string }{result})
    }
}