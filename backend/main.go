package main

import (
	"flag"
	"golang/calendar/infrastructure/router"
	"golang/conf"
	"log"
	"os"
)

func init() {
	conf.Init()
}
func main() {
	var databaseDatasource string
	var serviceAccountKeyPath string
	var port int
	flag.StringVar(&databaseDatasource, "databaseDatasource", "root:password@tcp(mysql_container:3306)/app", "Should looks like root:password@tcp(hostname:port)/dbname")
	flag.StringVar(&serviceAccountKeyPath, "serviceAccountKeyPath", "", "Path to service account key")
	flag.IntVar(&port, "port", 8080, "Web server port")
	flag.Parse()

	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)
	router.Run(databaseDatasource, serviceAccountKeyPath, port)
}
