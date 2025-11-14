# Vehicle Service
.PHONY: vehicle-local-build, vehicle-build, vehicle-run, vehicle-test, vehicle-testcover, vehicle-lint

vehicle-local-build:
	@go build -o vehicle-service -v ./services/vehicle/cmd/main.go
	@echo "Vehicle service successfully built"

vehicle-build:
	CGO_ENABLED=0 GOOS=linux go build -o /bin/vehicle -v ./services/vehicle/cmd/main.go
	@echo "Vehicle service successfully built"

vehicle-run:
	@VEHICLE_URL=:7071 ./vehicle-service

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
	@INSPECTION_URL=:7072 ./inspection-service

inspection-test:
	@go test -v ./services/inspection/...

inspection-testcover:
	@go test --cover ./services/inspection/... --coverprofile=testscoverprofile
	@go tool cover -html=testscoverprofile

inspection-lint:
	golangci-lint run ./services/inspection/...

inspection: inspection-lint inspection-test inspection-local-build inspection-run

# Common
.PHONY: lint
lint:
	golangci-lint run ./...