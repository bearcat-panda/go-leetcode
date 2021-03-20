package 面试题

import (
	"sync"
	"sync/atomic"
)

type example struct {
	name string
}

var mux sync.Mutex
var instance *example

/**
如果这样去做，每一次请求单例的时候，都会加锁和减锁，
而锁的用处只在于解决对象初始化的时候可能出现的并发问题
当对象被创建之后，加锁就失去了意义，会拖慢速度，
所以我们就引入了双重检查机制（Check-lock-Check）, 也叫DCL(Double Check Lock), 代码如下:
*/
func GetInstance() *example {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = &example{}
	}
	return instance
}

func GetInstance1() *example {
	if instance == nil {
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = &example{}
		}
	}

	return instance
}

//这样只有当对象未初始化的时候，才会又加锁和减锁的操作
//但是又出现了另一个问题：每一次访问都要检查两次，为了解决这个问题，我们可以使用golang标准包中的方法进行原子性操作:
var initialized uint32

func GetInstance2() *example {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mux.Lock()
	defer mux.Unlock()

	if initialized == 0 {
		instance = &example{}
		atomic.StoreUint32(&initialized, 1) //原子装载
	}

	return instance
}

var once sync.Once

func GetInstanc3() *example {
	once.Do(func() {
		instance = new(example)
	})
	return instance
}

//地址:https://www.cnblogs.com/lurenq/p/11988920.html
