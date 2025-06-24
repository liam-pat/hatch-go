package main

import (
	"container/ring"
	"fmt"
)

/**
  func New(n int) *Ring  // 初始化环
  func (r *Ring) Do(f func(interface{}))  // 循环环进行操作
  func (r *Ring) Len() int // 环长度
  func (r *Ring) Link(s *Ring) *Ring // 连接两个环
  func (r *Ring) Move(n int) *Ring // 指针从当前元素开始向后移动或者向前（n 可以为负数）
  func (r *Ring) Next() *Ring // 当前元素的下个元素
  func (r *Ring) Prev() *Ring // 当前元素的上个元素
  func (r *Ring) Unlink(n int) *Ring // 从当前元素开始，删除 n 个元素
*/

func main() {
	pointNums := 3
	rings := ring.New(pointNums)

	for i := 0; i < pointNums; i++ {
		rings.Value = i + 10
		rings = rings.Next()
	}

	sum := 0
	rings.Do(func(p interface{}) {
		sum += p.(int)
	})

	fmt.Println("ring len", rings.Len())
	fmt.Println("ring point value sum is", sum)
}
