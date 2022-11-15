setup:
	$(MAKE) install-hooks && $(MAKE) create-db && $(MAKE) install-deps

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
	docker-compose up -d