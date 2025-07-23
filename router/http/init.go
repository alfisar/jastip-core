package router

import (
	countriesControll "jastip-core/application/countries/controller/http"
	simpleControll "jastip-core/application/simple/controller/http"
	travelControll "jastip-core/application/travel_schedule/controller/http"

	countriesRepo "jastip-core/application/countries/repository"
	countriesServ "jastip-core/application/countries/service"
	travelRepo "jastip-core/application/travel_schedule/repository"
	travelServ "jastip-core/application/travel_schedule/service"
	"os"

	"github.com/alfisar/jastip-import/helpers/jwthandler"
	"github.com/alfisar/jastip-import/helpers/middlewere"
)

func SimpleInit() *simpleRouter {
	controll := simpleControll.NewSimpleController()

	return NewSimpleRouter(controll)
}

func TravelSchInit() *travelScheduleRouter {
	repo := travelRepo.NewTravelSchRepository()
	serv := travelServ.NewTravelSchService(repo)

	controll := travelControll.NewTravelController(serv)
	return NewTravelSchRouter(controll)
}

func CountriesInit() *countriesRouter {
	repo := countriesRepo.NewCountriesRepository()
	serv := countriesServ.NewCountriesService(repo)

	controll := countriesControll.NewCountriesController(serv)
	return NewCountriesRouter(controll)
}

func setMiddleware() *middlewere.AuthenticateMiddleware {
	jwtData := jwthandler.GetJwt()
	jwtData.Secret = os.Getenv("JWT_SECRET")
	middleWR := middlewere.NewAuthenticateMiddleware(jwtData)
	return middleWR
}
