// Package v1 defines version 1 of the role API
package v1

import (
	"fmt"
	"net/http"

	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/role"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// RegisterRoute function creates the role route
func RegisterRoute(e *echo.Group) {
	e.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "api/v1")
	})

	e.GET(role.Path, welcome)
	// Override
	// e.GET(role.Path, func(ctx echo.Context) error {
	// 	return ctx.String(http.StatusOK, "Override path")
	// })
}

// Welcome
func welcome(ctx echo.Context) error {
	if db, ok := ctx.Get("DB").(*gorm.DB); ok {
		// TODO: FAKE DB
		// role := model.Role{
		// 	Title: "test role",
		// 	Slug: "role slug",
		// 	CreatedAT: time.Now(),
		// }
		// result := db.Create(&role)
		// fmt.Println(result)
		service := InitializeRoleService(db)
		roles := []*model.Role{}
		err := service.FindAll(&roles)

		if err != nil {
			fmt.Println(err)
			ctx.String(http.StatusNotFound, "Has error - DB Query")
		}

		return ctx.JSON(http.StatusOK, roles)
	}

	return ctx.String(http.StatusInternalServerError, "Has error - DB context")
}