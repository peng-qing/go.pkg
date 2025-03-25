package options

import "go.pkg/common/constraint"

// Options 接口
type Options[T constraint.Pointer[U], U any] interface {
	Apply(*T)
}

// OptFunc 函数类型
type OptFunc[T constraint.Pointer[U], U any] func(T)

// Apply 方法 实现 Options 接口
func (f OptFunc[T, U]) Apply(t T) {
	f(t)
}
