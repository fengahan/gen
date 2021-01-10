//go
package main

import (
   "gen/internal/gen_config"
   "gen/internal/gen_route"
   "github.com/creack/pty"
   "io"
   "os"
   "os/exec"
)

func main()  {

   gen_config.GenConfig()
   gen_route.Gen(".","app/api/gen_build/router_gen_target.go")
   fs,_:=os.Getwd()
   c := exec.Command("/Users/desmond/go/bin/wire",fs+"/app/api/gen_build/wire_controller.go")
   f, err := pty.Start(c)
   if err != nil {
    panic(err)
   }
   io.Copy(os.Stdout, f)

}
