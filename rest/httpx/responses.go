package httpx

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-study/core/ecode"
	"go-zero-study/core/trace"
	"net/http"

	"go-zero-study/core/logx"
)

type JSONMsg struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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
	w.Header().Set("trace-id", trace.TraceIdFromContext(ctx))
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

type CSVMsg struct {
	Content []byte
	Title   string
}

func CSV(ctx context.Context, w http.ResponseWriter, csv CSVMsg, err error) {
	code := http.StatusOK
	if err != nil {
		logx.WithContext(ctx).Error(err)
	}
	bcode := ecode.Cause(err)
	if bcode.Code() <= http.StatusNetworkAuthenticationRequired && bcode.Code() >= http.StatusOK {
		code = bcode.Code()
	}

	w.Header().Set(ContentType, ApplicationCsv)
	w.Header().Set("trace-id", trace.TraceIdFromContext(ctx))
	w.WriteHeader(code)

	w.Header()["Content-Disposition"] = append(w.Header()["Content-Disposition"], fmt.Sprintf("attachment; filename=%s.csv", csv.Title))

	if _, err := w.Write(csv.Content); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if err != http.ErrHandlerTimeout {
			logx.Errorf("write response failed, error: %s", err)
		}
	}
}
