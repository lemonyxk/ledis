/**
* @program: ledis
*
* @description:
*
* @author: lemo
*
* @create: 2021-11-13 18:11
**/

package ledis

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Model struct {
	cache *Client
	name  string
}

func (p *Model) SetHandler(r *Client) *Model {
	p.cache = r
	return p
}

func (p *Model) key(id int) string {
	return p.name + ":" + strconv.Itoa(id)
}

func (p *Model) Get(id int) *redis.StringStringMapCmd {
	return p.cache.Handler.HGetAll(context.Background(), p.key(id))
}

func (p *Model) Delete(id int) *redis.IntCmd {
	return p.cache.Handler.Del(context.Background(), p.key(id))
}

func (p *Model) Create(id int, value map[string]interface{}) *redis.BoolCmd {
	return p.cache.Handler.HMSet(context.Background(), p.key(id), value)
}

func (p *Model) Update(id int, update map[string]interface{}) *redis.BoolCmd {
	return p.cache.Handler.HMSet(context.Background(), p.key(id), update)
}

func (p *Model) Expire(id int, time time.Duration) *redis.BoolCmd {
	return p.cache.Handler.Expire(context.Background(), p.key(id), time)
}

func (p *Model) Exists(id int) *redis.IntCmd {
	return p.cache.Handler.Exists(context.Background(), p.key(id))
}
