# マイグレーションにはflywayを使っている
DBNAME:=calendar_app
DOCKER_DNS:=db
FLYWAY_CONF?=-url=jdbc:mysql://$(DOCKER_DNS):3306/$(DBNAME) -user=root -password=mysql