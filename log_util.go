package log_watcher

import (
	"context"
	"go.uber.org/zap"
)

const appLogging = "app_logging"

// setLogData
// This function is use for set log data from context
func setLogData(ctx context.Context, msg string) (logRecord []zap.Field) {
	sess := getContext(ctx)

	logRecord = append(logRecord, zap.String("app_request_id", sess.ThreadID))
	logRecord = append(logRecord, zap.String("app_journey_id", sess.JourneyID))
	logRecord = append(logRecord, zap.String("message", msg))
	logRecord = append(logRecord, zap.String("app_name", sess.ServiceName))
	logRecord = append(logRecord, zap.String("app_tag", sess.Tag))
	logRecord = append(logRecord, zap.String("app_version", sess.ServiceVersion))
	logRecord = append(logRecord, zap.Int("app_port", sess.ServicePort))
	logRecord = append(logRecord, zap.String("app_req_ip", sess.SrcIP))
	logRecord = append(logRecord, zap.String("app_method", sess.ReqMethod))
	logRecord = append(logRecord, zap.String("app_uri", sess.ReqURI))
	logRecord = append(logRecord, zap.Any("app_req_header", sess.Header))

	if sess.Request != nil {
		logRecord = append(logRecord, zap.Any("app_req_body", sess.Request))
	} else {
		logRecord = append(logRecord, zap.Any("app_req_body", nil))
	}

	if sess.AdditionalData != nil {
		logRecord = append(logRecord, zap.Any("app_additional_data", sess.AdditionalData))
	} else {
		logRecord = append(logRecord, zap.Any("app_additional_data", nil))
	}

	return
}

// SetContext
// This function use for set logging struct to golang context
func SetContext(ctx context.Context, contextId string, value interface{}) context.Context {
	key := appLogging

	//if contextId != "" {
	//	key = contextId
	//}

	ctxi := context.WithValue(ctx, key, value)

	return ctxi
}

// getContext
// This function use for this package to get logging struct from golang context
func getContext(ctx context.Context) Context {
	if ctx == nil {
		return Context{}
	}

	val, err := ctx.Value(appLogging).(Context)

	if !err {
		return Context{}
	}

	return val
}
