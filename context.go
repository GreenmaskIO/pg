package pg

import "context"

type ctxQueryComment string

var (
	ctxQueryCommentKey = ctxQueryComment("dbQueryComment")
)

// SetQueryComment sets the query comment in the context for the query.
// This string will be added to the query statement as a comment in first line
func SetQueryComment(ctx context.Context, query string) context.Context {
	return context.WithValue(ctx, ctxQueryCommentKey, query)
}

// GetQueryComment returns the query comment from the context
func GetQueryComment(ctx context.Context) string {
	queryComment, _ := ctx.Value(ctxQueryCommentKey).(string)
	return queryComment
}
