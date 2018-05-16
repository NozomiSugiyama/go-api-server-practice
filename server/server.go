package server

import (
	"net/http"
	"strconv"
	"github.com/urfave/negroni"
)

func StartServer(host string, port int) error {
	router := NewRouter()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	addr := host + ":" + strconv.Itoa(port)

	if err := http.ListenAndServe(addr, n); err != nil {
		return err
	}

	return nil
}
