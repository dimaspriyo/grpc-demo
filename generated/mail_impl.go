package generated

import (
	"context"
	"fmt"
)

type MailImpl struct {
}

func (m MailImpl) mustEmbedUnimplementedRepositoryServiceServer() {
	panic("implement me")
}

func (m MailImpl) Receive(ctx context.Context, mail *Mail) (*Mail, error) {
	fmt.Println("Incoming Receive Method")
	return &Mail{
		Text: "Receive Method Executed",
	}, nil
}
