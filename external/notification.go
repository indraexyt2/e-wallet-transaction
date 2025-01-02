package external

import (
	"context"
	"e-wallet-ums/constants"
	"e-wallet-ums/external/proto/notification"
	"e-wallet-ums/helpers"
	"fmt"
	"google.golang.org/grpc"
)

func (e *External) SendNotification(ctx context.Context, recipient string, templateName string, placeholder map[string]string) error {
	conn, err := grpc.Dial(helpers.GetEnv("NOTIFICATION_GRPC_HOST", ""), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := notification.NewNotificationServiceClient(conn)
	request := &notification.SendNotificationRequest{
		TemplateName: templateName,
		Recipient:    recipient,
		Placeholders: placeholder,
	}

	resp, err := client.SendNotification(ctx, request)
	if err != nil {
		return err
	}

	if resp.Message != constants.SuccessMessage {
		return fmt.Errorf("failed to send notification: %s", resp.Message)
	}

	return nil
}
