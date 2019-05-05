package main

import (
	"github.com/signintech/gopdf"
	"log"
	"time"
)

const (
	TWO    = 22     // 二号字体
	FOUR   = 14     //四号字体
	W      = 595.28 // 页面宽度
	H      = 841.89 //页面高度
	BH     = 30     // 行距
	PD     = 70     // 段距
	START  = 42
	P1     = "申请执行人："
	P2     = "委托代理人："
	P3     = "被执行人："
	T1     = "请求事项"
	T2     = "事实与理由"
	Blod   = "blod"   //加粗笔触
	Normal = "normal" //普通笔触
	LayOut = "2006年01月02日"
	Font   = 35
)

type Data struct {
	Agent               string    `description:"委托代理人"`
	AgentWork           string    `description:"代理人工作单位"`
	AgentPosition       string    `description:"代理人职位"`
	AgentAccount        string    `description:"代理人电话"`
	PeopleName          string    `description:"被执行人名字"`
	PeopleSex           string    `description:"被执行人性别"`
	PeopleNation        string    `description:"被执行人民族"`
	PeopleBorn          time.Time `description:"被执行人出生年月"`
	PeopleAddress       string    `description:"被执行人住址"`
	PeopleAccount       string    `description:"被执行人电话"`
	PeopleId            string    `description:"被执行人身份证号"`
	ArbitrationDate     time.Time `description:"仲裁裁决书做出之日"`
	Date                string    `description:"仲裁裁决书做出之日年份？"`
	ArbitrationNumber   string    `description:"仲裁号码"`
	Principal           string    `description:"本金"`
	Interest            string    `description:"利息"`
	Penalty             string    `description:"违约金"`
	RepaymentDate       time.Time `description:"还款日的次日"`
	EnforcementDate     time.Time `description:"强制执行申请日"`
	ArbitrationPrice    string    `description:"仲裁费"`
	ArbitrationAccept   string    `description:"受理费"`
	DoubleInterest      string    `description:"加倍利息"`
	AllPrice            string    `description:"借款本金+案件受理费"`
	ArbitrationFullDate time.Time `description:"仲裁裁决书确定的履行期限届满的次日"`
	CourtName           string    `description:"法院名字"`
}

func main() {
	CreatePdf(Data{
		Agent:               "黄希云",
		AgentWork:           "杭州有个金融服务外包有限公司",
		AgentPosition:       "董事长",
		AgentAccount:        "13552780787",
		PeopleName:          "余周同",
		PeopleSex:           "男",
		PeopleNation:        "汉",
		PeopleBorn:          time.Date(1993, 12, 25, 0, 0, 0, 0, time.Local),
		PeopleAddress:       "合肥市高新区古佛我5组59号",
		PeopleAccount:       "15808055474",
		PeopleId:            "500221199308017610",
		ArbitrationDate:     time.Date(2018, 10, 10, 0, 0, 0, 0, time.Local),
		Principal:           "100000",
		Interest:            "1150",
		Penalty:             "4425",
		RepaymentDate:       time.Date(2018, 6, 25, 0, 0, 0, 0, time.Local),
		EnforcementDate:     time.Date(2018, 8, 16, 0, 0, 0, 0, time.Local),
		ArbitrationPrice:    "40",
		ArbitrationAccept:   "40",
		ArbitrationNumber:   "46705",
		DoubleInterest:      "29100",
		AllPrice:            "100040",
		ArbitrationFullDate: time.Date(2018, 8, 18, 0, 0, 0, 0, time.Local),
		CourtName:           "杭州市第一",
		Date:                "2018",
	})
}

//右对齐
func writeInRight(pdf *gopdf.GoPdf, text string) {
	wide, _ := pdf.MeasureTextWidth(text)
	offset := W - START - wide
	pdf.SetX(offset)
	pdf.Cell(nil, text)
	pdf.Br(BH)
}

//分割段落
func cut(str string, firstnum float64) []string {
	r := []rune(str)
	content := make([]string, 0)
	var s string
	j := 0
	times := 0
	sum := 0.0
	flag := false
	for i, _ := range r {
		if r[i] >= 48 && r[i] <= 57 || r[i] >= 97 && r[i] <= 122 || r[i] >= 65 && r[i] <= 90 || r[i] == 40 || r[i] == 41 || r[i] == 14909578 || r[i] == 14909579 {
			sum += 0.5
		} else {
			j++
		}
		s += string(r[i])
		if times == 0 {
			if sum+float64(j) >= firstnum {
				content = append(content, s)
				s = ""
				sum = 0
				j = 0
				times = 1
			}
		} else {
			if sum+float64(j) >= 37 {
				content = append(content, s)
				s = ""
				sum = 0
				j = 0
				flag = true
			} else {
				flag = false
			}
		}
	}
	if times == 0 {
		content = append(content, string(r[:]))
	}
	if !flag {
		content = append(content, s)
		s = ""
	}
	if s != "" {
		content = append(content, s)
	}
	return content
}

//写入
func write(content []string, title string, pdf *gopdf.GoPdf) {
	for i, _ := range content {
		if i == 0 {
			pdf.SetX(PD)
			if title != "" {
				pdf.SetFont(Blod, "", FOUR)
				pdf.Cell(nil, title)
			}
			pdf.SetFont(Normal, "", FOUR)
			pdf.Cell(nil, content[i])
		} else {
			pdf.Br(BH)
			pdf.SetX(START)
			pdf.Cell(nil, content[i])
		}
	}
}

//写入pdf
func writepdf(data Data, pdf *gopdf.GoPdf) {
	pdf.SetX(220)
	pdf.Cell(nil, "强制执行申请书")
	pdf.Br(50)
	pdf.SetX(PD)
	str := "杭州有个金融服务外包有限公司，法定代表人：崔音音，职务：总经理，住所地：浙江省杭州市江干区俞章路88号11幢8楼802室，统一社会信用代码：913301053282397957。"
	content := cut(str, Font-6)
	write(content, P1, pdf)
	pdf.Br(BH)
	str = data.Agent
	str += "，工作单位：" + data.AgentWork
	str += "，职务：" + data.AgentPosition
	str += "，联系电话：" + data.AgentAccount + "。"
	content = cut(str, Font-6)
	write(content, P2, pdf)
	pdf.Br(BH)
	str = data.PeopleName
	str += "，性别：" + data.PeopleSex
	str += "，民族：" + data.PeopleNation + "族"
	str += "，" + data.PeopleBorn.Format(LayOut) + "出生"
	str += "，住" + data.PeopleAddress
	str += "，联系电话：" + data.PeopleAccount
	str += "，身份证号码：" + data.PeopleId + "。"
	content = cut(str, Font-5)
	write(content, P3, pdf)
	pdf.Br(BH)
	str = "申请执行人与被执行人因借款合同纠纷一案，经衢州仲裁委员会审理并于"
	str += data.ArbitrationDate.Format(LayOut) + "作出(" + data.Date + ")衢仲网裁字第" + data.ArbitrationNumber
	str += "号裁决。现被执行人拒不遵照该裁决书内容履行，为此，特申请贵院给予强制执行。"
	content = cut(str, Font)
	write(content, "", pdf)
	pdf.Br(40)
	pdf.SetX(260)
	pdf.SetFont(Blod, "", FOUR)
	pdf.Cell(nil, "请求事项")
	sp := []string{"请求贵院依法强制执行被执行人向申请执行人支付："}
	pdf.Br(BH)
	write(sp[:], "", pdf)
	pdf.Br(BH)
	sp = []string{"（1）借款本金" + data.Principal + "元，利息" + data.Interest + "元；"}
	write(sp[:], "", pdf)
	pdf.Br(BH)
	s := "（2）逾期违约金" + data.Penalty + "元（自" + data.RepaymentDate.Format(LayOut) + "起，以" + data.Principal + "元为基数，按照年利率24%标准计算，暂计算至" + data.EnforcementDate.Format(LayOut) + "，要求计算至被执行人本金还清之日）；"
	content = cut(s, Font)
	write(content, "", pdf)
	pdf.Br(BH)
	sp = []string{"（3）为实现债权支出的仲裁服务费" + data.ArbitrationPrice + "元；"}
	write(sp[:], "", pdf)
	pdf.Br(BH)
	sp = []string{"（4）仲裁案件受理费" + data.ArbitrationAccept + "元；"}
	write(sp, "", pdf)
	pdf.Br(BH)
	s = "（5）迟延履行的加倍利息" + data.DoubleInterest + "元" +
		"（自" + data.ArbitrationFullDate.Format(LayOut) + "起，" +
		"以" + data.AllPrice + "元为基数，" +
		"按照日万分之一点七五为标准，" +
		"暂计算至" + data.EnforcementDate.Format(LayOut) + "，" +
		"要求计算至被执行人付清前述款项之日）。"
	content = cut(s, Font)
	write(content, "", pdf)
	pdf.AddPage()
	pdf.SetFont(Normal, "", FOUR)
	pdf.Br(60)
	pdf.SetFont(Blod, "", FOUR)
	pdf.SetX(260)
	pdf.Cell(nil, "事实与理由")
	pdf.Br(40)
	str = "申请执行人" + "杭州有个金融服务外包有限公司" +
		"与被执行人" + data.PeopleName +
		"借款合同纠纷一案，衢州仲裁委员会于" +
		data.ArbitrationDate.Format(LayOut) + "作出(" + data.Date + ")衢仲网裁字第" + data.ArbitrationNumber + "号裁决，" +
		"该裁决书现已生效。裁决书裁决被执行人柳康在裁决书作出之日起五日内支付申请执行人杭州有个金融服务外包有限公司借款本金" + data.Principal + "元，利息" + data.Interest + "元，" +
		"实现债权而支出的仲裁服务费" + data.ArbitrationPrice + "元，仲裁案件受理费" + data.ArbitrationAccept + "元，并支付自" + data.RepaymentDate.Format(LayOut) + "起，以" + data.Principal + "元为基数，按照年利率24%标准计算至本金还清之日的逾期违约金。"
	content = cut(str, Font)
	write(content, "", pdf)
	pdf.Br(BH)
	str = "裁决生效后，被执行人一直未履行生效法律文书确定的给付义务。" +
		"另根据《中华人民共和国民事诉讼法》第253条之规定，被申请执行人应加倍支付迟延履行期间的债务利息。" +
		"为维护申请人合法权益，根据《中华人民共和国民事诉讼法》之相关规定，特向贵院提出申请强制被执行人支付相应款项。"
	content = cut(str, Font)
	write(content, "", pdf)
	pdf.Br(BH)
	pdf.SetX(PD)
	pdf.Cell(nil, "此致")
	pdf.Br(BH)
	pdf.SetX(START)
	//参数没有传

	pdf.Cell(nil, data.CourtName+"中级人民法院")
	pdf.Br(BH)
	writeInRight(pdf, "申请人：杭州有个金融服务外包有限公司")
	writeInRight(pdf, time.Now().Format(LayOut))
	pdf.WritePdf("强制执行申请书.pdf")
}

//创建pdf文档
func CreatePdf(data Data) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: W, H: H}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont(Blod, "msyhbd.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.AddTTFFont(Normal, "SIMYOU.TTF")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont(Normal, "", FOUR)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Br(60)
	err = pdf.SetFont(Blod, "", TWO)
	if err != nil {
		log.Print(err.Error())
		return
	}
	writepdf(data, &pdf)
}
