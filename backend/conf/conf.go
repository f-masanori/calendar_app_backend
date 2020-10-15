package conf

type Database_ struct {
	Drivername string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
}

var Database Database_

func Init() {
	Database.Drivername = "mysql"
	Database.Host = "mysql_container"
	Database.Port = "3306"
	Database.User = "root"
	Database.Password = "mysql"
	Database.Dbname = "app"
}

func Test() {
	Database.Drivername = "mysql"
	Database.Host = "mysql_container"
	Database.Port = "3306"
	Database.User = "root"
	Database.Password = "mysql"
	Database.Dbname = "app"

}
