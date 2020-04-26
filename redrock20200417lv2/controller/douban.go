package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"redrock20200417lv2/mysql"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Douban(c *gin.Context) {
	t1 := time.Now() // get current time
	db := mysql.DbConn()
	parse(db)
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)
}

//定义 Spider get的方法
func (keyword Spider) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
		fmt.Println("err1:", err)
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err2:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err3:", err)
	}
	return string(body)
}

func parse(db *gorm.DB)  {
	header := map[string]string{
		"Host": "movie.douban.com",
		"Connection": "keep-alive",
		"Cache-Control": "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer": "https://movie.douban.com/top250",
	}

	//循环每页解析
	for i := 0; i < 10; i++{
		fmt.Println("正在抓取第"+strconv.Itoa(i+1)+"页......")
		url := "https://movie.douban.com/top250?start="+strconv.Itoa(i*25)+"&filter="
		fmt.Println(url)
		spider := &Spider{url, header}
		html := spider.get_html_header()

		//评语
		reBody  := strings.ReplaceAll(html,"\n","")//去除所有换行 用""替代
		s       := strings.ReplaceAll(reBody, " ", "")//去除所有空格 用""替代
		par     := `<divclass="item"><divclass="pic">(.*?)</div></div></div>`
		liReg   := regexp.MustCompile(par)
		find    := liReg.FindAllStringSubmatch(s, -1)

		//导演
		reBody0  := strings.ReplaceAll(string(html),"\n","")
		ss       := strings.ReplaceAll(reBody0, " ", "")
		pattern0 := `导演:(.*?)&`
		rp0      := regexp.MustCompile(pattern0)
		find_d   := rp0.FindAllStringSubmatch(ss,-1)

		//图片
		pattern1 := `src="(.*?)" class`
		rp1      := regexp.MustCompile(pattern1)
		find_Jpg := rp1.FindAllStringSubmatch(html, -1)

		//评价人数
		pattern2:=`<span>(.*?)人评价</span>`
		rp2 := regexp.MustCompile(pattern2)
		find_Num := rp2.FindAllStringSubmatch(html,-1)

		//评分
		pattern3 := `property="v:average">(.*?)</span>`
		rp3 := regexp.MustCompile(pattern3)
		find_Score := rp3.FindAllStringSubmatch(html,-1)

		//电影名称
		pattern4 := `alt="(?s:(.*?))"`
		rp4 := regexp.MustCompile(pattern4)
		find_Name := rp4.FindAllStringSubmatch(html,-1)

		//names := rp4.FindAllString(find_txt4[0][0], -1)
		//fmt.Println(find_txt2)
		//fmt.Println(find_txt3)
		//fmt.Println(find_txt4)
		//fmt.Println(names)
		//// 写入UTF-8 BOM
		//f.WriteString("\xEF\xBB\xBF")
		////  打印全部数据和写入excel文件
		//for i:=0;i<len(find_txt2);i++{
		//	fmt.Printf("%s %s %s\n",find_txt4[i][1],find_txt3[i][1],find_txt2[i][1], )
		//	f.WriteString(find_txt4[i][1]+"\t"+find_txt3[i][1]+"\t"+find_txt2[i][1]+"\t"+"\r\n")
		//
		//}
		Save(find_Name, find_Num, find_Score, find_Jpg, find_d, find, db)
	}
}

func Save(find_Name, find_Num, find_Score, find_Jpg, find_d, find [][]string, db *gorm.DB)  {
	movies := []Movie{}
	for i := 0; i < 25; i++ {
		par1    := `<spanclass="inq">(.*?)</span>`
		li      := regexp.MustCompile(par1)
		find_P  := li.FindAllStringSubmatch(find[i][1], -1)
		var str string
		if find_P == nil{
			str = "暂无短评"
		}else {
			str = find_P[0][1]
		}
		score, err := strconv.ParseFloat(find_Score[i][1], 64)
		if err != nil{
			fmt.Println("score err:", err)
		}
		movies := append(movies, Movie{
			Name:  find_Name[i][1],
			Num:   find_Num[i][1],
			Score: score,
			Jpg:   find_Jpg[i][1],
			P:     str,
			D:     find_d[i][1],
		})
		fmt.Println(movies)
		db.AutoMigrate(&movies)
		db.Create(&Movie{
			Name:  find_Name[i][1],
			Num:   find_Num[i][1],
			Score: score,
			Jpg:   find_Jpg[i][1],
			P:     str,
			D:     find_d[i][1],
			Ok:    1,
		})

	}
}
