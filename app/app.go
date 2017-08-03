package main

import (
	//"encoding/json"
	//"errors"
	"fmt"
	//"io"
	"log"
	"net/http"
	"os"
	//"path"
	//"strconv"

	//"cloud.google.com/go/pubsub"
	//"cloud.google.com/go/storage"

	//"golang.org/x/net/context"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/satori/go.uuid"

	"github.com/GoogleCloudPlatform/golang-samples/getting-started/productshelf"
	"google.golang.org/appengine"
)

var (
	// See template.go
	editTmpl = parseTemplate("edit.html")
)

func main() {
	registerHandlers()
	appengine.Main()
}

func registerHandlers() {

	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/products", http.StatusFound))

	r.Methods("GET").Path("/products").
		Handler(appHandler(addFormHandler))

	r.Methods("POST").Path("/products").
		Handler(appHandler(createHandler))

		// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	// [END request_logging]

}

// addFormHandler displays a form that captures details of a new book to add to
// the database.
func addFormHandler(w http.ResponseWriter, r *http.Request) *appError {
	return editTmpl.Execute(w, r, nil)
}

// bookFromForm populates the fields of a Book from form values
// (see templates/edit.html).
func productFromForm(r *http.Request) (*productshelf.Product, error) {

	product := &productshelf.Product{
		Title:         r.FormValue("title"),
		Author:        r.FormValue("author"),
		PublishedDate: r.FormValue("publishedDate"),
		Description:   r.FormValue("description"),
	}
	return product, nil
}

// createHandler adds a book to the database.
func createHandler(w http.ResponseWriter, r *http.Request) *appError {
	product, err := productFromForm(r)
	if err != nil {
		return appErrorf(err, "could not parse product from form: %v", err)
	}
	id, err := productshelf.DB.AddProduct(product)
	if err != nil {
		return appErrorf(err, "could not save product: %v", err)
	}
	//go publishUpdate(id)
	//fmt.Printf(id)
	http.Redirect(w, r, fmt.Sprintf("/products/%d", id), http.StatusFound)
	return nil
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}
