setup:
	$(MAKE) install-hooks && $(MAKE) create-db

install-hooks:
	cp ./tools/pre-push.sh .git/hooks/pre-push
	chmod 755 .git/hooks/pre-push

create-db:
	sqlite3 orders.db ".read ./tools/create-db.sql"