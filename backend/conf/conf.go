package conf

/* 設定情報はアプリ実行時に挿入されるが、テスト用のために残している */
type Database_ struct {
	Drivername string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
}

var Database Database_

func Test() {
	Database.Drivername = "mysql"
	Database.Host = "calendar-app-db"
	Database.Port = "3306"
	Database.User = "root"
	Database.Password = "mysql"
	Database.Dbname = "calendar_app"

}
