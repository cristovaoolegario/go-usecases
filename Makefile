# 'ps' get information for all processes incl. command
# 'grep' is filtering for your command
# 'tr' is removing duplicate spaces for cut
# 'cut' is getting you the column with the PID
PROCESS_LIST := $(shell ps | grep go | grep -v grep | tr -s ' ' | cut -d\  -f1)

setup:
	$(MAKE) install-hooks && $(MAKE) create-db && $(MAKE) install-deps && mkdir -p logs

install-hooks:
	cp ./tools/pre-push.sh .git/hooks/pre-push
	chmod 755 .git/hooks/pre-push

create-db:
	sqlite3 orders.sqlite ".read ./tools/create-db.sql"

install-deps:
	go mod tidy
	go get ./...

tests:
	go clean -testcache
	go test -v ./...

test-coverage:
	go test -coverprofile cover.out ./... && go tool cover -html=cover.out

run:
	docker-compose up -d && sleep 30 && $(MAKE) producer && $(MAKE) consumer && open http://localhost:8080/total

producer:
	go run cmd/producer/main.go > logs/output.producer.log 2>logs/error.producer.log&

consumer:
	go run cmd/consumer/main.go > logs/output.consumer.log 2>logs/error.consumer.log&

list-processes:
	@echo $(PROCESS_LIST)

kill:
# if there is some app process runnig, kill it
	@$(if $(strip $(PROCESS_LIST)), kill -9 $(PROCESS_LIST)) 	
	docker-compose down

build-image:
	docker build -t cristovaoolegario/go-usecases:latest -f Dockerfile.prod .