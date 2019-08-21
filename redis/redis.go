package redis

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"io"
	"reflect"
	"strings"
	"time"
)


type Redigo struct {
	addr string
	redisPool *redis.Pool
	conn redis.Conn
}

// 初始化redis配置
// 创建redis连接池
func MakeRedigo(redisUrl string) *Redigo {
	log.QyLogger.Info("init redis connection", zap.String("addr", redisUrl))
	db := 15
	rg := Redigo{addr:redisUrl}
	rg.redisPool = &redis.Pool{
		MaxIdle:     100,             // 最大空闲连接数
		MaxActive:   10000,           // 最大连接数
		IdleTimeout: 3 * time.Minute, // 空闲连接的最大等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", rg.addr, redis.DialDatabase(db))
			if err != nil {
				panic("连接redis出错：" + err.Error())
			}
			return c, err
		},
	}

	// 验证是否连接成功
	c := rg.redisPool.Get()
	defer c.Close()
	if err := c.Err(); err != nil {
		panic("连接redis失败：" + err.Error())
	}

	//rg.conn = c
	return &rg
}

func (r *Redigo) GetConn() *Redigo {
	if r.redisPool == nil {
		panic("没有连接到redis，请先调用MakeRedigo()方法")
	}
	r.conn = r.redisPool.Get()
	return r
}

func (r *Redigo) SetHash(table string, kv ...interface{}) error {
	r.GetConn()
	defer r.Close()
	tabWithKv := redis.Args{}.Add(table).AddFlat(kv)
	//tabWithKv := []interface{}{}
	//tabWithKv = append(tabWithKv, table)
	//tabWithKv = append(tabWithKv, kv...)
	_, err := r.conn.Do("HMSET", tabWithKv...)
	return err
}

func (r *Redigo) SetHashStruct(table string, record interface{}) error {
	r.GetConn()
	defer r.Close()
	tabWithKv := redis.Args{}.Add(table).AddFlat(record)
	//tabWithKv := []interface{}{}
	//tabWithKv = append(tabWithKv, table)
	//tabWithKv = append(tabWithKv, kv...)
	_, err := r.conn.Do("HMSET", tabWithKv...)
	return err
}

func (r *Redigo) GetHashStruct(table string, content interface{}) error {
	r.GetConn()
	defer r.Close()
	v, err := redis.Values(r.conn.Do("HGETALL", table))
	if err != nil {
		return err
	}
	fmt.Println(v, reflect.TypeOf(v[0]), string(v[0].([]byte)))
	err = redis.ScanStruct(v, content)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redigo) SetStr(key, value string) error {
	r.GetConn()
	defer r.Close()
	_, err := r.conn.Do("SET", key, value)
	r.conn.Flush()
	if IsConnError(err) {
		r.ReConnect()
		// do again
		_, err = r.conn.Do("SET", key, value)
		r.conn.Flush()
	}
	return err
}

func (r *Redigo) GetStr(key string) (string, error) {
	r.GetConn()
	defer r.Close()
	v, err := redis.String(r.conn.Do("GET", key))
	if IsConnError(err) {
		r.ReConnect()
		// do again
		v, err = redis.String(r.conn.Do("GET", key))
	}
	return v, err
}

func (r *Redigo) Close() {
	r.conn.Close()
}

func IsConnError(err error) bool {
	var needNewConn bool

	if err == nil {
		return false
	}

	if err == io.EOF {
		needNewConn = true
	}
	if strings.Contains(err.Error(), "use of closed network connection") {
		needNewConn = true
	}
	if strings.Contains(err.Error(), "connect: connection refused") {
		needNewConn = true
	}
	return needNewConn
}

func (r *Redigo) ReConnect() {
	c, _ := r.redisPool.Dial()
	if c == nil {
		return
	}
	r.conn = c
}

func (r *Redigo) Invalidate(key string) error {
	r.GetConn()
	defer r.Close()
	_, err := r.conn.Do("DEL", key)
	return err
}