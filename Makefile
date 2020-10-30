# マイグレーションにはflywayを使っている
DBNAME:=calendar_app
DOCKER_DNS:=db
FLYWAY_CONF?=-url=jdbc:mysql://$(DOCKER_DNS):3306/$(DBNAME) -user=root -password=mysql

export DATABASE_DATASOURCE:=root:mysql@tcp($(DOCKER_DNS):3306)/$(DBNAME)
export GOOGLE_APPLICATION_CREDENTIALS=$(HOME)/.config/gcloud/calendarapp-dcbd5-firebase-adminsdk-9fhv7-acb24a0067.json

dcu:
	docker-compose up

DB_SERVICE:=db
mysql/client:
	docker-compose exec $(DB_SERVICE) mysql -uroot -hlocalhost -pmysql $(DBNAME)

mysql/init:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -pmysql \
		-e "create database \`$(DBNAME)\`"

__mysql/drop:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -pmysql \
		-e "drop database \`$(DBNAME)\`"

MIGRATION_SERVICE:=migration
flyway/info:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) info

flyway/migrate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) migrate

flyway/repair:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) repair

flyway/baseline:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) baseline