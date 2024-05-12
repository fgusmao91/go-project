package runtime

import (
	"database/sql"
	"v1/controller"
	"v1/repository"
	"v1/service"

	_ "github.com/go-sql-driver/mysql"
)

type Controllers struct {
	LoginController *controller.AuthController
}

func InjectDependencies() Controllers {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/logindb")
	if err != nil {
		panic(err)
	}

	//repositories
	loginRepository := repository.NewAuthRepository(db)

	//services
	loginService := service.NewAuthService(*loginRepository)

	//controllers
	loginController := controller.NewAuthController(*loginService)

	return Controllers{
		LoginController: loginController,
	}
}
