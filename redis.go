package redis

import (
	"time"

	"github.com/go-redis/redis"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/redis", new(REDIS))
}

// REDIS is the k6 Redis extension.
type REDIS struct{}

// NewClient creates a new Redis client
func (*REDIS) NewClient(addr string, password string, bd int) *redis.Client {
	if addr == "" {
		addr = "localhost:6379"
	}
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       bd,       // use default DB
	})
}

func (*REDIS) NewClusterClient(addrs []string) *redis.ClusterClient {
	/*if addr == "" {
		addr = "localhost:6379"
	}*/
	return redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: addrs,
      //  Password: password,
    })
}

// Set adds a key/value
func (*REDIS) Set(client *redis.Client, key string, value interface{}, expiration time.Duration) {
	// TODO: Make expiration configurable. Or document somewhere the unit.
	err := client.Set(key, value, expiration*time.Second).Err()
	if err != nil {
		ReportError(err, "Failed to set the specified key/value pair")
	}
}


// Get gets a key/value
func (*REDIS) Get(client *redis.Client, key string) string {
	val, err := client.Get(key).Result()
	if err != nil {
		ReportError(err, "Failed to get the specified key")
	}
	return val
}

// Del removes a key/value
func (*REDIS) Del(client *redis.Client, key string) {
	err := client.Del(key).Err()
	if err != nil {
		ReportError(err, "Failed to remove the specified key")
	}
}

// Do runs arbitrary/custom commands
func (*REDIS) Do(client *redis.Client, cmd string, key string) string {
	val, err := client.Do(cmd, key).Result()
	if err != nil {
		if err == redis.Nil {
			ReportError(err, "Key does not exist")
		} else {
			ReportError(err, "Failed to do command")
		}
	}
	// TODO: Support more types, not only strings.
	return val.(string)
}
// Test Function, returns Hello + string
func (*REDIS) Sayhi(value string) string {
	val:= "Hello "+ value
	return val
}


// Set2 is like Set but it works inserting multiple fields with values.
func (*REDIS) Set2(client *redis.Client, key string,a []string,b []string) {
	// TODO: Make expiration configurable. Or document somewhere the unit.
	fields := make(map[string]interface{})
	for i := 0; i < len(a); i++ {
		fields[a[i]]=b[i]
	}
	//err := 
	client.HMSet(key, fields)
        //.Err()
	//if err != nil {
	//	ReportError(err, "Failed to set the specified key/value pair")
	//}
}


// Expire add the expiration time to the specified key
func (*REDIS) Set3(client *redis.Client, key string,values interface{}) {
	// TODO: Make expiration configurable. Or document somewhere the unit.
	err := client.LPush(key, values).Err()
	if err != nil {
		ReportError(err, "Failed to lpush the value to specified key")
	}
}

// Expire add the expiration time to the specified key
func (*REDIS) Expire(client *redis.Client, key string,expiration time.Duration) {
	// TODO: Make expiration configurable. Or document somewhere the unit.
	
	err := client.Expire(key, expiration*time.Second).Err()
	if err != nil {
		ReportError(err, "Failed to set the expiration to specified key")
	}
}
