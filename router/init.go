package router

import "jastip-core/application/simple/controller"

func SimpleInit() *simpleRouter {
	controll := controller.NewSimpleController()

	return NewSimpleRouter(controll)
}
