package zsh

import (
	"strings"
	"time"
)

type cachedRun struct {
	Stdout *strings.Builder
	Stderr *strings.Builder
	Err    error
	T      time.Time
}

var cachedRuns map[string]cachedRun = make(map[string]cachedRun)

func cacheKey(path string, args []string) (key string) {
	args = append(args, path)
	key = strings.Join(args, "--")
	return
}

func fromCache(path string, args []string) (cached cachedRun, found bool) {
	key := cacheKey(path, args)
	cached, found = cachedRuns[key]
	return
}

func toCache(
	path string,
	args []string,
	stdout *strings.Builder, stderr *strings.Builder, err error,
) {
	key := cacheKey(path, args)
	cachedRuns[key] = cachedRun{Stdout: stdout, Stderr: stderr, Err: err, T: time.Now().UTC()}
}
