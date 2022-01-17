package SQL

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	IP       string
	Port     int
	UserName string
	PassWord string
	DataBase string
)

type Yaml struct {
	Clickhouse struct {
		IP        string `yaml:"ip"`
		Port      int    `yaml:"port"`
		UserName  string `yaml:"username"`
		PassWord  string `password:"port"`
		DataBase  string `database:"port"`
		TableName string `tablename:"port"`
	}
	Mysql struct {
		IP        string `yaml:"ip"`
		Port      int    `yaml:"port"`
		UserName  string `yaml:"username"`
		PassWord  string `password:"port"`
		DataBase  string `database:"port"`
		TableName string `tablename:"port"`
	}
}

const yamlroutie = "./DataBase/yaml/clickhouse.yml"

func InitConfig() {
	var tmp Yaml
	config, err := ioutil.ReadFile(yamlroutie)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(config))

	err1 := yaml.Unmarshal(config, &tmp)

	if err1 != nil {
		log.Print(err1)
	}

	IP = tmp.Mysql.IP
	Port = tmp.Mysql.Port
	UserName = tmp.Mysql.UserName
	PassWord = tmp.Mysql.PassWord
	DataBase = tmp.Mysql.DataBase

}

func Connect() *gorm.DB {
	InitConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, PassWord, IP, Port, DataBase)

	//dsn := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s", IP, Port, UserName, PassWord, DataBase)
	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   true,
			},
		})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connet success DataBase:", DataBase)
	}

	return db
}
