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

type model struct {
	cache *Client
	name  string
}

func (p *model) SetHandler(r *Client) *model {
	p.cache = r
	return p
}

func (p *model) key(id int) string {
	return p.name + ":" + strconv.Itoa(id)
}

func (p *model) Get(id int) *redis.StringStringMapCmd {
	return p.cache.Handler.HGetAll(context.Background(), p.key(id))
}

func (p *model) Delete(id int) *redis.IntCmd {
	return p.cache.Handler.Del(context.Background(), p.key(id))
}

func (p *model) Create(id int, value map[string]interface{}) *redis.BoolCmd {
	return p.cache.Handler.HMSet(context.Background(), p.key(id), value)
}

func (p *model) Update(id int, update map[string]interface{}) *redis.BoolCmd {
	return p.cache.Handler.HMSet(context.Background(), p.key(id), update)
}

func (p *model) Expire(id int, time time.Duration) *redis.BoolCmd {
	return p.cache.Handler.Expire(context.Background(), p.key(id), time)
}

func (p *model) Exists(id int) *redis.IntCmd {
	return p.cache.Handler.Exists(context.Background(), p.key(id))
}
