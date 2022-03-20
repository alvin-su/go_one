package main

import "fmt"

func main() {
	//基于数组创建切片
	var foods = []string{"test1", "test2", "test3", "test4", "test5"}
	var foodsSlice []string = foods[0:3]
	fmt.Println(foodsSlice)
	fmt.Println(foods[:2])
	fmt.Println(foods[3:])

	//直接创建切片
	var foodsSlice2 []string = []string{"测试1", "测试2", "测试3", "测试4", "测试5"}
	fmt.Println(foodsSlice2)

	//使用内置函数make创建
	//foodsSlice3:=make([]string,6)
	//foodsSlice4:=make([]string,6,8)

	foodsSlice5 := foodsSlice2[1:4]
	fmt.Println(len(foodsSlice5))
	fmt.Println(cap(foodsSlice5))

	foodsSlice4 := []string{"红烧肉"}
	fmt.Println(foodsSlice4)
	fmt.Println(fmt.Sprintf("Len:%d", len(foodsSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodsSlice4)))

	foodsSlice4 = append(foodsSlice4, "清蒸鱼")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodsSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodsSlice4)))

	foodsSlice4 = append(foodsSlice4, "溜大虾")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodsSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodsSlice4)))

	foodsSlice4 = append(foodsSlice4, "溜大虾22")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodsSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodsSlice4)))

	foodsSlice4 = append(foodsSlice4, "溜大虾33")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodsSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodsSlice4)))

	foodsSlice6 := make([]string, 5, 8)
	foodsSlice6[0] = "t1"
	foodsSlice6[1] = "t2"
	foodsSlice6[2] = "t3"
	foodsSlice6[3] = "t4"
	foodsSlice6[4] = "t5"
	sliceForach(foodsSlice6)

	//复制切片 [:]
	var other_foods = foods[:]
	fmt.Println(other_foods)

	foodsSlice7 := []string{"米饭", "面条"}
	copy(foodsSlice7, foodsSlice6)
	fmt.Println(foodsSlice6)
	fmt.Println(foodsSlice7)

	foods8 := []string{"s1", "s2", "s3", "s4", "s5"}
	/*foods8 = append(foods8, make([]string, 5)...)
	fmt.Printf("foods8长度:%d\n", len(foods8))
	fmt.Printf("容量:%d\n", cap(foods8))*/

	//在切片foods8索引2的位置插入元素 ss3
	foods8 = append(foods8[:2], append([]string{"ss3"}, foods8[2:]...)...)
	fmt.Println(foods8)
	foods8 = append(foods8[:4], append(make([]string, 3), foods8[4:]...)...)
	fmt.Println(foods8)

	//从切片foods9中删除 元素 t4
	foods9 := []string{"t1", "t2", "t3", "t4", "t5"}
	foods10 := append(foods9[:3], foods9[4:]...)
	fmt.Println(foods10)

	var foods11 = []string{"红烧肉", "清蒸鱼", "溜大虾", "蒸螃蟹", "鲍鱼粥"}
	fmt.Println(foods11[:3])
	fmt.Println(foods11[4:])
	foods12 := append(foods11[:3], foods11[4:]...)
	fmt.Println(foods12)
}
