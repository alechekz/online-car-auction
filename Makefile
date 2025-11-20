# Vehicle Service
.PHONY: vehicle-local-build, vehicle-build, vehicle-run, vehicle-test, vehicle-testcover, vehicle-lint

vehicle-local-build:
	@go build -o vehicle-service -v ./services/vehicle/cmd/main.go
	@echo "Vehicle service successfully built"

vehicle-build:
	CGO_ENABLED=0 GOOS=linux go build -o /bin/vehicle -v ./services/vehicle/cmd/main.go
	@echo "Vehicle service successfully built"

vehicle-run:
	@VEHICLE_URL=:7071 INSPECTION_URL=:7073 PRICING_URL=:7075 ./vehicle-service

vehicle-test:
	@go test -v ./services/vehicle/...

vehicle-testcover:
	@go test --cover ./services/vehicle/... --coverprofile=testscoverprofile
	@go tool cover -html=testscoverprofile

vehicle-lint:
	golangci-lint run ./services/vehicle/...

vehicle: vehicle-lint vehicle-test vehicle-local-build vehicle-run

# Inspection Service
.PHONY: inspection-local-build, inspection-build, inspection-run, inspection-test, inspection-testcover, inspection-lint

inspection-local-build:
	@go build -o inspection-service -v ./services/inspection/cmd/main.go
	@echo "inspection service successfully built"

inspection-build:
	CGO_ENABLED=0 GOOS=linux go build -o /bin/inspection -v ./services/inspection/cmd/main.go
	@echo "inspection service successfully built"

inspection-run:
	@INSPECTION_HTTP=:7072 INSPECTION_GRPC=:7073 ./inspection-service

inspection-test:
	@go test -v ./services/inspection/...

inspection-testcover:
	@go test --cover ./services/inspection/... --coverprofile=testscoverprofile
	@go tool cover -html=testscoverprofile

inspection-lint:
	golangci-lint run ./services/inspection/...

inspection: inspection-lint inspection-test inspection-local-build inspection-run

# Pricing Service
.PHONY: pricing-local-build, pricing-build, pricing-run, pricing-test, pricing-testcover, pricing-lint

pricing-local-build:
	@go build -o pricing-service -v ./services/pricing/cmd/main.go
	@echo "pricing service successfully built"

pricing-build:
	CGO_ENABLED=0 GOOS=linux go build -o /bin/pricing -v ./services/pricing/cmd/main.go
	@echo "pricing service successfully built"

pricing-run:
	@PRICING_HTTP=:7074 PRICING_GRPC=:7075 INSPECTION_URL=:7073 ./pricing-service

pricing-test:
	@go test -v ./services/pricing/...

pricing-testcover:
	@go test --cover ./services/pricing/... --coverprofile=testscoverprofile
	@go tool cover -html=testscoverprofile

pricing-lint:
	golangci-lint run ./services/pricing/...

pricing: pricing-lint pricing-test pricing-local-build pricing-run

# Common
.PHONY: lint
lint:
	golangci-lint run ./...