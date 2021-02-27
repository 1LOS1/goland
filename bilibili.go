package main

import "fmt"



type video struct {		//一个视频的参数
	good int
	coin int
	cet int
}

type mouse interface {		//一个名为鼠标的接口
	makegood()		//点赞方法
	givecoin()		//投币方法
	makecet()		//收藏方法
}

func (mouse *video)makegood()  {			//video实现点赞方法
	mouse.good++
}

func (mouse *video)givecoin()  {			//video实现投币方法
	mouse.coin++
}

func (mouse *video)makecet()  {				//video实现收藏方法
	mouse.cet++
}

func main()  {
	 LOS := &video{100,20,50}		//一个叫LOS 的视频原来有100个赞，20个币，50个收藏
	var i mouse					//一个名为i的鼠标接口
	i = LOS						//
	i.makegood()				//给我点个赞
	fmt.Println("los = ",LOS)//验证实现了点赞

}
