package main

import (
   "gen/cmd/gen_route"
   "os"
   "os/exec"
)

func main()  {
   f,_:=os.Getwd()
   gen_route.Gen(".","app/api/gen_build/auto_gen_router.go")
   exec.Command("go","run",f+"/app/api/api_server.go").Run().Error()
}
