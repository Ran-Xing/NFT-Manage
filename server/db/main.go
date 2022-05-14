package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	db  *gorm.DB
	err error
)

type (
	AddressList struct {
		ID        uint   `gorm:"primaryKey"`
		Name      string `json:"name" gorm:"comment:名称"`
		Domain    string `json:"domain" gorm:"comment:域名"`
		Address   string `json:"address" gorm:"comment:URL地址"`
		Logo      string `json:"logo" gorm:"comment:图标地址"`
		CreatedAt time.Time
		UpdatedAt time.Time
		//DeletedAt gorm.DeletedAt
	}

	CookiesList struct {
		ID   uint   `gorm:"primaryKey"`
		Data string `json:"data" gorm:"comment:Cookie数据"`
	}
)

func init() {
	shellFolder, _ := os.Getwd()
	if _, err = os.Stat(shellFolder); err != nil {
		if err = os.Mkdir(shellFolder, os.ModePerm); err != nil {
			log.Errorf("创建目录失败")
			return
		}
	}
	if db, err = gorm.Open(sqlite.Open(shellFolder+"/nft.db"), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		log.Errorf("数据库连接失败")
		return
	}
}

func GetMenu(address string) AddressList {
	var (
		domain      string
		addressData *url.URL
		addressList AddressList
		num         int64
	)
	// Create a new table
	if addressData, err = url.Parse(address); err != nil {
		panic(err)
	}
	domain = addressData.Host
	log.Infof("domain: %s", domain)

	// Create a new table
	if err = db.AutoMigrate(&AddressList{}); err != nil {
		log.Errorf("自动创建表失败, %s", err)
		return AddressList{}
	}
	if err = db.AutoMigrate(&CookiesList{}); err != nil {
		log.Errorf("自动创建表失败, %s", err)
		return AddressList{}
	}
	//db.Create(&AddressList{Name: "test", Domain: domain, Address: address, Logo: "test"})
	//db.Create(&CookiesList{ID: 1, Data: "test"})
	if err = db.Model(&AddressList{}).Where("domain = ?", domain).Count(&num).Scan(&addressList).Error; err != nil {
		log.Errorf("查询失败, %s", err)
		return AddressList{}
	}
	switch num {
	case 1:
		log.Infof("查询菜单成功, %s", "数据库中存在数据")
		return addressList
	case 0:
		log.Errorf("查询菜单成功, %s", "数据库中不存在数据")
		return AddressList{}
	default:
		log.Errorf("查询菜单成功, %s", "数据库中存在多条数据")
		return AddressList{}
	}
}

func AddMenu(name, address, logo string) {
	// Create a new table
}

func GetCookies(address string) []string {
	var (
		cookies []string
		menus   AddressList
	)
	if menus = GetMenu(address); menus == (AddressList{}) {
		log.Errorf("查询菜单失败, %s", "数据库中不存在数据")
		return []string{}
	}
	if err = db.Model(&CookiesList{}).Where("id = ?", strconv.Itoa(int(menus.ID))).Pluck("data", &cookies).Error; err != nil {
		log.Errorf("查询Cookies失败, %s", err)
		return nil
	}
	log.Infof("查询Cookies成功, Cookies数量: %v", len(cookies)) // 返回Cookies数量
	return cookies
}

func Addcookies(address string, cookie string) bool {
	var (
		cookies []string
		menus   AddressList
	)
	if menus = GetMenu(address); menus == (AddressList{}) {
		log.Errorf("查询菜单失败, %s", "数据库中不存在数据")
		return false
	}
	if cookies = GetCookies(address); len(cookies) == 0 {
		log.Errorf("查询Cookies失败, %s", "数据库中不存在数据")
	}
	for i := range cookies {
		if cookies[i] == cookie {
			log.Errorf("查询Cookies失败, %s", "数据库中已存在数据")
			return false
		}
	}
	if err = db.Create(&CookiesList{ID: menus.ID, Data: cookie}).Error; err != nil {
		log.Errorf("添加Cookies失败, %s", err)
		return false
	}
	return true
}
