package main

import "C"
import (
	"fmt"
)

func main() {
	/*接口和结构体之间的联系和使用规范
	  接口：一套标准，适合统一定义共性，抽离并定义出一套标准
	  结构体：实体的描述和定义
	*/
	//接口：有统一的共性进行判断是，往往使用的是接口
	//
	Person1 := NewChinese()
	age := Person1.Agenum()
	if age<23 {
		fmt.Println("年龄上符合")
	}else {
		fmt.Println("年龄上不符合")
	}
}

type Person interface {
	Shanliang() bool
	WRCS()      bool
	Height()    int
	Weight()    int
	Age()       int
	Salary()    int
}

type Chinese struct {
	Name string
	Sex  string
	IsShanliang bool
	High int
	Wei  int
	AgeNum int
	Money int
}

func NewChinese() *Chinese  {
	c := &Chinese{
		Name:        "x",
		Sex:         "女",
		IsShanliang:  true,
		High:        164,
		Wei:         97,
		AgeNum:      20,
		Money:       10000,
	}
	return c
}

func (c *Chinese)Shanliang() bool {
	return c.IsShanliang
}
func (c *Chinese)WRCS()  {
	fmt.Println(C.name+"wsrs很好")
}
func (c *Chinese) Hight() int  {
	return c.High
}

func (c *Chinese)Weigh() int {
	return c.Wei
}

func (c *Chinese)Agenum()int  {
	return c.AgeNum
}

func (c *Chinese)Moneynum()int  {
	return c.Money
}


type USA struct {
	Name string
	Sex string
	Isshanliang bool
	Height int
	weight int
	Age  int
	Money int
}

func NewUsa() *USA {
	u := &USA{
		Name:        "y",
		Sex:         "女",
		Isshanliang: true,
		Height:      178,
		weight:      110,
		Age:         21,
		Money:       300000,
	}
	return u
}

func (u *USA)name() string  {
	return u.Name
}

func (u *USA)sex()string  {
	return u.Sex
}

func (u * USA)shanliang() bool {
	return u.shanliang()
}

func (u *USA)heih() int  {
	return u.Height
}

func (u *USA)wei()int  {
	return u.weight
}

func (u *USA)agenum()int  {
	return u.Age
}

func (u *USA)moneynum()int  {
	return u.Money
}