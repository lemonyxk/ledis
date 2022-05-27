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
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/lemonyxk/ledis"
)

func main() {

	var client = ledis.NewFailover(&redis.FailoverOptions{
		MasterName:    "master",
		Password:      "1354243",
		SentinelAddrs: []string{"192.168.0.100:16379", "192.168.0.100:16380", "192.168.0.100:16381"},
	})

	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}

	var client1 = ledis.NewFailover(&redis.FailoverOptions{
		MasterName:    "master",
		Password:      "1354243",
		SentinelAddrs: []string{"192.168.0.100:16379", "192.168.0.100:16380", "192.168.0.100:16381"},
	})

	err = client1.Ping().Err()
	if err != nil {
		panic(err)
	}

	var handler = ledis.NewCmd(client)

	log.Println(handler.HGetAll("ACCOUNT:100013643").String())
}
