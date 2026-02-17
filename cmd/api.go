package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/majesticdrag0n/ecom/internal/adapters/postgresql/sqlc"
	"github.com/majesticdrag0n/ecom/internal/customers"
	"github.com/majesticdrag0n/ecom/internal/orders"
	"github.com/majesticdrag0n/ecom/internal/products"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) //ratelimiting
	r.Use(middleware.RealIP)    //rate limiiting and analytics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})
	productSevice := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productSevice)
	r.Get("/products", productHandler.ListProducts)
	r.Get("/products/{id}", productHandler.GetProduct)

	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)
	r.Post("/orders", ordersHandler.PlaceOrder)

	customerService := customers.NewService(repo.New(app.db))
	customerHandler := customers.NewHandler(customerService)
	r.Post("/customers", customerHandler.CreateCustomer)
	r.Get("/customers", customerHandler.ListCustomers)
	r.Get("/customers/{id}", customerHandler.GetCustomer)
	r.Put("/customers/{id}", customerHandler.UpdateCustomer)
	r.Delete("/customers/{id}", customerHandler.DeleteCustomer)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.address,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has Started at address %s", app.config.address)

	return srv.ListenAndServe()

}

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	address string
	db      dbconfig
}

type dbconfig struct {
	dsn string
}
