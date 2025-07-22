package router

import (
	simpleControll "jastip-core/application/simple/controller/tcp"
)

func SimpleInit() *simpleRouter {
	controll := simpleControll.NewSimpleController()

	return NewSimpleRouter(*controll)
}
