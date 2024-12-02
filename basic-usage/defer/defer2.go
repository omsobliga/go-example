// 延迟调用使得代码更简洁和鲁棒
import "os"

func withoutDefers(filepath string, head, body []byte) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	_, err = f.Seek(16, 0)
	if err != nil {
		f.Close()
		return err
	}

	_, err = f.Write(head)
	if err != nil {
		f.Close()
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		f.Close()
		return err
	}

	err = f.Sync()
	f.Close()
	return err
}

func withDefers(filepath string, head, body []byte) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Seek(16, 0)
	if err != nil {
		return err
	}

	_, err = f.Write(head)
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	return f.Sync()
}


// 下面是另外一个延迟调用使得代码更鲁棒的例子。 如果doSomething函数产生一个恐慌，则函数f2在退出时将导致互斥锁未解锁。 所以函数f1更鲁棒。
var m sync.Mutex

func f1() {
	m.Lock()
	defer m.Unlock()
	doSomething()
}

func f2() {
	m.Lock()
	doSomething()
	m.Unlock()
}
