package main

import (
	"jiacrontab/server/rpc"
	"jiacrontab/server/store"
	"jiacrontab/server/config"
	_ "net/http/pprof"
	"runtime"
	"time"
)

var globalConfig *config.Config
var globalStore *store.Store
var globalJwt *mjwt
var globalReqFilter *reqFilter
var startTime = time.Now()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	globalConfig = config.NewConfig()
	globalJwt = newJwt(globalConfig.TokenExpires, globalConfig.TokenCookieName, globalConfig.JWTSigningKey, globalConfig.TokenCookieMaxAge)

	globalStore = store.NewStore(globalConfig.DataFile)
	globalStore.Load()

	globalReqFilter = newReqFilter()

	go rpc.InitSrvRpc(globalConfig.DefaultRPCPath, globalConfig.DefaultRPCDebugPath, globalConfig.RpcAddr, &Logic{})

	initServer()
}
