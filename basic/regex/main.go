package main

import (
	"fmt"
	"regexp"
)
//								.任意字符 ｜ + 多个 ｜ () 提取子串
//								. +
const text1  = `This is my email 819243521@qq.com ,ok
This is email1 121@sina.com  111
This is email2 n2d22@google.com.cn 12
`
const textRg1  = `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`

const text2  = `<div data-v-11592496="" class="des f-cl">天津 | 33岁 | 中专 | 离异 | 158cm | 3001-5000元<a data-v-11592496="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1149119961" target="_self" class="online f-fr">查看最后登录时间</a></div>`
const textRg2  = `<div data-v-11592496="" class="des f-cl">([^<]+)<a`

const text3  = `<div class="m-btn purple" data-v-8b1eac0c>离异</div><div class="m-btn purple" data-v-8b1eac0c>50岁</div><div class="m-btn purple" data-v-8b1eac0c>魔羯座(12.22-01.19)</div><div class="m-btn purple" data-v-8b1eac0c>171cm</div><div class="m-btn purple" data-v-8b1eac0c>工作地:天津宝坻区</div><div class="m-btn purple" data-v-8b1eac0c>月收入:3-5千</div><div class="m-btn purple" data-v-8b1eac0c>销售总监</div><div class="m-btn purple" data-v-8b1eac0c>高中及以下</div>`
const textRg3  = `<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`
func main() {
	reg := regexp.MustCompile(textRg3)
	//match := reg.FindString(text)	//单个
	//match := reg.FindAllString(text2, -1)	//全部个
	//match := reg.FindAllStringSubmatch(text2, -1)	//全部个
	//for _, m := range match {
	//	fmt.Println(m[1])
	//}
	match := reg.FindAllStringSubmatch(text3, -1)
	//match1 := match[len(match)-2:][0]
	fmt.Println(match)

}
