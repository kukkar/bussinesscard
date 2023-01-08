package bussinesscard

import "context"

type bussinesscard interface {
	FetchProfile(ctx context.Context, userID int)
}
