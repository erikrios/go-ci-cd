package main

import (
	"fmt"
	"log"
	"os"

	"github.com/erikrios/go-clean-arhictecture/config"
	"github.com/erikrios/go-clean-arhictecture/controller"
	"github.com/erikrios/go-clean-arhictecture/repository"
	"github.com/erikrios/go-clean-arhictecture/service"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("$PORT must be set")
	}

	db, err := config.NewMySQLDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	err = config.MigrateMySQLDatabase(db)
	if err != nil {
		log.Fatalln(err)
	}

	userRepository := repository.NewUserRepositoryImpl(db)

	userService := service.NewUserServiceImpl(userRepository)
	authService := service.NewAuthServiceImpl()

	userController := controller.NewUserController(userService, authService)

	app := echo.New()

	userController.Route(app)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", port)))
}
