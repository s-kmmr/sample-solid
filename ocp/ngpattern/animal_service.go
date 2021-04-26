package ngpattern

func MakeSound(animal Animal) {
	switch animal.Name {
	case "dog":
		println("ワンワン")
	case "cat":
		println("ニャーニャー")
		//   ・・・牛を追加するとした際にここをいじらないと・・・
	}
}
