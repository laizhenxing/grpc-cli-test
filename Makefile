run:
		MICRO_REGISTRY=consul MICRO_REGISTRY_ADDRESS=consul-server:8500 MICRO_SERVER_ADVERTISE=127.0.0.1:9052 go run main.go