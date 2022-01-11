package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

const (
	ADDRESS  string = "10.188.88.193:6379"
	PASSWORD string = ""
	DATABASE int    = 0
)

func main() {
	client := GetNewClient()
	result, err := GetPong(client)
	errorHandler(err)
	fmt.Println(result)
	// res, err := GetValueFromKey(client, "name")
	// errorHandler(err)
	// fmt.Println(res)
	err = SetValueForKey(client, "name", "liurui1")
	errorHandler(err)
	fmt.Println("添加Key完成")
	key := "name"
	res, err := GetValueFromKey(client, key)
	errorHandler(err)
	fmt.Printf("Key:%s的Value是:%s\n", key, res)
	ListOperator(client)
	SetOperator(client)
	HashSetOperator(client)
}

func GetNewClient() *redis.Client {
	var opt = redis.Options{
		Addr:     ADDRESS,
		Password: PASSWORD,
		DB:       DATABASE,
		PoolSize: 10, //连接池数量
	}
	client := redis.NewClient(&opt)
	return client
}

func GetPong(client *redis.Client) (string, error) {
	return client.Ping().Result()
}

func errorHandler(err error) {
	if err != nil {
		logger := log.Default()
		logger.Fatalln(err)
	}
}

/*
set( key, value )：给数据库中名称为 key 的 string 赋予值 value
get( key )：返回数据库中名称为 key 的 string 的 value
getset( key, value )：给名称为 key 的 string 赋予上一次的 value
mget( key1, key2,…, key N )：返回库中多个 string 的 value
setnx( key, value )：添加 string，名称为 key，值为 value
setex( key, time, value )：向库中添加 string，设定过期时间 time
mset( key N, value N )：批量设置多个 string 的值
msetnx( key N, value N )：如果所有名称为 key 的 string 都不存在
incr( key )：名称为 key 的 string 增 1 操作
incrby( key, integer )：名称为 key 的 string 增加 integer
decr( key )：名称为 key 的 string 减 1 操作
decrby( key, integer )：名称为 key 的 string 减少 integer
append( key, value )：名称为 key 的 string 的值附加 value
substr( key, start, end )：返回名称为 key 的 string 的 value 的子串
*/
func GetValueFromKey(client *redis.Client, key string) (string, error) {
	return client.Get(key).Result()
}

func SetValueForKey(client *redis.Client, key string, val string) error {
	//如果需要设置k-v的过期时间(比如过期时间为10秒)，0可以改成 time.Second * 10
	return client.Set(key, val, time.Second*10).Err()
}

/*
rpush( key, value )：在名称为 key 的 list 尾添加一个值为 value 的元素
lpush( key, value )：在名称为 key 的 list 头添加一个值为 value 的元素
llen( key )：返回名称为 key 的 list 的长度
lrange( key, start, end )：返回名称为 key 的 list 中 start 至 end 之间的元素
ltrim( key, start, end )：截取名称为 key 的 list
lindex( key, index )：返回名称为 key 的 list 中 index 位置的元素
lset( key, index, value )：给名称为 key 的 list 中 index 位置的元素赋值
lrem( key, count, value )：删除 count 个 key 的 list 中值为 value 的元素
lpop( key )：返回并删除名称为 key 的 list 中的首元素
rpop( key )：返回并删除名称为 key 的 list 中的尾元素
blpop( key1, key2,… key N, timeout )：lpop 命令的 block 版本。
brpop( key1, key2,… key N, timeout )：rpop 的 block 版本。
rpoplpush( srckey, dstkey )：返回并删除名称为 srckey 的 list 的尾元素，并将该元素添加到名称为 dstkey 的 list 的头部
*/
//list操作
func ListOperator(client *redis.Client) {
	client.RPush("fruit", "apple")  //在名称为fruit的list尾部添加apple
	client.LPush("fruit", "banana") //在名称为fruit的list头部添加banana
	client.RPush("fruit", "orange")
	client.LPush("fruit", "peer")
	len, err := client.LLen("fruit").Result() //得到list的长度
	errorHandler(err)
	fmt.Println(len)
	val, err := client.LRange("fruit", 0, len-1).Result() //得到list中所有的元素
	errorHandler(err)
	fmt.Println(val)
	client.LSet("fruit", 0, "watermeter") //0为index，就是将list中第index个元素进行修改
	val, err = client.LRange("fruit", 0, len-1).Result()
	errorHandler(err)
	fmt.Println(val)
	val1, err := client.RPop("fruit").Result() //弹出list中首部(右边)的元素
	errorHandler(err)
	fmt.Println(val1)
	val2, err := client.LPop("fruit").Result() //弹出list中尾部(左边)的元素
	errorHandler(err)
	fmt.Println(val2)
}

/*
sadd( key, member )：向名称为 key 的 set 中添加元素 member
srem( key, member )：删除名称为 key 的 set 中的元素 member
spop( key )：随机返回并删除名称为 key 的 set 中一个元素
smove( srckey, dstkey, member )：移到集合元素
scard( key )：返回名称为 key 的 set 的基数
sismember( key, member )：member 是否是名称为 key 的 set 的元素
sinter( key1, key2,…key N )：求交集
sinterstore( dstkey, ( keys ))`：求交集并将交集保存到 dstkey 的集合
sunion( key1, ( keys ))`：求并集
sunionstore( dstkey, ( keys ))`：求并集并将并集保存到 dstkey 的集合
sdiff( key1, ( keys ))`：求差集
sdiffstore( dstkey, ( keys ))`：求差集并将差集保存到 dstkey 的集合
smembers( key )：返回名称为 key 的 set 的所有元素
srandmember( key )：随机返回名称为 key 的 set 的一个元素
*/
//set操作
func SetOperator(client *redis.Client) {
	client.SAdd("black", "Obama")   //向black set中添加Obama
	client.SAdd("black", "Hillary") //再添加一个
	client.SAdd("black", "Elder")   //再添加一个
	client.SAdd("write", "Hillary")
	client.SAdd("write", "Elder")
	client.SAdd("write", "Sam")
	//判断元素是否在集合中
	val, err := client.SIsMember("black", "Bush").Result()
	errorHandler(err)
	fmt.Println(val)
	//求交集，既在black中，又在write中
	names, err := client.SInter("black", "write").Result()
	errorHandler(err)
	fmt.Println(names)

	//获得执行集合的所有元素
	val1, err := client.SMembers("black").Result()
	errorHandler(err)
	fmt.Println(val1)

	//求并集
	val2, err := client.SUnion("black", "write").Result()
	errorHandler(err)
	fmt.Println(val2)
}

/*
hset( key, field, value )：向名称为 key 的 hash 中添加元素 field
hget( key, field )：返回名称为 key 的 hash 中 field 对应的 value
hmget( key, ( fields ) )：返回名称为 key 的 hash 中 field 对应的 value
hmset( key, ( fields ) )：向名称为 key 的 hash 中添加元素 field
hincrby( key, field, integer )：将名称为 key 的 hash 中 field 的 value 增加 integer
hexists( key, field )：名称为 key 的 hash 中是否存在键为 field 的域
hdel( key, field )：删除名称为 key 的 hash 中键为 field 的域
hlen( key )：返回名称为 key 的 hash 中元素个数
hkeys( key )：返回名称为 key 的 hash 中所有键
hvals( key )：返回名称为 key 的 hash 中所有键对应的 value
hgetall( key )：返回名称为 key 的 hash 中所有的键（ field ）及其对应的 value
*/
//hashset操作
func HashSetOperator(client *redis.Client) {
	client.HSet("user", "name", "liurui") //向uesr这个hashset中添加key为name，value为刘锐的元素
	client.HSet("user", "age", "39")
	len, err := client.HLen("user").Result()
	errorHandler(err)
	fmt.Println("length of hashset is:", len)
	name, err := client.HGet("user", "name").Result()
	errorHandler(err)
	fmt.Println("the name in hashset is:", name)
	client.HDel("user", "name") //删除user这个hashset的name字段
	client.HDel("user", "age")  //删除user这个hashset的age字段
	var mymap map[string]interface{} = make(map[string]interface{})
	mymap["name"] = "liurui1"
	mymap["age"] = "40"
	client.HMSet("user_map", mymap)
	val, err := client.HMGet("user_map", "name", "age").Result()
	errorHandler(err)
	fmt.Println(val)
	val1, err := client.HGetAll("user_map").Result()
	fmt.Println(val1)
}
