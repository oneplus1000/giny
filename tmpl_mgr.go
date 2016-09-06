package giny

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//ErrTmplNotFound tmpl not found
var ErrTmplNotFound = errors.New("tmpl not found")

//TmplMgr template manager
type TmplMgr struct {
	tmplsPath string
	viewinfos map[string]ViewInfo           //ข้อมูลสำหรับสร้าง (compile) template
	views     map[string]*template.Template //tmeplate ที่ compile ไว้แล้ว
	isDebug   bool
	tmplFuncs map[string]interface{}
}

//NewTmplMgr new TmplMgr
func NewTmplMgr(tmplsPath string, isDebug bool) *TmplMgr {
	var tmplMgr TmplMgr
	//int
	tmplMgr.tmplsPath = tmplsPath
	tmplMgr.viewinfos = make(map[string]ViewInfo)
	tmplMgr.views = make(map[string]*template.Template)
	tmplMgr.tmplFuncs = make(map[string]interface{})
	tmplMgr.isDebug = isDebug
	return &tmplMgr
}

//RegistFunc regit html/template function
func (t *TmplMgr) RegistFunc(name string, fn interface{}) {
	t.tmplFuncs[name] = fn
}

//RegistTmpl regit template
func (t *TmplMgr) RegistTmpl(name string, startTmpl string, files ...string) {
	var vinfo ViewInfo
	vinfo.startTmpl = startTmpl
	vinfo.files = files
	t.viewinfos[name] = vinfo
}

//MakeAll create all template
func (t *TmplMgr) MakeAll() error {
	for name, vinfo := range t.viewinfos {
		tmpl, err := t.Make(name, vinfo)
		if err != nil {
			return err
		}
		t.views[name] = tmpl
	}
	return nil
}

//Make create template
func (t *TmplMgr) Make(name string, vinfo ViewInfo) (*template.Template, error) {
	var files []string
	for _, f := range vinfo.files {
		files = append(files, filepath.Join(t.tmplsPath, f))
	}
	tmpl, err := template.New(name).Delims("[[", "]]").Funcs(t.tmplFuncs).ParseFiles(files...)
	if err != nil {
		log.Printf("err : %s", err.Error())
		return nil, err
	}
	return tmpl, nil
}

//RenderWithCtx reander template
func (t *TmplMgr) RenderWithCtx(
	ctx *gin.Context,
	name string,
	data interface{},
) error {
	return t.Render(ctx.Writer, ctx.Request, name, data)
}

//Render reander template
func (t *TmplMgr) Render(
	w io.Writer,
	r *http.Request,
	name string,
	data interface{},
) error {
	if vinfo, ok := t.viewinfos[name]; ok {
		if t.isDebug { //ถ้าอยู่ใน debug mode make tmpl ใหม่ทุกครั้ง
			view, err := t.Make(name, vinfo)
			if err != nil {
				fmt.Fprintf(w, "TmplMgr.Render  : %s", err.Error()) //echo error
				return err
			}
			err = view.ExecuteTemplate(w, vinfo.startTmpl, ViewModel{
				R: r,
				D: data,
			})

			if err != nil {
				fmt.Fprintf(w, "TmplMgr.Render  : %s", err.Error()) //echo error
				return err
			}

			return nil
		} else { //ถ้าอยู่ใน deploy mode ใช้ tmpl ที่ compile ไว้แล้ว
			if view, ok := t.views[name]; ok {
				err := view.ExecuteTemplate(w, vinfo.startTmpl, ViewModel{
					R: r,
					D: data,
				})
				if err != nil {
					fmt.Fprintf(w, "TmplMgr.Render  : %s", err.Error()) //echo error
					return err
				}
				return nil
			}
		}
	}
	fmt.Fprintf(w, "TmplMgr.Render error : %s", ErrTmplNotFound.Error()) //echo error
	return ErrTmplNotFound
}

//ViewInfo view info
type ViewInfo struct {
	startTmpl string
	files     []string
}
