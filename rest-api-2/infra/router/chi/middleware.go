package routerChi

import (
	"context"
	"log"
	"net/http"
	"rest-api-2/utility"
	"strings"
	"time"
)

func Tracer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before Request")

		h.ServeHTTP(w, r)

		log.Println("After Request")
	})
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		log.Printf("method=%v path=%v type=[INFO] message='incoming request'", r.Method, r.URL.Path)

		h.ServeHTTP(w, r)

		end := time.Since(now).Milliseconds()

		log.Printf("method=%v path=%v type=[INFO] message='finish request' response_time=%vms", r.Method, r.URL.Path, end)
	})
}

func CheckToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")

		if bearer == "" {
			resp := APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Error:   "no token provided",
			}
			WriteJsonResponse(w, resp)
			return
		}

		tokSlice := strings.Split(bearer, "Bearer ")

		if len(tokSlice) < 2 {
			resp := APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Error:   "invalid token",
			}
			WriteJsonResponse(w, resp)
			return
		}

		tokString := tokSlice[1]

		token, err := utility.VerifyToken(tokString)

		if err != nil {
			resp := APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Error:   err.Error(),
			}
			WriteJsonResponse(w, resp)
			return
		}

		ctx := context.WithValue(r.Context(), "AUTH_ID", token.Id)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
