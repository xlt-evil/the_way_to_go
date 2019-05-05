package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问路由
	http.HandleFunc("/login", login)   //设置访问路由
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello my master")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("D:/HxyGo/src/the-way-to-go/learn_7/login.html")
		t.Execute(w, token)
	} else { //form["key"] 是slice类型的
		//fmt.Println("username",r.Form["username"])
		//fmt.Println("password",r.Form["password"])
		fmt.Println("username", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("username", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //默认过滤html标签
	}
}

//上传图片文件并copy到服务器下
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println(666)
	fmt.Println(r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("D:/HxyGo/src/the-way-to-go/learn_7/upload.html")
		fmt.Println(777)
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20) //超出内存剩下的文件存在系统内部中
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test"+handler.Filename, os.O_WRONLY|os.O_CREATE, 066)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintf(w, "图片上传成功")
	}
}

func setCookie(w http.ResponseWriter, r *http.Request) {

	time1 := time.Now()
	time1.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "hxy", MaxAge: 86400}
	http.SetCookie(w, &cookie)
	w.Header().Set("Set-Cookie", cookie.String())
	w.Header().Add("Set-Cookie", cookie.String())
	fmt.Fprintf(w, "SetCookie")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	name, _ := r.Cookie("username")
	fmt.Fprintf(w, name.String())
}
