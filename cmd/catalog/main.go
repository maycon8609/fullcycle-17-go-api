package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maycon8609/fullcycle-17-go-api/internal/database"
	"github.com/maycon8609/fullcycle-17-go-api/internal/service"
	"github.com/maycon8609/fullcycle-17-go-api/internal/webserver"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandle := webserver.NewWebCategoryHandler(categoryService)
	webProductHandle := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	// category
	c.Get("/category", webCategoryHandle.GetCategories)
	c.Get("/category/{id}", webCategoryHandle.GetCategory)
	c.Post("/category", webCategoryHandle.CreateCategory)

	// product
	c.Get("/product", webProductHandle.GetProducts)
	c.Get("/product/{categoryID}", webProductHandle.GetProductsByCategoryID)
	c.Get("/product/{id}", webProductHandle.GetProduct)
	c.Post("/product", webProductHandle.CreateProduct)

	fmt.Println("Server is ruining on port 8080")
	http.ListenAndServe(":8080", c)
}
