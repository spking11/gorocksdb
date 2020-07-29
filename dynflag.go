// +build !linux !static

package gorocksdb

// #cgo LDFLAGS: -lrocksdb -lstdc++ -lm -ldl -lz -lbz2 -lsnappy -llz4 -lzstd
import "C"
