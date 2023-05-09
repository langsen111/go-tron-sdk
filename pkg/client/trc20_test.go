package client_test

import (
	"encoding/hex"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math/big"
	"testing"

	"github.com/langsen111/go-tron-sdk/pkg/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestTRC20_Balance(t *testing.T) {
	trc20Contract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // USDT
	address := "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9"

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())
	require.Nil(t, err)

	balance, err := conn.TRC20ContractBalance(address, trc20Contract)
	assert.Nil(t, err)
	assert.Greater(t, balance.Int64(), int64(0))
}

func TestGrpcClient_TRC20ContractAllowance(t *testing.T) {
	trc20Contract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // USDT
	address := "TDjfwCF8EcUph49MQj9u3fuNePMDELfAZr"
	spender := "TKcEU8ekq2ZoFzLSGFYCUY6aocJBX9X31b"

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())
	require.Nil(t, err)

	allowance, err := conn.TRC20ContractAllowance(address, spender, trc20Contract)
	println(allowance.String())
	assert.Nil(t, err)
	assert.Greater(t, allowance.Int64(), int64(0))
}
func Test_TransferTrc20(t *testing.T) {
	//c, err := grpcs.NewClient("54.168.218.95:50051")
	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(grpc.WithInsecure())
	require.Nil(t, err)

	amount := big.NewInt(1)
	amount = amount.Mul(amount, big.NewInt(1000000)) //TE9LmR3z5QoWCVmL9GQ2jdXonzeeFMw33G
	tx, err := conn.TRC20Send("TE9LmR3z5QoWCVmL9GQ2jdXonzeeFMw33G", "TCQDpWEfmugTwm7tqeMFEpvUt3cMJb5pvd",
		"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", amount, 5000000)
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	hexRawData := hex.EncodeToString(rawData)
	fmt.Println(hexRawData)

}
