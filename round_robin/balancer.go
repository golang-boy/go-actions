package roundrobin

import "google.golang.org/grpc/balancer"

// Pick returns the connection to use for this RPC and related information.
// Pick should not block.  If the balancer needs to do I/O or any blocking
// or time-consuming work to service this call, it should return
// ErrNoSubConnAvailable, and the Pick call will be repeated by gRPC when
// the Picker is updated (using ClientConn.UpdateState).
//
// If an error is returned:
//
//   - If the error is ErrNoSubConnAvailable, gRPC will block until a new
//     Picker is provided by the balancer (using ClientConn.UpdateState).
//
//   - If the error is a status error (implemented by the grpc/status
//     package), gRPC will terminate the RPC with the code and message
//     provided.
//
//   - For all other errors, wait for ready RPCs will wait, but non-wait for
//     ready RPCs will be terminated with this error's Error() string and
//     status code Unavailable.
func (b *Balancer) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	panic("not implemented") // TODO: Implement
}
