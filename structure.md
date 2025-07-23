.
├── Dockerfile
├── Format 1 (tree-like):
├── app
├── application
│   ├── countries
│   │   ├── controller
│   │   │   └── http
│   │   │       ├── countries_controller.go
│   │   │       └── countries_controller_contract.go
│   │   ├── repository
│   │   │   ├── countries_repository.go
│   │   │   └── countries_repository_contract.go
│   │   └── service
│   │       ├── countries_service.go
│   │       ├── countries_service_contract.go
│   │       └── countries_service_utility.go
│   ├── simple
│   │   └── controller
│   │       ├── http
│   │       │   ├── simple_controller.go
│   │       │   └── simple_controller_contract.go
│   │       └── tcp
│   │           └── simple_controller.go
│   └── travel_schedule
│       ├── controller
│       │   └── http
│       │       ├── travel_schedule_controller.go
│       │       └── travel_schedule_controller_contract.go
│       ├── repository
│       │   ├── coverage
│       │   │   └── coverageTravelSch.html
│       │   ├── travel_schedule_repository.go
│       │   ├── travel_schedule_repository_contract.go
│       │   └── travel_schedule_repository_test.go
│       └── service
│           ├── travel_schedule_service.go
│           ├── travel_schedule_service_contract.go
│           └── travel_schedule_service_utility.go
├── components
├── go.mod
├── go.sum
├── main.go
├── router
│   ├── http
│   │   ├── countries_router.go
│   │   ├── init.go
│   │   ├── router.go
│   │   ├── simple_router.go
│   │   └── travel_schedule_router.go
│   └── tcp
│       ├── init.go
│       ├── router.go
│       └── simple_router.go
├── src
└── structure.md

24 directories, 32 files
