package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

//模板的使用
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

type Friend struct {
	Fname string
}

func main() {
	nestingTemplate()
}

func easyTempateUse() {
	t := template.New("filename example") //创建一个模板
	t, _ = t.Parse("hello {{.UserName}}") //小写字母是不能导出的
	p := Person{UserName: "hxy", Emails: []string{"958752538@qq.com"}}
	t.Execute(os.Stdout, p) //将p导出
}

func LoopTemplateUse() { //模板循环
	f1 := Friend{Fname: "miunx.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname examples")
	t, _ = t.Parse(`hello {{.UserName}}!
					{{range .Emails }}
						an email {{.}}
					{{end}}
					{{with .Friends}}//将改变的值赋值给.
					{{ range .}}
						my friend name is {{.Fname}}
					{{end}}
					{{end}}
				`)
	p := Person{UserName: "hxy", Emails: []string{"qqwe@qwe.com", "dfas@163.com"}, Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func IfTemplateUse() {
	t := template.New("template test")
	//Must用与检测模板是否正确
	t = template.Must(t.Parse("空 pipeline if demo :{{if ``}} 不会输出 {{end}}"))
	t.Execute(os.Stdout, nil)
	t1 := template.New("template test")
	t1 = template.Must(t1.Parse("不为空的 pipeline if demo: {{if `any`}}qwe{{end}}"))
	t1.Execute(os.Stdout, nil)
	t2 := template.New("template test")
	t2 = template.Must(t2.Parse("if-else demo:{{if `anything`}}if 部分{{else}}haha{{end}}"))
	t2.Execute(os.Stdout, nil)
}

//pipelines 的使用 在go中任何 {{}}都是pipelines的数据
func variablePipeUse() {
	t := template.New("template test")
	t, _ = t.Parse(`{{with $k := "output" | printf "1+%s" }}
							{{$k}}
						{{end}}
					`) //with在他的作用域内取别名
	t.Execute(os.Stdout, nil)
}

func funcTemplate(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	return (substrs[0] + "at" + substrs[1])
}

//模板函数使用
func UseTemplateFunc() {
	f1 := Friend{Fname: "miunx.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("tempalte test")
	t = t.Funcs(template.FuncMap{"emailDeal": funcTemplate}) //注册函数
	t, _ = t.Parse(`hello {{.UserName}}!
					{{range .Emails}}
							an emails {{.|emailDeal}}
					{{end}}
					{{with .Friends}}
					{{range .}}
						my friend name is {{.Fname}}
					{{end}}
					{{end}}
				`)
	p := Person{UserName: "hxy", Emails: []string{"hxy@163.com", "hxy@gmail.com"}, Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

//子模板嵌套使用
func nestingTemplate() {
	s1, _ := template.ParseFiles("D:/HxyGo/src/the-way-to-go/learn_10/header.tmpl", "D:/HxyGo/src/the-way-to-go/learn_10/content.tmpl", "D:/HxyGo/src/the-way-to-go/learn_10/footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)

}
