package util

import (
	// "html/template"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gocraft/web"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./templates/*.html",
	}

	rnd = renderer.New(opts)
}

func RenderTemplate(w web.ResponseWriter, temphtml string, params interface{}) {

	post := make(map[string]interface{})

	vFields := reflect.TypeOf(params).Elem()
	rr := reflect.ValueOf(params).Elem()
	for i := 0; i < vFields.NumField(); i++ {
		fieldName := vFields.Field(i).Name

		ff := rr.FieldByName(fieldName)
		ffValue := ff.Interface()

		post[fieldName] = ffValue;
		// send.(map[string]interface{})[fieldName] = ffValue
	}

	r := reflect.ValueOf(params).Elem()
	f := r.FieldByName("PageRenderStart")
	fieldValue := f.Interface()

	renderStart := fieldValue.(time.Time)
	renderEnd := time.Now()
	renderTimeSeconds := uint64(renderEnd.Sub(renderStart).Nanoseconds() / 1000000)

	f = r.FieldByName("PageRenderTime")
	// make sure that this field is defined, and can be changed.
	if f.IsValid() && f.CanSet() && f.Kind() == reflect.Uint64 {
		f.SetUint(renderTimeSeconds)
	}
	
	rnd.HTML(w, http.StatusOK, temphtml, post)
}


func APIResponse(w web.ResponseWriter, r *web.Request, params interface{}) {
	json, err := json.MarshalIndent(params, "", "\t")
	if err != nil {
		http.NotFound(w, r.Request)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(json)
	w.Write(json)
}

func RedirectOrAPIResponse(w web.ResponseWriter, r *web.Request, redirectUrl string, params interface{}) {
	if len(r.URL.Query()["json"]) > 0 {
		APIResponse(w, r, params)
	} else {
		http.Redirect(w, r.Request, redirectUrl, 302)
		return
	}
}
