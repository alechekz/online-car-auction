# Generate gRPC code
protoc -I=services/inspection/delivery/grpc/proto \
  --go_out=paths=source_relative:services/inspection/delivery/grpc/proto \
  --go-grpc_out=paths=source_relative:services/inspection/delivery/grpc/proto \
  services/inspection/delivery/grpc/proto/inspection.proto

protoc -I=services/pricing/delivery/grpc/proto \
  --go_out=paths=source_relative:services/pricing/delivery/grpc/proto \
  --go-grpc_out=paths=source_relative:services/pricing/delivery/grpc/proto \
  services/pricing/delivery/grpc/proto/pricing.proto





# vehicle service API examples
curl -i -X POST http://localhost:8081/vehicles \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","year":2022,"odometer":12000}'

curl -i http://localhost:8081/vehicles

curl -i http://localhost:8081/vehicles/5YJSA1E26MF168123

curl -i -X PUT http://localhost:8081/vehicles/5YJSA1E26MF168123 \
  -H "Content-Type: application/json" \
  -d '{
    "vin":"5YJSA1E26MF168123",
    "year":1999,
    "odometer":125000,
    "exteriorColor":"Red",
    "interiorColor":"Black"
  }'

curl -i -X DELETE http://localhost:8081/vehicles/5YJSA1E26MF168123

# Inspection Service API examples
curl -i http://localhost:8082/inspections/get-build-data/5YJSA1E26MF168123

curl -i -X POST http://localhost:8082/inspections/inspect \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","year":2022}'

# Pricing Service API examples
curl -i -X POST http://localhost:8084/pricing/get-recommended-price \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","grade":47,"odometer":30000}'

# Bulk Vehicle example
curl -i -X POST http://localhost:8081/vehicles \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168300","year":2022,"odometer":1000}'

curl -X POST http://localhost:8081/vehicles/bulk \
  -H "Content-Type: application/json" \
  -d '{
      "vehicles": [
        {"vin":"1HGCM82633A004352","year":2018,"odometer":45000},
        {"vin":"1FTFW1E50JFA12345","year":2020,"odometer":15000}
      ]
}'
