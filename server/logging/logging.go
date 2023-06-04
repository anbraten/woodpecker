package logging

import (
	"context"
	"errors"
	"io"

	"github.com/woodpecker-ci/woodpecker/server/model"
)

// ErrNotFound is returned when the log does not exist.
var ErrNotFound = errors.New("stream: not found")

// Handler defines a callback function for handling log entries.
type Handler func(...*model.LogEntry)

// Log defines a log multiplexer.
type Log interface {
	// Open opens the log.
	Open(c context.Context, path string) error

	// Write writes the entry to the log.
	Write(c context.Context, path string, entry *model.LogEntry) error

	// Tail tails the log.
	Tail(c context.Context, path string, handler Handler) error

	// Close closes the log.
	Close(c context.Context, path string) error

	// Snapshot snapshots the stream to Writer w.
	Snapshot(c context.Context, path string, w io.Writer) error

	// Info returns runtime information about the multiplexer.
	// Info(c context.Context) (interface{}, error)
}

// // global streamer
// var global = New()
//
// // Set sets a default global logger.
// func Set(log Log) {
// 	global = log
// }
//
// // Open opens the log stream.
// func Open(c context.Context, path string) error {
// 	return global.Open(c, path)
// }
//
// // Write writes the log entry to the stream.
// func Write(c context.Context, path string, entry *Entry) error {
// 	return global.Write(c, path, entry)
// }
//
// // Tail tails the log stream.
// func Tail(c context.Context, path string, handler Handler) error {
// 	return global.Tail(c, path, handler)
// }
//
// // Close closes the log stream.
// func Close(c context.Context, path string) error {
// 	return global.Close(c, path)
// }
//
// // Snapshot snapshots the stream to Writer w.
// func Snapshot(c context.Context, path string, w io.Writer) error {
// 	return global.Snapshot(c, path, w)
// }
