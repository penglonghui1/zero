package queue

import "time"

type Config struct {
	EtcdAddress          string //连接地址
	NSQAddress           string //连接地址
	MaxInFlight          int    //最大NSQ消费并发数
	MaxNSQGoroutineCount int    //NSQ用来处理消息的协程数
	RedisPoolSize        int    //Redis连接池
	RedisMinIdleConns    int    //最小连接空闲数
	MySQLMaxOpenConns    int
	MySQLMaxIdleConns    int
	MySQLMaxLifetime     time.Duration
}

type Options func(parameter *Config)
