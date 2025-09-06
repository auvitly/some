package variable

import "maps"

var Lib = func() any { return entity }

type lib struct {
	storage map[any]any
}

var entity = &lib{
	storage: make(map[any]any),
}

func (l *lib) Set(key any, value any) any {
	l.storage[key] = value
	return value
}

func (l *lib) Get(key any) any {
	return l.storage[key]
}

func (l *lib) Del(key any) any {
	value := l.storage[key]

	delete(l.storage, key)

	return value
}

func (l *lib) All() map[any]any {
	return maps.Clone(l.storage)
}
