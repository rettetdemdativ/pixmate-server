/*
 * IMGTURLTE
 * GO PROTOTYPE
 * 2015
 */

package http

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

// Start is the http packages launch method
// It creates all the routes, adds negroni logging and starts the server
func Start() {
	r := mux.NewRouter().StrictSlash(false)

	/*
	 * main page
	 */
	r.HandleFunc("/", mainPageHandler)

	/*
	 * custom profile url
	 */
	r.HandleFunc("/u/{name}", peoplePageHandler)

	/*
	 * personal profile url
	 */
	r.HandleFunc("/me", mePageHandler)

	/*
	 * images
	 */
	//r.HandleFunc("/img/{id}.jpg", imageJPGHandler)
	r.HandleFunc("/img/{id}", imageHandler)

	/*
	 * init negroni middleware
	 */
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MiddleWare),
		//negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)
	n.UseHandler(r)
	color.Green("imgturtle Server running on port 8000")

	http.ListenAndServe(":8000", n)
	//http.ListenAndServeTLS(port, certificate.pem, key.pem, nil) for https
}
