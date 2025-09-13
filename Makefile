COMPOSE_FILE = build/docker-compose.runners.yaml


#runners.rebuild.server image containing binary representing gRPC server. This binary should be placed in runner container
runners.rebuild.server:
	@docker build -f build/runner.dockerfile -t runner_store:latest .

#runners.rebuild.start.all.rebuild.server rebuilds server and runners containers
runners.rebuild.start.all.rebuild.server: runners.rebuild.server
	@docker compose -f ${COMPOSE_FILE} up -d --build

#runners.start.all starts all runners containers
runners.start.all:
	@docker compose -f ${COMPOSE_FILE} up -d

runners.rebuild.start.all:
	@docker compose -f ${COMPOSE_FILE} up -d --build

runners.stop.all:
	@docker compose -f ${COMPOSE_FILE} down

runners.start.py:
	@docker compose -f ${COMPOSE_FILE} up py_runner -d

runners.rebuild.start.py:
	@docker compose -f ${COMPOSE_FILE} up py_runner -d --build

runners.stop.py:
	@docker compose -f ${COMPOSE_FILE} down py_runner

runners.start.js:
	@docker compose -f ${COMPOSE_FILE} up js_runner -d

runners.rebuild.start.js:
	@docker compose -f ${COMPOSE_FILE} up js_runner -d --build

runners.stop.js:
	@docker compose -f ${COMPOSE_FILE} down js_runner

runners.start.go:
	@docker compose -f ${COMPOSE_FILE} up go_runner -d

runners.rebuild.start.go:
	@docker compose -f ${COMPOSE_FILE} up go_runner -d --build

runners.stop.go:
	@docker compose -f ${COMPOSE_FILE} down go_runner

runners.start.cs:
	@docker compose -f ${COMPOSE_FILE} up cs_runner -d

runners.rebuild.start.cs:
	@docker compose -f ${COMPOSE_FILE} up cs_runner -d --build

runners.stop.cs:
	@docker compose -f ${COMPOSE_FILE} down cs_runner

runners.start.cpp:
	@docker compose -f ${COMPOSE_FILE} up cpp_runner -d

runners.rebuild.start.cpp:
	@docker compose -f ${COMPOSE_FILE} up cpp_runner -d --build

runners.stop.cpp:
	@docker compose -f ${COMPOSE_FILE} down cpp_runner

