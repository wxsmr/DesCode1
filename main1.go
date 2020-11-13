package main

import "fmt"

func main() {
	person1 := NewChinese()
	age := person1.Age()
	if age < 80 {
		fmt.Println("在年龄上符合择偶标准")
	} else {
		fmt.Println("年龄不符合择偶标准")
	}
	//person1.IsMajiang

	person2 := NewJapanese()
	person2.Age()
	//person2.Php

}

type Person interface {
	Shanliang()    bool
	weiRenChuShi()
	Height()       int
	weight()       int
	Age()          int
	Salary()       int
}

type Chinese struct {
	Name        string
	Sex         string
	IsShanliang bool
	High        int   //身高
	Wei         int   //体重
	AgeNum      int  //年龄
	Money       int  //挣钱能力
	IsMajiang   bool //打或者不打麻将
}

func NewChinese()*Chinese  {
	c:=&Chinese{
		Name:        "名媛",
		Sex:         "女",
		IsShanliang: true,
		High:        170,
		Wei:         98,
		AgeNum:      21,
		Money:       300000,
		IsMajiang:   false,
	}
	return c
}

func (c *Chinese)Shanliang()bool  {
	return c.IsShanliang
}

func (c *Chinese)weiRenChuShi()  {
	fmt.Println(c.Name+"伟人处理能力很好")
}

func (c *Chinese)Height()int  {
	return c.High
}

func (c *Chinese)Weight()int  {
	return c.Wei
}
func (c *Chinese)Age()int  {
	return c.AgeNum
}

func (c *Chinese)Salary()int  {
	return c.Money
}

func (c *Chinese)Majiang()  {
	if c.IsMajiang {
		fmt.Println(c.Name + "会打麻将")
	}else {
		fmt.Println(c.Name + "不会打麻将")
	}
}

type Japanese struct {
	Name        string
	AgeNum      int
	IsShanliang bool
	High        int
	wei         int
	Money       int
	Php         bool
}

func NewJapanese()*Japanese  {
	j:=&Japanese{
		Name:        "桥本淦江华",
		AgeNum:      23,
		IsShanliang: true,
		High:        165,
		wei:         90,
		Money:       20000000,
		Php:         true,
	}
	return  j
}

func (j *Japanese)shanliang()bool  {
	return j.IsShanliang
}

func (j *Japanese)WeiRenChuShi()  {
	fmt.Println(j.Name+"为人处事能力好")
}

func (j*Japanese)Height()int  {
	return j.High
}

func (j *Japanese)Weight()int  {
	return j.wei
}

func (j *Japanese)Age()int  {
	return j.AgeNum
}

func (j *Japanese)Salary()int  {
	return j.Money
}

func (j *Japanese)Eat(food string)  {
	fmt.Println(j.Name+"喜欢吃:",food)
}

func (j *Japanese)Phps()bool  {
	return j.Php
}