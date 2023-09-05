package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// ``` text
// 一班 张小丙 第3名
// 一班 张小甲 第1名
// 一班 张小乙 第2名

// 二班 王七六 第5名
// 二班 王九七 第1名
// 二班 胡八一 第2名
// 二班 王六零 第6名
// 二班 刘八一 第2名
// 二班 李八一 第2名
// ```

// StudentScore 学生成绩
type StudentScore struct {
	Name  string
	Score int
}

// printRank 获取班级排名并打印
func printRank() {
	jsonData := readJson("./score.json")
	dataMap := make(map[string]map[string]int)
	err := json.Unmarshal(jsonData, &dataMap)
	if err != nil {
		fmt.Println("error unmarshal json data")
		return
	}
	for k, v := range dataMap {
		sortPrint(k, v)
		fmt.Printf("\n")
	}
}

// 按顺序输出学生成绩
func sortPrint(class string, scoreMap map[string]int) {
	sortList := []*StudentScore{}
	for k, v := range scoreMap {
		item := &StudentScore{
			Name:  k,
			Score: v,
		}
		sortList = append(sortList, item)
	}
	sort.Slice(sortList, func(i, j int) bool {
		return sortList[i].Score > sortList[j].Score
	})
	count := 0
	score := 0
	// 60 80 80 90
	for i := 0; i < len(sortList); i++ {
		rank := i + 1
		if sortList[i].Score == score {
			rank = count
		} else {
			count++
			score = sortList[i].Score
		}
		fmt.Printf("%s %s 第%d名\n", class, sortList[i].Name, rank)
	}
}

// readJson 读取Json数据
func readJson(filePath string) []byte {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening json file")
		return nil
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading json file")
		return nil
	}
	return jsonData
}
