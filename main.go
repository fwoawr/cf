package main

import (
	"github.com/fwoawr/cf/cmd"
	_ "github.com/fwoawr/cf/cmd/alibaba"
	_ "github.com/fwoawr/cf/cmd/aws"
	_ "github.com/fwoawr/cf/cmd/huawei"
	_ "github.com/fwoawr/cf/cmd/tencent"
)

func main() {
	cmd.Execute()
}
