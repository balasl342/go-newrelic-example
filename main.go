package main

import (
	"fmt"
	"net/http"
	"relicapp/database"
	"relicapp/handlers"
	"relicapp/middleware"
)

func main() {
	// Initialize New Relic application
	app, err := middleware.NewRelicApp()
	if err != nil {
		fmt.Println("error creating New Relic application:", err)
		return
	}

	// Initialize database
	database.InitDatabase()

	// txn := app.StartTransaction("transaction_name")
	// defer txn.End()
	// Wrap handlers with New Relic
	http.HandleFunc(middleware.WrapHandleFunc(app, "/", handlers.Hello))
	http.HandleFunc(middleware.WrapHandleFunc(app, "/create_user", handlers.CreateUser))
	http.HandleFunc(middleware.WrapHandleFunc(app, "/get_user/{id}", handlers.GetUser))
	http.HandleFunc(middleware.WrapHandleFunc(app, "/get_all_users", handlers.GetAllUsers))

	// Start the server
	http.ListenAndServe(":8000", nil)
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/newrelic/go-agent/v3/newrelic"
// )

// func main() {
// 	// Initialize New Relic application
// 	app, err := newrelic.NewApplication(
// 		newrelic.ConfigAppName("ApplicationName"),
// 		newrelic.ConfigLicense("69314b6dea4e28d5cfb141000d0c7a20FFFFNRAL"),
// 		newrelic.ConfigDistributedTracerEnabled(true),
// 	)

// 	if err != nil {
// 		fmt.Println("error creating New Relic application:", err)
// 		return
// 	}

// 	// Wrap handlers with New Relic
// 	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", hello))

// 	// Start the server
// 	http.ListenAndServe(":8000", nil)
// }

// // Handler function
// func hello(w http.ResponseWriter, r *http.Request) {
// 	// Create a New Relic transaction
// 	txn := newrelic.FromContext(r.Context())
// 	defer txn.StartSegment("hello-segment").End()

// 	// Write response
// 	w.Write([]byte("Hello, World!"))
// }
