package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//学习mgo的基本使用方法,mongodb 不是内存数据数据库 ，一般存储与关系无关的数据 mongodb的速度大于mysql

const URL = "localhost:27017"

type Person struct {
	NAME  string `bson:"name"`
	PHONE string `bson:"phone"`
}

type Men struct {
	Person []Person
}

var (
	session    *mgo.Session    //会话
	db         *mgo.Database   //数据库
	collection *mgo.Collection //集合
	err        error
)

func main() {
	GetSession()
	fmt.Println("连接成功")
	SelectDB("GODs")
	fmt.Println("连接数据库成功！")
	SelectCollection("person")
	fmt.Println("得到集合")
	ElementCount()
	temps := make([]Person, 3)
	one := Person{"123", "123"}
	temps[0] = one
	temps[1] = one
	temps[2] = one
	Insert(temps)
	fmt.Println("增加元素成功")
	ElementCount()
	result := Person{}
	s := FindOneByBosnM("name", "123", result)
	fmt.Println("找出的数据是", s)
	ss := FindAll()
	for ss.Next(&result) {
		fmt.Println(result)
	}
	err = collection.Update(bson.M{"name": "123"}, bson.M{"$set": bson.M{"name": "ddd"}})
	fmt.Println("修改成功")
	ssss := FindOneByBosnM("name", "ddd", result)
	fmt.Println(ssss)
	Remove("name", "123")
	fmt.Println("删除数据完毕")
	ElementCount()
	session.Close()
}

//得到会话  会话表示与数据库的通信会话,会话相当与一个会场，比如公司开会先要通知人，然后在指定的场所开会
func GetSession() {
	session, err = mgo.Dial(URL) //拨号连接数据库
	Exit(err)
}

//选择数据库
func SelectDB(dbname string) *mgo.Database { //会场里面可以有很多的会议
	db = session.DB(dbname)
	return db
}

//选择指定数据库里的集合 如果集合已经存在，则直接返回,
func SelectCollection(name string) *mgo.Collection {
	collection = db.C(name)
	return collection
}

//查看选中的集合有多少元素
func ElementCount() {
	n, err := collection.Count()
	Exit(err)
	fmt.Println("该集合共有", n, "个元素")
}

//插入元素
func Insert(s []Person) {
	for _, ss := range s {
		err = collection.Insert(ss)
	}
	Exit(err)
}

//BSON 是一种类json的二进制存储格式
//BSON.M{}键值对
//查询单条数据
func FindOneByBosnM(key, value string, result interface{}) interface{} {
	err = collection.Find(bson.M{key: value}).One(&result)
	Exit(err)
	return result
}

//查询多条语句
func FindAll() *mgo.Iter {
	iter := collection.Find(nil).Iter() //nil查全部
	return iter
}

//移除数据
func Remove(key, value string) {
	_, err = collection.RemoveAll(bson.M{key: value})
	Exit(err)
}

//错误直接结束程序
func Exit(err error) {
	if err != nil {
		fmt.Println("程序失败，退出")
		panic(err)
	}
}
