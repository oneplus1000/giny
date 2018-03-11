package giny

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

//PubFileMgr manage public file like .js .css
type PubFileMgr struct {
	pubPath       string
	pubVirPath    string
	scriptVersion string
	jsfiles       map[string]([]string)
	cssfiles      map[string]([]string)
}

//ErrDupJsName dup js name
var ErrDupJsName = errors.New("dup JsBundle name")

//ErrDupCSSName dup cs name
var ErrDupCSSName = errors.New("dup CssBundle name")

//NewPubFileMgr new PubFileMgr
func NewPubFileMgr(pubPath string, pubVirPath string, scriptVersion string) *PubFileMgr {
	var pubMgr PubFileMgr
	pubMgr.pubPath = pubPath
	pubMgr.pubVirPath = pubVirPath
	pubMgr.scriptVersion = scriptVersion
	pubMgr.jsfiles = make(map[string]([]string))
	pubMgr.cssfiles = make(map[string]([]string))
	return &pubMgr
}

//RegistJs regist javascript bundle
func (p *PubFileMgr) RegistJs(name string, files ...string) error {
	if _, ok := p.jsfiles[name]; ok {
		return ErrDupJsName
	}
	p.jsfiles[name] = files
	return nil
}

//RegistCSS regist css bundle
func (p *PubFileMgr) RegistCSS(name string, files ...string) error {
	if _, ok := p.cssfiles[name]; ok {
		return ErrDupCSSName
	}
	p.cssfiles[name] = files
	return nil
}

//JsBundle func for html/template add js bundle to page
func (p *PubFileMgr) JsBundle(data interface{}, name string) template.HTML {
	if files, ok := p.jsfiles[name]; ok {
		var buff bytes.Buffer
		buff.WriteString(fmt.Sprintf("<!-- JsBundle %s -->\n", name))
		for _, file := range files {
			buff.WriteString(fmt.Sprintf("<script src=\"/%s/%s?version=%s\" ></script>\n", p.pubVirPath, file,p.scriptVersion))
		}
		return template.HTML(buff.String())
	}
	return template.HTML(fmt.Sprintf("<!-- Error:JsBundle %s not found  -->\n", name))
}

//CSSBundle func for html/template add css bundle to page
func (p *PubFileMgr) CSSBundle(data interface{}, name string) template.HTML {
	if files, ok := p.cssfiles[name]; ok {
		var buff bytes.Buffer
		buff.WriteString(fmt.Sprintf("<!-- CssBundle %s -->\n", name))
		for _, file := range files {
			buff.WriteString(fmt.Sprintf("<link rel=\"stylesheet\"  href=\"/%s/%s?version=%s\" />\n", p.pubVirPath, file,p.scriptVersion))
		}
		return template.HTML(buff.String())
	}
	return template.HTML(fmt.Sprintf("<!-- Error:CssBundle %s not found  -->\n", name))
}

//Static setup static file path
func (p *PubFileMgr) Static(router *gin.Engine) {
	router.Static(p.pubVirPath, p.pubPath)
}
