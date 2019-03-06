package test
import (
	"fmt"
	"testing"
)
/*
对于功能测试函数来说，其名称必须以Test为前缀，并且参数列表中应有一个*testing.T类型的参数声明
对于性能测试函数来说,其名称必须以Benchmark为前缀,并且唯一参数的类型必须是*testing.B类型的
对于示例测试函数来说,其名称必须以Example为前缀,但对函数的参数列表没有强制规定

go test -bench=. -run=^$ +main文件
-bench=. 只有有了这个标记才能进行性能测试
-run=^$ 用于表明需要执行哪些功能测试函数
*/
func BenchmarkHello(t *testing.B) {
	var name string
	greeting, err := hello(name)
	if err == nil {
		t.Errorf("The error is nil, but it should not be. (name=%q)",
			name)
	}
	if greeting != "" {
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)",
			name)
	}
	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("The error is not nil, but it should be. (name=%q)",
			name)
	}
	if greeting == "" {
		t.Errorf("Empty greeting, but it should not be. (name=%q)",
			name)
	}
	expected := fmt.Sprintf("Hello, %s!", name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not the expected. (name=%q)",
			greeting, name)
	}
	t.Logf("The expected greeting is %q.\n", expected)
}

func TestIntroduce(t *testing.T) {
	intro := it()
	expected := "Welcome to my Golang column."
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.",
			intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}