package parser

import (
	"github.com/zoulongbo/learngo/crawler/engine"
	"github.com/zoulongbo/learngo/crawler/model"
	"github.com/zoulongbo/learngo/crawler_distributed/config"
	"regexp"
	"strconv"
	"strings"
)

var desRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)
//var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-11592496>([^<]+)</h1>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)cm</div>`)
//var weightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var constellationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+座)([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+车)</div>`)
var nativeRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^<]+)</div>`)
var genderRe = regexp.MustCompile(`<div class="m-title" data-v-8b1eac0c>([^<]+)的动态</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	//取简介 des 0=城市 1=年龄 2=学历 3=marriage 4=身高 5=收入
	desc := extractString(contents, desRe)
	if desc != "" {
		des := strings.Split(desc, " | ")
		profile.Education = des[2]
		profile.Marriage = des[3]
		profile.Income = des[5]
	}
	//年龄 单独取
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if  err == nil {
		profile.Age = age
	}
	//身高 单独取
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if  err == nil {
		profile.Height = height
	}
	//名字
	//profile.Name = extractString(contents, nameRe)
	profile.Name = name
	//籍贯
	profile.Native = extractString(contents, nativeRe)
	//性别
	profile.Gender = convertGender(extractString(contents, genderRe))
	//房况
	profile.House = extractString(contents, houseRe)
	//车况
	profile.Car = extractString(contents, carRe)
	//职业
	math := occupationRe.FindAllSubmatch(contents, -1)
	if len(math) > 2 {
		profile.Occupation = string(math[len(math) -2:][0][1])
	}
	//星座
	profile.Constellation = extractString(contents, constellationRe)

	result := engine.ParseResult{
		Items:[]engine.Item{
			{
				Url:url,
				Id:extractString([]byte(url), idUrlRe),
				Type:"zhenai",
				Payload:profile,
			},
		},
	}
	return result
}

func convertGender(title string) string {
	if  strings.Compare(title, "她") < 0  {
		return "男士"
	} else {
		return "女士"
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	math := re.FindSubmatch(contents)
	if len(math) > 1 {
		return string(math[1])
	} else {
		return ""
	}
}


//特殊 参数解析器方法 接口实现
type ProfileParser struct {
	username string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.username)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.username
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{username:name}
}

