package database

import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
)

const Dbname = ""
const Dbhost = "127.0.0.1"
const Dbport = "8000" //default is 5432
const Dbuser = "postgres" //default is postgres
const Dbpassword = ""

func Connect()(*gorm.DB, error){

	return gorm.Open("postgres","host="+Dbhost+" port="+Dbport+" user="+Dbuser+" dbname="+Dbname+" password="+Dbpassword+" sslmode=disable")
}