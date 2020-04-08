/**
* @program: ledis
*
* @description:
*
* @author: lemo
*
* @create: 2020-04-08 17:18
**/

package ledis

import (
	"time"

	"github.com/go-redis/redis/v7"
)

/**
* @program: center-server-go
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-01 14:06
**/

func NewHandler(client Handler) *Client {
	return &Client{Handler: client}
}

type Client struct {
	Handler Handler
}

func (client *Client) Transaction(fn func(pipeClient *Client) error) error {
	var pipe = client.Handler.TxPipeline()
	var err = fn(NewHandler(pipe))
	if err != nil {
		return err
	}
	_, err = pipe.Exec()
	return err
}

type ScanResult struct {
	err error
	res string
}

func (res *ScanResult) LastError() error {
	return res.err
}

func (res *ScanResult) Result() string {
	return res.res
}

func (client *Client) Scan(key string, count int) chan *ScanResult {

	var ch = make(chan *ScanResult, 1)

	var cursor uint64 = 0

	go func() {
		for {

			var keys []string
			var err error
			keys, cursor, err = client.Handler.Scan(cursor, key, int64(count)).Result()
			if err != nil {
				ch <- &ScanResult{err: err}
				close(ch)
				return
			}

			for i := 0; i < len(keys); i++ {
				ch <- &ScanResult{res: keys[i]}
			}

			if cursor == 0 {
				break
			}
		}
		close(ch)
	}()

	return ch
}

type Handler interface {
	HMSet(key string, values ...interface{}) *redis.BoolCmd
	Del(keys ...string) *redis.IntCmd
	HGetAll(key string) *redis.StringStringMapCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	HIncrByFloat(key, field string, incr float64) *redis.FloatCmd
	Exists(keys ...string) *redis.IntCmd
	FlushAll() *redis.StatusCmd
	LPush(key string, values ...interface{}) *redis.IntCmd
	TxPipeline() redis.Pipeliner
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Scan(cursor uint64, match string, count int64) *redis.ScanCmd
	Do(args ...interface{}) *redis.Cmd
}

func NewClient(option *redis.Options) *redis.Client {
	return redis.NewClient(option)
}

func NewCluster(option *redis.ClusterOptions) *redis.ClusterClient {
	return redis.NewClusterClient(option)
}

func NewFailover(option *redis.FailoverOptions) *redis.Client {
	return redis.NewFailoverClient(option)
}
