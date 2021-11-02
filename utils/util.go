package utils

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"html/template"
	"github.com/russross/blackfriday"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"github.com/sourcegraph/syntaxhighlight"
	"crypto/md5"
	"time"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql....")
	driverName, _ := beego.AppConfig.String("driverName")

	//注册数据库驱动
	//orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user ,_:= beego.AppConfig.String("mysqluser")
	pwd,_ := beego.AppConfig.String("mysqlpwd")
	host,_ := beego.AppConfig.String("host")
	port ,_:= beego.AppConfig.String("port")
	dbname,_ := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		//创建用户表
		CreateTableWithUser()
		//创建文章表
		CreateTableWithArticle()
		//创建相册数据表
		CreateTableWithAlbum()
	}
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

/**
 * 将文章详情的内容，转换成HTMl语句
 */
func SwitchMarkdownToHtml(content string) template.HTML {

	markdown := blackfriday.MarkdownCommon([]byte(content))

	//获取到html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	 */
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}


//传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

//将传入的时间戳转为时间
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
