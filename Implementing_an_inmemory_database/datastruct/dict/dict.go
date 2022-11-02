package dict

type Consumer func(key string, val interface{}) bool

type Dict interface {
	// Get 空接口，是否存在
	Get(key string) (val interface{}, exists bool)
	Len() int
	// Put 存放了几个
	Put(key string, val interface{}) (result int)
	PutIfAbset(key string, val interface{}) (result int)
	PutIfExists(key string, val interface{}) (result int)
	Remove(key string) (result int)
	ForEach(consumer Consumer)
	Keys() []string
	RandomKeys(limit int) []string
	RandomDistinctKeys(limit int) []string
	Clear()
}
