package functions

import (
	"github.com/rohan-sharma92/accountapi-client-lib/form3/core"
)

type Form3 struct {
	Client core.Client
}

func Init(host string) *Form3 {
	client := core.CreateClient(host)
	return &Form3{
		Client: *client,
	}
}

func InitWithClient(c core.Client) *Form3 {
	return &Form3{
		Client: c,
	}
}
