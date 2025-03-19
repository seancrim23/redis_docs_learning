package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-13594.c283.us-east-1-4.ec2.redns.redis-cloud.com:13594",
		Username: "default",
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	//SET GET UNLINK
	rdb.Set(ctx, "foo", "bar", 0)
	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar

	rdb.Set(ctx, "foo", "updated bar", 0)
	result, err = rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> updated bar

	//rdb.Unlink(ctx, "foo")
	//result, err = rdb.Get(ctx, "foo").Result()

	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(result) // >>> empty string?

	//LIST
	rdb.LPush(ctx, "foo:bar:list", "value1")
	rdb.LPush(ctx, "foo:bar:list", "value2")
	rdb.LPush(ctx, "foo:bar:list", "value3")
	rdb.LPush(ctx, "foo:bar:list", "value4")
	rdb.LPush(ctx, "foo:bar:list", "value5")
	rdb.LPush(ctx, "foo:bar:list", "value6")

	len, err := rdb.LLen(ctx, "foo:bar:list").Result()

	if err != nil {
		panic(err)
	}
	fmt.Println(len) // should be 6

	lRangeResult, err := rdb.LRange(ctx, "foo:bar:list", 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(lRangeResult) // should print first 3 values in practice list

	lRangeResult, err = rdb.LRange(ctx, "foo:bar:list", 0, -1).Result() //-1 indicates end of list
	if err != nil {
		panic(err)
	}
	fmt.Println(lRangeResult) // should print entire list

	lIndexResult, err := rdb.LIndex(ctx, "foo:bar:list", 0).Result() //get first element of list
	if err != nil {
		panic(err)
	}
	fmt.Println(lIndexResult) // should print first element of list

	//SETS
	rdb.SAdd(ctx, "foo:bar:set", "bob", "joe", "tim", "jim")
	sMembsResult, err := rdb.SMembers(ctx, "foo:bar:set").Result() //members of set
	if err != nil {
		panic(err)
	}
	fmt.Println(sMembsResult)                                  // should print all set members
	sCardResult, err := rdb.SCard(ctx, "foo:bar:set").Result() //cardinality (number of elements) of set
	if err != nil {
		panic(err)
	}
	fmt.Println(sCardResult) // should print number of elements of set
	rdb.SRem(ctx, "foo:bar:set", "tim", "jim")
	sMembsResult, err = rdb.SMembers(ctx, "foo:bar:set").Result() //members of set
	if err != nil {
		panic(err)
	}
	fmt.Println(sMembsResult) // should print all set members
	rdb.HSet(ctx, "foo:bar:hash:example", "name", "joe", "color", "blue", "quantity", 20)
	hGetResult, err := rdb.HGet(ctx, "foo:bar:hash:example", "name").Result() //name field from hash
	//you can also do hgetall to get everything from a hash
	if err != nil {
		panic(err)
	}
	fmt.Println(hGetResult)
}
