package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	cost   float64
	weight float64
}

/*
Первая строка содержит количество предметов и вместимость рюкзака
Каждая из следующих n строк задаёт стоимость и объём предмета
Выведите максимальную стоимость частей предметов
(от каждого предмета можно отделить любую часть, стоимость и объём при этом пропорционально уменьшатся),
помещающихся в данный рюкзак, с точностью не менее трёх знаков после запятой.
*/

func continuousBackpack() {
	var itemsCount, backpackCapacity float64
	fmt.Scan(&itemsCount, &backpackCapacity)
	var items []Item
	sc := bufio.NewScanner(os.Stdin)
	var i float64
	for i = 0; i < itemsCount; i++ {
		sc.Scan()
		items = append(items, makeItem(sc.Text()))
	}

	var result float64

	sortItems(items)

	for _, v := range items {
		if backpackCapacity == 0 {
			break
		}

		if v.weight <= backpackCapacity {
			backpackCapacity = backpackCapacity - v.weight
			result += v.cost
		} else {
			oneItemCost := v.cost / v.weight
			canTakeCost := backpackCapacity * oneItemCost
			backpackCapacity = 0
			result += canTakeCost
		}
	}

	fmt.Printf("%.3f", result)

}

func makeItem(input string) Item {
	nums := strings.Split(input, " ")
	itemCost, _ := strconv.ParseFloat(nums[0], 64)
	itemWeight, _ := strconv.ParseFloat(nums[1], 64)
	return Item{cost: itemCost, weight: itemWeight}
}

func sortItems(items []Item) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].cost/items[i].weight > items[j].cost/items[j].weight
	})
}
