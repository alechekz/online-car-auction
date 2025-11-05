.PHONY: build-vehicle, run-vehicle, vehicle, lint, inspection, pricing

build-vehicle:
	@go build -o vehicle-service -v ./services/vehicle/cmd/main.go
	@echo "Vehicle service successfully built"

run-vehicle:
	@VEHICLE_DB=localhost VEHICLE_PORT=8081 ./vehicle-service

vehicle: build-vehicle run-vehicle

lint:
	golangci-lint run ./...