package emap

// mapKey 定义一个键的类型
type mapKey interface {
	~int | ~uint | ~string
}

// Map 定义一个字典
type Map[K mapKey, V any] map[K]V

// MapItem 键值对
type MapItem[K mapKey, V any] struct {
	Key   K
	Value V
}

// NewMap 创建一个字典
func NewMap[Key mapKey, Value any]() Map[Key, Value] {
	return make(map[Key]Value)
}

// HasKey 判断字典中是否存在某个键
func (m Map[K, V]) HasKey(k K) bool {
	_, ok := m[k]
	return ok
}

// Get 返回字典的值
func (m Map[K, V]) Get(k K) V {
	return m[k]
}

// Keys 返回字典的键
func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values 返回字典的值
func (m Map[K, V]) Values() []V {
	vals := make([]V, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

// Len 返回字典的长度
func (m Map[K, V]) Len() int {
	return len(m)
}

// Items 返回字典的键值对
func (m Map[K, V]) Items() []MapItem[K, V] {
	items := make([]MapItem[K, V], 0, len(m))
	for k, v := range m {
		items = append(items, MapItem[K, V]{Key: k, Value: v})
	}
	return items
}

// Update 更新字典
func (m Map[K, V]) Update(mm Map[K, V]) {
	for k, v := range mm {
		m[k] = v
	}
}
