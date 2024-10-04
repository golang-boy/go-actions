AOP设计
---


## select


  get接口内实现中间件的调用，中间类似web中的函数式中间件调用，通过函数式编程，将中间件和业务逻辑分离，实现中间件的可插拔，可扩展。

  ```go

type Middleware func(next HandleFunc) HandleFunc

type HandleFunc func(ctx context.Context, qc *QueryContext) *QueryResult

type core struct {
	r          model.Registry
	dialect    Dialect
	valCreator valuer.Creator
	ms         []Middleware
	model      *model.Model
}

func get[T any](ctx context.Context, c core, sess Session, qc *QueryContext) *QueryResult {
	var handler HandleFunc = func(ctx context.Context, qc *QueryContext) *QueryResult {
		return getHandler[T](ctx, sess, c, qc)
	}
	ms := c.ms
	for i := len(ms) - 1; i >=0; i-- {
		handler = ms[i](handler)
	}
	return handler(ctx, qc)
}
  ```

## insert

  在exec中类似的方式实现中间件