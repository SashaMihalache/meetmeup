package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/sashamihalache/meetmeup/models"
)

const userloaderKey = "userloader"

func DataloaderMiddlerware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloder := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					return nil, []error{err}
				}

				return users, nil
			},
		}

		ctx := context.WithValue(r.Context(), userloaderKey, &userloder)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func getUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userloaderKey).(*UserLoader)
}
