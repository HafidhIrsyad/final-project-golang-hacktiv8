package middleware

import (
	"context"
	"encoding/json"
	"final-project/entity"
	"final-project/helper"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"
)

var ctxKey = &contextKey{"user"}

type contextKey struct {
	data string
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.Contains(header, "Bearer") {
			response, _ := json.Marshal(helper.APIResponseFailed("Unauthorized", http.StatusUnauthorized, false))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)

			return
		}

		tokenString := ""
		arrayToken := strings.Split(header, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//validate jwt token
		token, err := ValidateToken(tokenString)
		if err != nil {
			response, _ := json.Marshal(helper.APIResponseFailed("Invalid token", http.StatusBadRequest, false))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)

			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			response, _ := json.Marshal(helper.APIResponseFailed("Unauthorized", http.StatusUnauthorized, false))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(response)

			return
		}

		userID := payload["id"].(string)

		id, _ := strconv.Atoi(userID)
		user := entity.User{ID: id}

		ctx := context.WithValue(r.Context(), ctxKey, &user)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) *entity.User {
	raw, _ := ctx.Value(ctxKey).(*entity.User)
	return raw
}
