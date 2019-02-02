package main

import (
	"context"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Dish struct {
	Info string
	Name string
}

type ViewDishes struct {
	Dishes []Dish
	Title  string
}

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		data := ViewDishes{
			Title: "Еда",
			Dishes: []Dish{
				Dish{Info: "Tom", Name: "Суп"},
				Dish{Info: "Kate", Name: "Мясо"},
				Dish{Info: "Alice", Name: "Рыба"},
			},
		}

		tmpl, _ := template.ParseFiles("static/index.html")

		tmpl.Execute(w, data)
	})

	// RESTy routes for "articles" resource
	r.Route("/articles", func(r chi.Router) {
		// Subrouters:
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Get("/", AddToSalesOrder)
		})
	})

	http.ListenAndServe(":3333", r)
}

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, "articleID")
		ctx := context.WithValue(r.Context(), "article", articleID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AddToSalesOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(string)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(article))
}

func returnDushTemplate() string {
	return "<div><button name = 'ВЗЯТЬ'><img src='pictures/eda.png' alt='альтернативный текст'></button></div>"
}
