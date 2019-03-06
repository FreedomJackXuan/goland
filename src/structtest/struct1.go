package structtest

import "fmt"

type AnimalCategory struct {
	kingdom string // 界
	phylum string // 门
	class string // 纲
	order string // 目
	family string // 科
	genus string // 属
	species string // 种

}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory // 动物基本分类。
}


func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

func (a Animal) Category() string{
	return a.AnimalCategory.String()
}

func Struct1(){
	cat := AnimalCategory{species:"cat"}
	animal:=Animal{
		scientificName:"american shorthair",
		AnimalCategory:cat,
	}
	fmt.Println("the animal",animal)
}
