# マイグレーションにはflywayを使っている
DBNAME:=calendar_app
DOCKER_DNS:=db
FLYWAY_CONF?=-url=jdbc:mysql://$(DOCKER_DNS):3306/$(DBNAME) -user=root -password=mysql

export DATABASE_DATASOURCE:=root:mysql@tcp($(DOCKER_DNS):3306)/$(DBNAME)
export GOOGLE_APPLICATION_CREDENTIALS=$(HOME)/.config/gcloud/treasure_app_service_account.json

dcu:
	docker-compose up