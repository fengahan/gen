//go
package main

import (
   "gen/cmd/gen_route"
   "github.com/creack/pty"
   "io"
   "os"
   "os/exec"
)

func main()  {
   fs,_:=os.Getwd()
   gen_route.Gen(".","app/api/gen_build/auto_gen_router.go")
   c := exec.Command("go","run",fs+"/app/api/api_server.go")
   f, err := pty.Start(c)
   if err != nil {
      panic(err)
   }
   io.Copy(os.Stdout, f)
}
