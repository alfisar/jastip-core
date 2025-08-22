package router

import (
	countriesControll "jastip-core/application/countries/controller/http"
	orderControll "jastip-core/application/order/controller/http"
	productsControll "jastip-core/application/products/controller/http"
	productTravelControll "jastip-core/application/products_travel/controller/http"
	simpleControll "jastip-core/application/simple/controller/http"
	travelControll "jastip-core/application/travel_schedule/controller/http"

	countriesRepo "jastip-core/application/countries/repository"
	countriesServ "jastip-core/application/countries/service"

	travelRepo "jastip-core/application/travel_schedule/repository"
	travelServ "jastip-core/application/travel_schedule/service"

	productsRepo "jastip-core/application/products/repository"
	productsServ "jastip-core/application/products/service"

	productsTravelRepo "jastip-core/application/products_travel/repository"
	productsTravelServ "jastip-core/application/products_travel/service"
	"os"

	addrOrderRepo "jastip-core/application/address_order/repository"
	orderRepo "jastip-core/application/order/repository"
	orderServ "jastip-core/application/order/service"
	orderDetailRepo "jastip-core/application/order_detail/repository"

	"github.com/alfisar/jastip-import/domain"
	pb "github.com/alfisar/jastip-import/proto/auth"
	clientAuth "github.com/alfisar/jastip-import/proto/auth/client"

	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/jwthandler"
	"github.com/alfisar/jastip-import/helpers/middlewere"
)

func SimpleInit() *simpleRouter {
	controll := simpleControll.NewSimpleController()

	return NewSimpleRouter(controll)
}

func TravelSchInit() *travelScheduleRouter {
	repoContries := countriesRepo.NewCountriesRepository()
	repo := travelRepo.NewTravelSchRepository()
	serv := travelServ.NewTravelSchService(repo, repoContries)

	controll := travelControll.NewTravelController(serv)
	return NewTravelSchRouter(controll)
}

func CountriesInit() *countriesRouter {
	repo := countriesRepo.NewCountriesRepository()
	serv := countriesServ.NewCountriesService(repo)

	controll := countriesControll.NewCountriesController(serv)
	return NewCountriesRouter(controll)
}

func ProductsInit() *productsRouter {
	repo := productsRepo.NewProductsRepository()
	serv := productsServ.NewProductsService(repo)

	controll := productsControll.NewProductController(serv)
	return NewProductRouter(controll)
}

func ProductsTravelInit() *productsTravelRouter {
	repoProduct := productsRepo.NewProductsRepository()
	repoTravel := travelRepo.NewTravelSchRepository()
	productstravelRepo := productsTravelRepo.NewProductTravelrepository()
	servProductTravel := productsTravelServ.NewProductsTravelService(productstravelRepo, repoProduct, repoTravel)

	controll := productTravelControll.NewProductsTravelController(servProductTravel)
	return NewProductTravelRouter(controll)
}

func OrderInit() *orderRouter {
	repoProduct := productsRepo.NewProductsRepository()
	repoTravel := travelRepo.NewTravelSchRepository()
	repoAddrOrder := addrOrderRepo.NewAddressOrderRepository()
	repoOrderDetail := orderDetailRepo.NewOrderDetailsRepository()
	repoOrder := orderRepo.NewOrderRepository()
	productstravelRepo := productsTravelRepo.NewProductTravelrepository()

	poolData := domain.DataPool.Get().(*domain.Config)
	connAuthClient := poolData.GRPC[consts.GrpcAuth]
	profileClient := pb.NewProfileClient(connAuthClient)
	clientAuths := clientAuth.NewAuthClient(profileClient)

	servOrder := orderServ.NewOrderService(repoOrder, repoOrderDetail, repoAddrOrder, repoProduct, productstravelRepo, repoTravel, clientAuths)

	controll := orderControll.NewOrderService(servOrder)
	return NewOrderRouter(controll)
}

func setMiddleware() *middlewere.AuthenticateMiddleware {
	jwtData := jwthandler.GetJwt()
	jwtData.Secret = os.Getenv("JWT_SECRET")
	middleWR := middlewere.NewAuthenticateMiddleware(jwtData)
	return middleWR
}
