package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserId    int    `orm:"PK;auto"`
	UserName  string `orm:"unique"`
	PassWord  string
	PublicKey string
}

//仓单
type Receipt struct {
	ReceiptId      int    `orm:"PK;auto"` //仓单编号
	Class          string //品种
	ProductionDate string //产期
	Level          string //等级
	Warehouse      string //仓库
	Provenance     string //产地
}

//用户所拥有的仓单
type UserReceipt struct {
	Id                int //id 主键
	UserId            int //用户id
	ReceiptId         int //仓单编号
	TotalQuantity     int //仓单总量
	RemainingQuantity int //剩余数量
	FrozenQuantity    int //冻结数量
}

//资金表
type UserFunds struct {
	UserId         int `orm:"PK"` //用户id
	TotalFunds     int //资金总量
	RemainingFunds int //剩余资金
	FrozenFunds    int //冻结资金
}

var o orm.Ormer //定义Ormer接口变量

func init() {
	//连接数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/receipt_trade?charset=utf8", 30)
	//注册声明的model
	//orm.RegisterModel(new(User), new(Receipt), new(UserReceipt), new(UserFunds))
	orm.RegisterModel(new(User), new(Receipt), new(UserReceipt), new(UserFunds))
	//创建表
	orm.RunSyncdb("default", false, true)

	o = orm.NewOrm() //创建orm结构体实例
}

func (user User) Insert() {
	id, err := o.Insert(&user)
	if err == nil {
		fmt.Printf("插入数据库 id = %d\n", id)
	} else {
		fmt.Printf("插入数据库失败err: %v\n", err)
	}
}

func (receipt Receipt) Insert() {
	id, err := o.Insert(&receipt)
	if err == nil {
		fmt.Printf("插入数据库 id = %d\n", id)
	} else {
		fmt.Printf("插入数据库失败err: %v\n", err)
	}
}
func (userRct UserReceipt) Insert() {
	id, err := o.Insert(&userRct)
	if err == nil {
		fmt.Printf("UserReceipt.Insert() 为用户添加仓单成功 id = %d\n", id)
	} else {
		fmt.Printf("UserReceipt.Insert() 为用户添加仓单失败 err: %v\n", err)
	}
}

func (userFus UserFunds) Insert() {
	id, err := o.Insert(&userFus)
	if err == nil {
		fmt.Printf("UserFunds 插入数据库 id = %d\n", id)
	} else {
		fmt.Printf("UserFunds 插入数据库失败err: %v\n", err)
	}
}
