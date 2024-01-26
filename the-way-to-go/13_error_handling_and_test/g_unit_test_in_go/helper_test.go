package main

import (
	"cource-practice-go/the-way-to-go/13_error_handling_and_test/g_unit_test_in_go/helper"
	"fmt"
	"testing"
)

// 0. 测试文件需要符合 *_test.go 的模式

// 1. 所有单元测试的命名都要符合TestXxx的形式
func TestAdd(t *testing.T) {
	c := helper.Add(1, 2)
	if c != 3 {
		panic("invalid result")
	}
}

func TestWrongAdd_0(t *testing.T) {
	c := helper.WrongAdd(1, 2)
	if c != 3 {
		// 2. 标记这个test为失败，然后继续执行后面的test
		t.Fail()
	}

	// 所以这里能够看到有打印输出
	fmt.Println("[TestWrongAdd_0]: after test")
}

func TestWrongAdd_1(t *testing.T) {
	c := helper.WrongAdd(1, 2)
	if c != 3 {
		// 4. 这样可以在控制台看到失败的原因
		t.Log("invalid add result")
		// 3. 标记这个test为失败，并立刻终止执行
		t.FailNow()
	}

	// 所以这里是看不到有打印输出的
	fmt.Printf("after test")
}

func TestWrongAdd_2(t *testing.T) {
	c := helper.WrongAdd(1, 2)
	if c != 3 {
		// 5. 相当于先执行t.Log(...) 然后再执行 t.FailNow()
		t.Fatal("invalid add result")
	}

	// 所以这里是看不到有打印输出的
	fmt.Printf("after test")
}

// 6. Benchmark打头的函数，会被识别为基准测试函数，控制台会给出执行n次这个函数，一次需要多少时间
// func BenchmarDemo(b *testing.B) {
func BenchmarkDemo(b *testing.B) {
	forBenchmarkTesting()
}

func forBenchmarkTesting() {
	for i := 0; i < 100000000; i++ {

	}
}
