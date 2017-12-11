package view

import (
	"fmt"
	"html/template"
	"jiacrontab/server/rpc"
	"jiacrontab/server/store"
	"jiacrontab/server/config"
	"runtime"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// ModelView 数据绑定，view渲染
type ModelView struct {
	StartTime float64
	shareData map[string]interface{}
	locals    map[string]interface{}
	config *config.Config 
	S         *store.Store
	rw        http.ResponseWriter
	
}

func NewModelView(rw http.ResponseWriter, s *store.Store, c *config.Config) *ModelView {
	mv := &ModelView{
		shareData: make(map[string]interface{}),
		locals:    make(map[string]interface{}),
		rw:        rw,
		S:         s,
		config : c,
	}

	mv.locals["appInfo"] = c.AppName + " 当前版本:" + c.Version
	mv.locals["goVersion"] = runtime.Version()
	mv.locals["appName"] = c.AppName
	mv.locals["version"] = c.Version
	return mv
}

func (self *ModelView) Locals(key string, val interface{}) {
	self.locals[key] = val
}

func (self *ModelView) ShareData(key string, val interface{}) {
	self.shareData[key] = val
}

func (self *ModelView) RpcCall(addr string, method string, args interface{}, reply interface{}) error {

	v, ok := self.S.SearchRPCClientList(addr)
	if !ok {
		return fmt.Errorf("not found %s", addr)
	}
	c, err := rpc.NewRpcClient(addr)
	if err != nil {
		self.S.Wrap(func(s *store.Store) {
			v.State = 0
			s.RpcClientList[addr] = v

		}).Sync()
		log.Println(err)
		return err
	}

	if err = c.Call(method, args, reply); err != nil {
		err = fmt.Errorf("failded to call %s %+v %s", method, args, err)
		log.Println(err)
	}
	return err

}

func (self *ModelView) RenderHtml(viewPath []string, locals map[string]interface{}, funcMap template.FuncMap) error {
	var fp []string
	var tplName string
	var tempStart = float64(time.Now().UnixNano())
	for _, v := range viewPath {
		tmp := filepath.Join(".", self.config.TplDir, v+self.config.TplExt)
		fp = append(fp, tmp)
		if tplName == "" {
			tplName = filepath.Base(tmp)
		}

	}
	if locals == nil {
		locals = make(map[string]interface{})
	}
	for k, v := range self.locals {
		locals[k] = v
	}
	t := template.Must(template.New(fp[0]).Funcs(funcMap).ParseFiles(fp...))

	endTime := float64(time.Now().UnixNano())
	tempCostTime := fmt.Sprintf("%.5fms", (endTime-tempStart)/1000000)

	locals["pageCostTime"] = fmt.Sprintf("%.5fms", (endTime-self.StartTime)/1000000)
	locals["tempCostTime"] = tempCostTime
	err := t.ExecuteTemplate(self.rw, tplName, locals)
	if err != nil {
		log.Println(err)
	}
	self.rw.Header().Set("Content-Type", "text/html")
	return err
}

// include user info and template header footer
func (self *ModelView) RenderHtml2(viewPath []string, locals map[string]interface{}, funcMap template.FuncMap) error {
	var fp []string
	var tplName string
	var tempStart = float64(time.Now().UnixNano())
	// pubViews := []string{viewPath, "header", "footer"}
	viewPath = append(viewPath, []string{"public/head", "public/header", "public/footer"}...)
	for _, v := range viewPath {
		tmp := filepath.Join(".", self.config.TplDir, v+self.config.TplExt)
		fp = append(fp, tmp)
		if tplName == "" {
			tplName = filepath.Base(tmp)
		}

	}

	if locals == nil {
		locals = make(map[string]interface{})
	}
	for k, v := range self.locals {
		locals[k] = v
	}

	t := template.Must(template.New(fp[0]).Funcs(funcMap).ParseFiles(fp...))

	endTime := float64(time.Now().UnixNano())
	tempCostTime := fmt.Sprintf("%.5fms", (endTime-tempStart)/1000000)

	locals["pageCostTime"] = fmt.Sprintf("%.5fms", (endTime-self.StartTime)/1000000)
	locals["tempCostTime"] = tempCostTime
	locals["staticDir"] = self.config.StaticDir

	err := t.ExecuteTemplate(self.rw, tplName, locals)

	if err != nil {
		log.Println(err)
	}
	self.rw.Header().Set("Content-Type", "text/html")

	return err
}
