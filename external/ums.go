package external

import (
	"context"
	"e-wallet-transaction/external/proto/tokenvalidation"
	"e-wallet-transaction/helpers"
	"e-wallet-transaction/internal/models"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type External struct {
}

func (e *External) ValidateToken(ctx context.Context, token string) (*models.TokenData, error) {
	var (
		resp models.TokenData
	)

	conn, err := grpc.Dial(helpers.GetEnv("UMS_GRPC_HOST", ""), grpc.WithInsecure())
	if err != nil {
		return &resp, errors.Wrap(err, "failed to connect to grpc server")
	}
	defer conn.Close()

	client := tokenvalidation.NewTokenValidationClient(conn)

	req := &tokenvalidation.TokenRequest{
		Token: token,
	}
	response, err := client.ValidateToken(ctx, req)
	if err != nil {
		return &resp, errors.Wrap(err, "failed to validate token")
	}

	if response.Message != "success" {
		return &resp, errors.New("failed to validate token")
	}

	fmt.Println(resp.Username)

	resp.UserID = response.Data.UserId
	resp.Username = response.Data.Username
	resp.FullName = response.Data.FullName
	resp.Email = response.Data.Email

	return &resp, nil
}
