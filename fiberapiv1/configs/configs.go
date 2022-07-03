package configs

import (
	"fiberapiv1/entity"
	"fiberapiv1/logs"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	//"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func InitTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func InitDatabase() *gorm.DB {
	//for mysql
	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
	// 	viper.GetString("db.username"),
	// 	viper.GetString("db.password"),
	// 	viper.GetString("db.host"),
	// 	viper.GetInt("db.port"),
	// 	viper.GetString("db.database"),
	// )
	dsn := fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	//dsn := "sqlserver://sa:12345Cat@localhost:1433?database=banking"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{}) //for sql server
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //for mysql
	if err != nil {
		logs.Info("Open db err : " + err.Error())
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{}, &entity.ProductGroup{}, &entity.Product{}, &entity.Orders{})
	if err != nil {
		panic(err)
	}

	Database = DbInstance{
		Db: db,
	}
	logs.Info("Database Connect : Success!")
	// defer db.Close()
	return db
}

//Json Token SecretKey
func GetSecretKey() string {
	return viper.GetString("jwt.secretKey")
}

//OAuth google
func GetGoogleClientID() string {
	return viper.GetString("oAuth.google.clientID")
}
func GetGoogleClientSecret() string {
	return viper.GetString("oAuth.google.clientSecret")
}

//OAuth line
func GetLineClientID() string {
	return viper.GetString("oAuth.line.clientID")
}
func GetLineClientSecret() string {
	return viper.GetString("oAuth.line.clientSecret")
}
