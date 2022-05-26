/**
* @program: ledis
*
* @description:
*
* @author: lemo
*
* @create: 2020-05-07 16:58
**/

package main

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/lemonyxk/ledis"
)

func main() {

	var client = ledis.NewFailover(&redis.FailoverOptions{
		MasterName:    "master",
		Password:      "1354243",
		SentinelAddrs: []string{"192.168.0.100:16379"},
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

}
