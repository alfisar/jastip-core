package router

import (
	simpleControll "jastip-core/application/simple/controller/http"
	travelControll "jastip-core/application/travel_schedule/controller"
	"jastip-core/application/travel_schedule/repository"
	"jastip-core/application/travel_schedule/service"
	"os"

	"github.com/alfisar/jastip-import/helpers/jwthandler"
	"github.com/alfisar/jastip-import/helpers/middlewere"
)

func SimpleInit() *simpleRouter {
	controll := simpleControll.NewSimpleController()

	return NewSimpleRouter(controll)
}

func TravelSchInit() *travelScheduleRouter {
	repo := repository.NewTravelSchRepository()
	serv := service.NewTravelSchService(repo)

	controll := travelControll.NewTravelController(serv)
	return NewTravelSchRouter(controll)
}

func setMiddleware() *middlewere.AuthenticateMiddleware {
	jwtData := jwthandler.GetJwt()
	jwtData.Secret = os.Getenv("JWT_SECRET")
	middleWR := middlewere.NewAuthenticateMiddleware(jwtData)
	return middleWR
}
