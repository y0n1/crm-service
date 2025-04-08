package collections

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any](size int) Map[K, V] {
	return make(Map[K, V], size)
}

func (m Map[K, V]) Get(key K) (V, bool) {
	value, exists := m[key]
	if !exists {
		var zero V
		return zero, false
	}
	return value, true
}

func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

func (m Map[K, V]) Delete(key K) {
	delete(m, key)
}

func (m Map[K, V]) Has(key K) bool {
	_, has := m[key]
	return has
}

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) Clear() {
	for key := range m {
		delete(m, key)
	}
}

func (m Map[K, V]) ForEach(fn func(key K, value V)) {
	for key, value := range m {
		fn(key, value)
	}
}
