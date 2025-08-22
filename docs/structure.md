.
├── Dockerfile
├── README.md
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
│   ├── products
│   │   ├── controller
│   │   │   └── http
│   │   │       ├── products_controller.go
│   │   │       └── products_controller_contract.go
│   │   ├── repository
│   │   │   ├── products_repository.go
│   │   │   └── products_repository_contract.go
│   │   └── service
│   │       ├── products_service.go
│   │       ├── products_service_contract.go
│   │       └── products_service_utility.go
│   ├── products_travel
│   │   ├── controller
│   │   │   └── http
│   │   │       ├── products_travel_controller.go
│   │   │       └── products_travel_controller_contract.go
│   │   ├── repository
│   │   │   ├── products_travel_repository.go
│   │   │   └── products_travel_repository_contract.go
│   │   └── service
│   │       ├── products_travel_service.go
│   │       ├── products_travel_service_contract.go
│   │       └── products_travel_utility.go
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
├── docs
│   ├── architecture-diagram.png
│   └── structure.md
├── go.mod
├── go.sum
├── main.go
└── router
    ├── http
    │   ├── countries_router.go
    │   ├── init.go
    │   ├── product_router.go
    │   ├── product_travel_router.go
    │   ├── router.go
    │   ├── simple_router.go
    │   └── travel_schedule_router.go
    └── tcp
        ├── init.go
        ├── router.go
        └── simple_router.go


31 directories, 50 files
