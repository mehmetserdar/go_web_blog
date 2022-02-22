package main

import (
	"net/http"

	admin_models "github.com/mehmetserdar/go_web_blog/admin/models"
	"github.com/mehmetserdar/go_web_blog/config"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
