package grpc

import "google.golang.org/grpc/resolver"

type Builder struct {
}

func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	r := &Resolver{cc: cc}

	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

func (b *Builder) Scheme() string {
	return "http"
}

type Resolver struct {
	cc resolver.ClientConn
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {
	err := r.cc.UpdateState(resolver.State{
		Addresses: []resolver.Address{
			{
				Addr: "http://localhost:8080",
			},
		},
	})

	if err != nil {
		r.cc.ReportError(err)
	}
}

func (r *Resolver) Close() {

}
