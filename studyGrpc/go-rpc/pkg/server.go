package pkg

import "net/rpc"

var (
	ServerName = "RpcServer"
	Addr       = ":1234"
)

type RpcServer interface {
	Server(request string, rsp *string) error
}

func Register(svr interface{}) error {
	if err := rpc.RegisterName(ServerName, svr); err != nil {
		return err
	}

	return nil
}
