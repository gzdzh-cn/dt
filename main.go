package main

import (
	"dzhgo-cli/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	cmd.Root.Run(ctx)
}
