package main

/*
怎么理解  interface
- interface 是方法的集合
- 从实现上来理解，任何对象实现再 interface 里面的所有方法，表明其实现了该接口
- interface 可以作为一种数据类型，它可以作为任意数据类型
*/

type Phone interface {
	Price() uint32
	Brand() string
}

type Iphone struct {
	Logo string
}

type XiaoMi struct {
	Logo string
}

func (ip *Iphone) Price() uint32 {
	return 1000
}

func (ip *Iphone) Brand() string {
	return ip.Logo + "iphone"
}

func (xm *XiaoMi) Price() uint32 {
	return 500
}

func (xm *XiaoMi) Brand() string {
	return xm.Logo + "小米"
}

func main() {

}
