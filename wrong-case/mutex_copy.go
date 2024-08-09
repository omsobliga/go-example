/*
复制sync标准库包中的类型的值
在实践中，sync标准库包中的类型（除了Locker接口类型）的值不应该被复制。 我们只应该复制它们的指针值。

下面是一个有问题的并发编程的例子。 在此例子中，当Counter.Value方法被调用时，一个Counter属主值将被复制，此属主值的字段Mutex也将被一同复制。 此复制并没有被同步保护，因此复制结果可能是不完整的，并非被复制的属主值的一个快照。 即使此Mutex字段得以侥幸完整复制，它的副本所保护的是对字段n的一个副本的访问，因此一般是没有意义的。
*/
import "sync"

type Counter struct {
	sync.Mutex
	n int64
}

// 此方法实现是没问题的。
func (c *Counter) Increase(d int64) (r int64) {
	c.Lock()
	c.n += d
	r = c.n
	c.Unlock()
	return
}

// 此方法的实现是有问题的。当它被调用时，
// 一个Counter属主值将被复制。
func (c Counter) Value() (r int64) {
	c.Lock()
	r = c.n
	c.Unlock()
	return
}
