package httpx

import (
	"context"
	"encoding/json"
	"go-zero-study/core/ecode"
	"go-zero-study/core/trace/tracespec"
	"net/http"

	"go-zero-study/core/logx"
)

type JSONMsg struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func traceIdFromContext(ctx context.Context) string {
	t, ok := ctx.Value(tracespec.TracingKey).(tracespec.Trace)
	if !ok {
		return ""
	}

	return t.TraceId()
}

func JSON(ctx context.Context, w http.ResponseWriter, data interface{}, err error) {
	code := http.StatusOK
	if err != nil {
		logx.WithContext(ctx).Error(err)
	}
	bcode := ecode.Cause(err)
	if bcode.Code() <= http.StatusNetworkAuthenticationRequired && bcode.Code() >= http.StatusOK {
		code = bcode.Code()
	}

	jsonMsg := JSONMsg{
		Code:    bcode.Code(),
		Message: bcode.Message(),
		Data:    data,
	}

	w.Header().Set(ContentType, ApplicationJson)
	w.Header().Set("trace-id", traceIdFromContext(ctx))
	w.WriteHeader(code)

	if bs, err := json.Marshal(jsonMsg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	} else if n < len(bs) {
		logx.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}
}
