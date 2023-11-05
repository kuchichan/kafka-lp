.PHONY: docker/up
docker/up:
	@echo "Starting docker..."
	docker-compose up -d --remove-orphans

.PHONY: docker/create-topics
docker/create-topics:
	@echo "Create topics..."
	docker exec -it kafka-lp_kafka_1 kafka-topics.sh --create --topic order-received --bootstrap-server localhost:9092
	docker exec -it kafka-lp_kafka_1 kafka-topics.sh --create --topic order-confirmed --bootstrap-server localhost:9092
	docker exec -it kafka-lp_kafka_1 kafka-topics.sh --create --topic order-fulfilled --bootstrap-server localhost:9092
	docker exec -it kafka-lp_kafka_1 kafka-topics.sh --create --topic notif-sent --bootstrap-server localhost:9092
	docker exec -it kafka-lp_kafka_1 kafka-topics.sh --create --topic event-not-processed --bootstrap-server localhost:9092

	# change retention for order-received topic to three days
	@echo "Alter 'order-received-topic..."
	docker exec -it kafka-lp_kafka_1 kafka-configs.sh --alter --add-config retention.ms=259200000 --topic order-received --bootstrap-server localhost:9092

	@echo "List created topics..."
	docker exec -it kafka-lp_kafka_1 kafka-topics.sh --bootstrap-server localhost:9092 --list

