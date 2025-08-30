py.runner.start:
	@docker compose -f build/runners/docker-compose.runners.yaml up

py.runner.rebuild:
	@docker compose -f build/runners/docker-compose.runners.yaml up --build