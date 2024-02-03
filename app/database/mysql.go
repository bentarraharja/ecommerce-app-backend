package database

import (
	"MyEcommerce/app/config"
	"fmt"

	ud "MyEcommerce/features/user/data"
	pd "MyEcommerce/features/product/data"
	cd "MyEcommerce/features/cart/data"
	od "MyEcommerce/features/order/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	// declare struct config & variable connectionString
	// username:password@tcp(hostdb:portdb)/db_name
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&ud.User{}, &pd.Product{}, &cd.Cart{}, &od.Order{}, &od.OrderItem{})

	return DB
}
