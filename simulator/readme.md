# How to test
1. run app - `go run main.go`
2. run kafka - `docker-compose -f .docker/kafka/docker-compose.yml -d up`
3. get into kafka container - `docker exec -it kafka_kafka_1 bash`
4. start producer - `kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction`
5. start consumer - `kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal`
6. send message through producer - `{"clientId":"1","routeId":"1"}`
7. See the magic happening
