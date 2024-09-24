package pg

import "context"

type ctxQueryComment string

var (
	ctxQueryCommentKey = ctxQueryComment("dbQueryComment")
)

func SetQueryComment(ctx context.Context, query string) context.Context {
	return context.WithValue(ctx, ctxQueryCommentKey, query)
}

func GetQueryComment(ctx context.Context) string {
	queryComment, _ := ctx.Value(ctxQueryCommentKey).(string)
	return queryComment
}
