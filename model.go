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
	Cache *Client
	Name  string
}

func (p *Model) SetHandler(r *Client) *Model {
	p.Cache = r
	return p
}

func (p *Model) Key(id int) string {
	return p.Name + ":" + strconv.Itoa(id)
}

func (p *Model) Get(id int) *redis.StringStringMapCmd {
	return p.Cache.Handler.HGetAll(context.Background(), p.Key(id))
}

func (p *Model) Delete(id int) *redis.IntCmd {
	return p.Cache.Handler.Del(context.Background(), p.Key(id))
}

func (p *Model) Create(id int, value map[string]interface{}) *redis.BoolCmd {
	return p.Cache.Handler.HMSet(context.Background(), p.Key(id), value)
}

func (p *Model) Update(id int, update map[string]interface{}) *redis.BoolCmd {
	return p.Cache.Handler.HMSet(context.Background(), p.Key(id), update)
}

func (p *Model) Expire(id int, time time.Duration) *redis.BoolCmd {
	return p.Cache.Handler.Expire(context.Background(), p.Key(id), time)
}

func (p *Model) Exists(id int) *redis.IntCmd {
	return p.Cache.Handler.Exists(context.Background(), p.Key(id))
}
