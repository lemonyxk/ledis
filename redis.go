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
	"context"

	"github.com/go-redis/redis/v8"
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

func NewModel(client redis.Cmdable) *Client {
	return &Client{Handler: client}
}

type Client struct {
	Handler redis.Cmdable
}

func (client *Client) Transaction(fn func(pipe redis.Cmdable) error) error {
	var pipe = client.Handler.TxPipeline()
	var err = fn(pipe)
	if err != nil {
		return err
	}
	_, err = pipe.Exec(context.Background())
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
			keys, cursor, err = client.Handler.Scan(context.Background(), cursor, key, int64(count)).Result()
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

func NewClient(option *redis.Options) *redis.Client {
	return redis.NewClient(option)
}

func NewCluster(option *redis.ClusterOptions) *redis.ClusterClient {
	return redis.NewClusterClient(option)
}

func NewFailover(option *redis.FailoverOptions) *redis.Client {
	return redis.NewFailoverClient(option)
}
