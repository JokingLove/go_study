package main

import "fmt"

func main() {
	demo1()
	demo2()
}

func demo2() {
	countryCapitaMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	for country := range countryCapitaMap {
		fmt.Println(country, "首都是", countryCapitaMap[country])
	}

	delete(countryCapitaMap, "France")

	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后的地图")

	for country := range countryCapitaMap {
		fmt.Println(country, "首都是", countryCapitaMap[country])
	}
}

func demo1() {

	var siteMap map[string]string
	siteMap = make(map[string]string)

	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	name, ok := siteMap["Facebook"]

	if ok {
		fmt.Println("Facebook 的站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}
}
