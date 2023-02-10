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
	"log"

	"github.com/lemonyxk/ledis"
	"github.com/redis/go-redis/v9"
)

func main() {

	var client = ledis.NewFailover(&redis.FailoverOptions{
		MasterName:    "master",
		Password:      "1354243",
		SentinelAddrs: []string{"127.0.0.1:16379", "127.0.0.1:16380", "127.0.0.1:16381"},
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	var client1 = ledis.NewFailover(&redis.FailoverOptions{
		MasterName:    "master",
		Password:      "1354243",
		SentinelAddrs: []string{"127.0.0.1:16379", "127.0.0.1:16380", "127.0.0.1:16381"},
	})

	err = client1.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	var handler = ledis.NewCmd(client)

	log.Println(handler.HGetAll(context.Background(), "ACCOUNT:100013643").String())
}
