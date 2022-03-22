package controller

import (
	"net/http"

	"github.com/erikrios/go-clean-arhictecture/middleware"
	"github.com/erikrios/go-clean-arhictecture/model"
	"github.com/erikrios/go-clean-arhictecture/service"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserController(userService service.UserService, authService service.AuthService) *UserController {
	return &UserController{userService: userService, authService: authService}
}

func (u *UserController) Route(e *echo.Echo) {
	e.GET("/users", u.GetAll, middleware.JWTMiddleware())
	e.POST("/users", u.Create)
	e.POST("/users/login", u.Login)
}

func (u *UserController) GetAll(c echo.Context) error {
	responses, err := u.userService.List()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": responses,
	})
}

func (u *UserController) Create(c echo.Context) error {
	var request model.CreateUserRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	response, err := u.userService.Create(request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": response,
	})
}

func (u *UserController) Login(c echo.Context) error {
	var request model.LoginUserRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	response, err := u.userService.Login(request)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := u.authService.GenerateToken(response.Id, response.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Something went wrong",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
