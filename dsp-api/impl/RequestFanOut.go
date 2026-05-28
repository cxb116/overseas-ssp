package impl

import (
	"context"
	"fmt"
	"time"

	ants "github.com/panjf2000/ants/v2"
)

// ==================== Future ====================

type Future[T any] struct {
	ch    chan struct{}
	value T
	err   error
}

func newFuture[T any]() *Future[T] {
	return &Future[T]{ch: make(chan struct{})}
}

func (f *Future[T]) Value() (T, error) {
	<-f.ch
	return f.value, f.err
}

// ==================== Pool ====================

type Pool[T any] struct {
	inner *ants.Pool
}

func NewPool[T any](cap int) *Pool[T] {
	pool, _ := ants.NewPool(cap)
	return &Pool[T]{inner: pool}
}

// Submit 提交任务到池中，支持超时和 panic 恢复
func (p *Pool[T]) Submit(ctx context.Context, timeout time.Duration, fn func() (T, error)) *Future[T] {
	future := newFuture[T]()

	p.inner.Submit(func() {
		defer func() {
			if r := recover(); r != nil {
				var zero T
				future.value, future.err = zero, fmt.Errorf("panic: %v", r)
				close(future.ch)
			}
		}()

		done := make(chan struct{})
		go func() {
			defer close(done)
			future.value, future.err = fn()
		}()

		select {
		case <-done:
		case <-ctx.Done():
			var zero T
			future.value, future.err = zero, ctx.Err()
		case <-time.After(timeout):
			var zero T
			future.value, future.err = zero, fmt.Errorf("timeout after %v", timeout)
		}

		if future.err == nil {
			close(future.ch)
		}
	})

	return future
}

func (p *Pool[T]) Cap() int     { return p.inner.Cap() }
func (p *Pool[T]) Running() int { return p.inner.Running() }
func (p *Pool[T]) Release()     { p.inner.Release() }

//// ==================== Demo ====================
//
//func RunFanoutDemo() {
//	pool := NewPool[string](5)
//	defer pool.Release()
//
//	services := []struct {
//		Name string
//		Fn   func() (string, error)
//	}{
//		{"库存服务", func() (string, error) {
//			time.Sleep(150 * time.Millisecond)
//			return "库存扣减成功", nil
//		}},
//		{"积分服务", func() (string, error) {
//			time.Sleep(100 * time.Millisecond)
//			return "积分增加", nil
//		}},
//		{"优惠券服务", func() (string, error) {
//			time.Sleep(80 * time.Millisecond)
//			return "优惠券核销成功", nil
//		}},
//		{"风控服务", func() (string, error) {
//			time.Sleep(200 * time.Millisecond)
//			return "风控检查通过", nil
//		}},
//	}
//
//	ctx := context.Background()
//	timeout := 500 * time.Millisecond
//	futures := make([]*Future[string], len(services))
//
//	fmt.Println("开始并行调用下游服务...")
//	for i, svc := range services {
//		futures[i] = pool.Submit(ctx, timeout, svc.Fn)
//	}
//
//	for i, future := range futures {
//		val, err := future.Value()
//		if err != nil {
//			fmt.Printf("  ❌ [%s] %v\n", services[i].Name, err)
//		} else {
//			fmt.Printf("  ✅ [%s] %s\n", services[i].Name, val)
//		}
//	}
//}
