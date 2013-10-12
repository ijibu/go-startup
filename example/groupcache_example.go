/*
参考自官方提供的实例，有些许改动
官方实例：http://talks.golang.org/2013/oscon-dl/groupcache.go
*/

package main

import (
	"fmt"
	groupcache "github.com/golang/groupcache"
)

func main() {
	// STARTINIT OMIT
	me := "http://127.0.0.1:11211"
	peers := groupcache.NewHTTPPool(me)

	// Whenever peers change:
	peers.Set("http://127.0.0.1:11211")
	// ENDINIT OMIT

	// STARTGROUP OMIT
	var thumbNails = groupcache.NewGroup("thumbnail", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			dest.SetBytes(generateThumbnail(key))
			return nil
		}))
	// ENDGROUP OMIT

	var ctx groupcache.Context

	// STARTUSE OMIT
	var data []byte
	thumbNails.Get(ctx, "big-file.jpg", groupcache.AllocatingByteSliceSink(&data))

	fmt.Println(string(data))
}

func generateThumbnail(filename string) []byte {
	// ...
	return []byte(filename)
	//return nil
}
