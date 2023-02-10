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

	"github.com/redis/go-redis/v9"
)

type Model struct {
	Handler Cmdable
	Name    string
}

func NewModel(name string) *Model {
	return &Model{Name: name}
}

func (p *Model) SetHandler(r Cmdable) *Model {
	p.Handler = r
	return p
}

func (p *Model) Key(id int) string {
	return p.Name + ":" + strconv.Itoa(id)
}

func (p *Model) Get(id int) *redis.MapStringStringCmd {
	return p.Handler.HGetAll(context.Background(), p.Key(id))
}

func (p *Model) Delete(id int) *redis.IntCmd {
	return p.Handler.Del(context.Background(), p.Key(id))
}

func (p *Model) Create(id int, value map[string]interface{}) *redis.BoolCmd {
	return p.Handler.HMSet(context.Background(), p.Key(id), value)
}

func (p *Model) Update(id int, update map[string]interface{}) *redis.BoolCmd {
	return p.Handler.HMSet(context.Background(), p.Key(id), update)
}

func (p *Model) Expire(id int, time time.Duration) *redis.BoolCmd {
	return p.Handler.Expire(context.Background(), p.Key(id), time)
}

func (p *Model) Exists(id int) *redis.IntCmd {
	return p.Handler.Exists(context.Background(), p.Key(id))
}
