#runners.rebuild.server image containing binary representing gRPC server. This binary should be placed in runner container
runners.rebuild.server:
	@docker build -f build/runner.dockerfile -t runner_store:latest .

#runners.rebuild.start.all.rebuild.server server and rebuild runners
runners.rebuild.start.all.rebuild.server: runners.rebuild.server
	@docker compose -f build/docker-compose.runners.yaml up -d --build

#runners.start.all starts all runners containers
runners.start.all:
	@docker compose -f build/docker-compose.runners.yaml up -d

runners.rebuild.start.all:
	@docker compose -f build/docker-compose.runners.yaml up -d --build

runners.stop.all:
	@docker compose -f build/docker-compose.runners.yaml down

runners.start.py:
	@docker compose -f build/docker-compose.runners.yaml up py_runner -d

runners.rebuild.start.py:
	@docker compose -f build/docker-compose.runners.yaml up py_runner -d --build

runners.stop.py:
	@docker compose -f build/docker-compose.runners.yaml down py_runner

runners.start.js:
	@docker compose -f build/docker-compose.runners.yaml up js_runner -d

runners.rebuild.start.js:
	@docker compose -f build/docker-compose.runners.yaml up js_runner -d --build

runners.stop.js:
	@docker compose -f build/docker-compose.runners.yaml down js_runner

runners.start.go:
	@docker compose -f build/docker-compose.runners.yaml up go_runner -d

runners.rebuild.start.go:
	@docker compose -f build/docker-compose.runners.yaml up go_runner -d --build

runners.stop.go:
	@docker compose -f build/docker-compose.runners.yaml down go_runner



