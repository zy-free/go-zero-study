package handler

import (
	"fmt"
	"go-zero-study/core/logx"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
)

func RecoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var rawReq []byte
		defer func() {
			if result := recover(); result != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				if r != nil {
					rawReq, _ = httputil.DumpRequest(r, false)
				}
				pl := fmt.Sprintf("http call panic: %s\n%s\n\n", string(rawReq), buf)
				fmt.Fprintf(os.Stderr, pl)
				logx.WithContext(r.Context()).Error(pl)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
