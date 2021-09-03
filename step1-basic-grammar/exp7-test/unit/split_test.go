package unit

import (
	"reflect"
	"testing"
)

// 基础单元测试
func TestSplit1(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("aabaa", "b")         // 程序输出的结果
	want := []string{"aa", "aa"}       // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

// 测试组和子测试
func TestSplit2(t *testing.T) {
	// 定义一个test结构体
	type test struct {
		input  string
		sep    string
		output []string
	}

	// 通过map存储测试用例
	tests := map[string]test{
		"test1": {input: "你好世界", sep: "好", output: []string{"你", "世界"}},
		"test2": {input: "hello,world", sep: ",", output: []string{"hello", "world"}},
		"test3": {input: "aabaa", sep: "b", output: []string{"aa", "aa"}},
		"test4": {input: "aabbaa", sep: "b", output: []string{"aa", "", "aa"}},
	}

	for name, test := range tests {
		// t.Run(): 执行子测试，可以使go test -v输出结果更为清晰
		t.Run(name, func(t *testing.T) {
			ret := Split(test.input, test.sep)
			if !reflect.DeepEqual(test.output, ret) {
				t.Errorf("test %s failed, excepted: %#v, but got: %#v", name, test.output, ret)
			}
		})
	}
}

// 性能基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("hello,world!", ",")
	}
}
