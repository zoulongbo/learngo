package model

import "encoding/json"

type Profile struct {
	Name          string		//姓名
	Gender        string		//性别
	Age           int			//年龄
	Height        int			//身高
	Weight        int			//体重
	Income        string		//收入
	Marriage      string		//婚否
	Education     string		//学历
	Occupation    string		//职业
	Native        string		//籍贯
	Constellation string		//星座
	House         string		//房否
	Car           string		//车否
}

func FromJsonObj(o interface{}) (Profile, error)  {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
} 
