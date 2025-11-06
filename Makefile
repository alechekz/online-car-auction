.PHONY: build-vehicle, run-vehicle, vehicle, test-vehicle, testcover-vehicle, lint-vehicle
build-vehicle:
	@go build -o vehicle-service -v ./services/vehicle/cmd/main.go
	@echo "Vehicle service successfully built"

run-vehicle:
	@VEHICLE_DB=localhost VEHICLE_PORT=8081 ./vehicle-service

test-vehicle:
	@go test -v ./services/vehicle/...

testcover-vehicle:
	@go test --cover ./services/vehicle/... --coverprofile=testscoverprofile
	@go tool cover -html=testscoverprofile

lint-vehicle:
	golangci-lint run ./services/vehicle/...

vehicle: lint-vehicle test-vehicle build-vehicle run-vehicle

.PHONY: lint
lint:
	golangci-lint run ./...