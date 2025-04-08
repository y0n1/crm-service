package collections

type List[T any] []T

func NewList[T any](size int) List[T] {
	return make(List[T], size)
}

func (l *List[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(*l) {
		var zero T
		return zero, false // Index out of bounds
	}
	return (*l)[index], true
}

func (l *List[T]) Set(index int, item T) bool {
	if index < 0 || index >= len(*l) {
		return false // Index out of bounds
	}
	(*l)[index] = item
	return true
}

func (l *List[T]) Length() int {
	return len(*l)
}

func (l *List[T]) Clear() {
	*l = make(List[T], 0)
}

func (l *List[T]) ForEach(fn func(index int, item T)) {
	for i, item := range *l {
		fn(i, item)
	}
}

func (l *List[T]) Filter(fn func(item T) bool) List[T] {
	var result List[T]
	for _, item := range *l {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func (l *List[T]) Map(fn func(item T) T) List[T] {
	var result List[T]
	for _, item := range *l {
		result = append(result, fn(item))
	}
	return result
}

func (l *List[T]) Reduce(fn func(acc T, item T) T, initial T) T {
	acc := initial
	for _, item := range *l {
		acc = fn(acc, item)
	}
	return acc
}

func (l *List[T]) IsEmpty() bool {
	return len(*l) == 0
}

func (l *List[T]) ToSlice() []T {
	return *l
}

func (l *List[T]) FromSlice(slice []T) {
	*l = make(List[T], len(slice))
	copy(*l, slice)
}
