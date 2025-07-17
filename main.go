package main

import (
	"github.com/gzdzh-cn/dt/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	cmd.Root.Run(ctx)
}
