package response

import (
	"context"
	"fmt"
	"testing"

	"github.com/tyr-tech-team/hawk/status"
)

func Test_Error(t *testing.T) {
	x := status.NotFound.Err()

	ctx := context.TODO()

	fmt.Println(Error(ctx, x))

}

func Test_Resp(t *testing.T) {
	fmt.Println(Resp(context.TODO(), "123"))
}
