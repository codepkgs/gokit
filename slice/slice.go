package slice

import "fmt"

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Int | Uint
}

type Slice[T Integer | Float | string] struct {
	elements []T
}

// Append 追加元素
func (s *Slice[T]) Append(v T) {
	s.elements = append(s.elements, v)
}

// Len 获取长度
func (s *Slice[T]) Len() int {
	return len(s.elements)
}

// Count 统计元素个数
func (s *Slice[T]) Count(v T) int {
	t := 0
	for _, e := range s.elements {
		if e == v {
			t++
		}
	}
	return t
}

// Remove 移除第一个找到的元素
func (s *Slice[T]) Remove(v T) error {
	find := false
	for i, e := range s.elements {
		if e == v {
			s.elements = append(s.elements[:i], s.elements[i+1:]...)
			find = true
			break
		}
	}

	if find {
		return nil
	}
	return fmt.Errorf("%v not in slice", v)
}

// Pop 弹出指定位置的元素，并返回弹出的元素
func (s *Slice[T]) Pop(idx uint) (T, error) {
	var zero T
	if idx >= uint(len(s.elements)) {
		return zero, fmt.Errorf("idx out of range")
	}
	popEle := s.elements[idx]
	s.elements = append(s.elements[:idx], s.elements[idx+1:]...)
	return popEle, nil
}

// Index 查找元素在切片中的位置，返回索引
func (s *Slice[T]) Index(value T, start, stop uint) (int, error) {
	var es []T
	if start >= stop {
		return 0, fmt.Errorf("start must be less than stop")
	}

	if start >= uint(len(s.elements)) {
		return 0, fmt.Errorf("start out of range")
	}

	if stop >= uint(len(s.elements)) {
		es = s.elements[start:]
	} else {
		es = s.elements[start:stop]
	}

	for i, e := range es {
		if e == value {
			return int(start) + i, nil
		}
	}
	return 0, fmt.Errorf("%v not in slice", value)
}

// Clear 清空Slice
func (s *Slice[T]) Clear() {
	s.elements = nil
}

// Insert 在指定位置插入元素
// 如果idx大于切片长度，则插入到末尾
func (s *Slice[T]) Insert(idx uint, value T) {
	if idx >= uint(len(s.elements)) {
		s.elements = append(s.elements, value)
	} else {
		s.elements = append(s.elements[:idx], append([]T{value}, s.elements[idx:]...)...)
	}
}

// Extend 合并切片
func (s *Slice[T]) Extend(st Slice[T]) {
	s.elements = append(s.elements, st.elements...)
}

// Contain 判断元素是否在切片中
func (s *Slice[T]) Contain(v T) bool {
	for _, e := range s.elements {
		if e == v {
			return true
		}
	}
	return false
}

// Add 添加多个元素
func (s *Slice[T]) Add(vals ...T) {
	s.elements = append(s.elements, vals...)
}
