package runtime

import (
	"v1/controller"
	"v1/service"
)

type Controllers struct {
	LoginController *controller.LoginController
}

func InjectDependencies() Controllers {
	//services
	loginService := service.NewLoginService()

	//controllers
	loginController := controller.NewLoginController(*loginService)

	return Controllers{
		LoginController: loginController,
	}
}
