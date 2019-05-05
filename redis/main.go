package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

type Conn struct {
	redis.Conn
}

func main() {
	this := Conn{}
	this.ConnectionRedis()
	this.SetString("yzt", "yzt")
	fmt.Println(this.GetString("yzt"))
	this.Close()
}

//连接redis
func (this *Conn) ConnectionRedis() {
	this.Conn, _ = redis.Dial("tcp", "127.0.0.1:6379") //连接
	_, err := this.Conn.Do("auth", "xltevil")
	if err != nil {
		this.Close()
	}
}

//关闭连接
func (this *Conn) Close() (err error) {
	err = this.Conn.Close()
	os.Exit(1)
	return
}

//redis string
func (this *Conn) SetString(key, value string) {
	_, err := this.Conn.Do("SET", key, value)
	if err != nil {
		this.Close()
	}
}

func (this *Conn) GetString(key string) interface{} {
	result, err := this.Conn.Do("GET", key)
	if err != nil {
		this.Close()
	}
	return result
}

//redis list
func (this *Conn) SetList(key, value string, flag bool) { //true l false r
	if flag == true {
		_, err := this.Conn.Do("lpush", key, value)
		if err != nil {
			this.Close()
		}
	} else {
		this.Conn.Do("rpush", key, value)
	}
}

func (this *Conn) GetList(key string, start, end int) interface{} {
	result, err := this.Conn.Do("lrange", key, start, end)
	if err != nil {
		this.Close()
	}
	return result
}

//redis set
func (this *Conn) Setset(key string, value ...interface{}) {
	_, err := this.Conn.Do("sadd", key, value)
	if err != nil {
		this.Close()
	}
}

func (this *Conn) Getset(key string) interface{} {
	result, err := this.Conn.Do("semeber", key)
	if err != nil {
		this.Close()
	}
	return result
}

//redis hash
func (this *Conn) SetHash(key string, value ...interface{}) {
	_, err := this.Conn.Do("HMSET", key, value)
	if err != nil {
		this.Close()
	}
}

func (this *Conn) GetHash(key string) interface{} {
	result, err := this.Conn.Do("HMGETALL", key)
	if err != nil {
		this.Close()
	}
	return result
}

//redis zset
func (this *Conn) Setzset(key, value string, score int) {
	_, err := this.Conn.Do("zadd", key, score, value)
	if err != nil {
		this.Close()
	}
}

func (this *Conn) Getzset(key string, index1, index2 int) interface{} {
	result, err := this.Conn.Do("ZRANGEBYSCORE", key, index1, index2)
	if err != nil {
		this.Close()
	}
	return result
}
