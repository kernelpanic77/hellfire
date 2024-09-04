// interface for a generic client
// could be extended based on the usecases
// TODO: httpClient and grpcClient

package hellfire

type Client interface {
	SendReq()
}
