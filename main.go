package main

import "fmt"

func main() {
	fmt.Println("Расчет вашего одноповторного максимума в паурлифтинге и приблизительный подсчёт разряда")
	var weight float32
	var benchReps float32
	var benchP float32
	var deadl float32
	var deadliftReps float32
	var squatt float32
	var squatReps float32
	var squat float32
	var deadlift float32
	var benchPress float32
	fmt.Print("Напишите ваш вес в килограммах: ")
	fmt.Scan(&weight)
	fmt.Print("Напишите вес штанги в жиме: ")
	fmt.Scan(&benchP)
	fmt.Print("Напишите количество повторений в подходе: ")
	fmt.Scan(&benchReps)
	benchPress = (weight * benchReps * 0.0333) + benchP
	fmt.Print("Напишите вес штанги в становой: ")
	fmt.Scan(&deadl)
	fmt.Print("Напишите количество повторений в подходе: ")
	fmt.Scan(&deadliftReps)
	deadlift = deadl / (1.0278 - 0.0278*deadliftReps)
	fmt.Print("Напишите вес штанги в приседаниях: ")
	fmt.Scan(&squatt)
	fmt.Print("Напишите количество повторений в подходе: ")
	fmt.Scan(&squatReps)
	squat = squatt / (1.0278 - 0.0278*squatReps)

	powerliftingCheck(weight, benchPress, deadlift, squat)
}

func powerliftingCheck(weight, benchPress, deadlift, squat float32) {
	if weight >= 53 && weight <= 59 {
		if benchPress+deadlift+squat >= 260 && benchPress+deadlift+squat <= 282.5 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 282.5 && benchPress+deadlift+squat <= 325 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 325 && benchPress+deadlift+squat <= 410 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 410 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		}
	} else if weight >= 59 && weight <= 66 {
		if benchPress+deadlift+squat >= 290 && benchPress+deadlift+squat <= 315 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 315 && benchPress+deadlift+squat <= 362.5 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 362.5 && benchPress+deadlift+squat <= 455 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 455 && benchPress+deadlift+squat <= 540 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 540 && benchPress+deadlift+squat <= 635 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 635 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK разряда")
		}
	} else if weight >= 66 && weight <= 74 {
		if benchPress+deadlift+squat >= 320 && benchPress+deadlift+squat <= 350 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 350 && benchPress+deadlift+squat <= 402.5 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 402.5 && benchPress+deadlift+squat <= 510 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 510 && benchPress+deadlift+squat <= 595 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 595 && benchPress+deadlift+squat <= 720 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 720 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK")
		}
	} else if weight >= 74 && weight <= 83 {
		if benchPress+deadlift+squat >= 352.5 && benchPress+deadlift+squat <= 385 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 385 && benchPress+deadlift+squat <= 440 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 440 && benchPress+deadlift+squat <= 537.5 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 537.5 && benchPress+deadlift+squat <= 675 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 675 && benchPress+deadlift+squat <= 785 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 785 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK разряда")
		}
	} else if weight >= 83 && weight <= 93 {
		if benchPress+deadlift+squat >= 387.5 && benchPress+deadlift+squat <= 422.5 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 422.5 && benchPress+deadlift+squat <= 482.5 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 482.5 && benchPress+deadlift+squat <= 582.5 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 582.5 && benchPress+deadlift+squat <= 775 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 775 && benchPress+deadlift+squat <= 850 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 850 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK разряда")
		}
	} else if weight >= 93 && weight <= 110 {
		if benchPress+deadlift+squat >= 412.5 && benchPress+deadlift+squat <= 465 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 465 && benchPress+deadlift+squat <= 520 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 520 && benchPress+deadlift+squat <= 610 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 610 && benchPress+deadlift+squat <= 800 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 800 && benchPress+deadlift+squat <= 925 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 925 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK разряда")
		}
	} else if weight >= 110 && weight <= 120 {
		if benchPress+deadlift+squat >= 460 && benchPress+deadlift+squat <= 500 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 500 && benchPress+deadlift+squat <= 552.5 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 552.5 && benchPress+deadlift+squat <= 645 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 645 && benchPress+deadlift+squat <= 840 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 840 && benchPress+deadlift+squat <= 970 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 970 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK разряда")
		}
	} else if weight >= 120 && weight <= 121 {
		if benchPress+deadlift+squat >= 497.5 && benchPress+deadlift+squat <= 530 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 530 && benchPress+deadlift+squat <= 600 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 600 && benchPress+deadlift+squat <= 687.5 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 687.5 && benchPress+deadlift+squat <= 875 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 875 && benchPress+deadlift+squat <= 1005 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 1005 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK разряда")
		} else {
			fmt.Println("Вы не соответствуете требованиям ни для одного разряда")
		}
	} else if weight >= 93 {
		if benchPress+deadlift+squat >= 510 && benchPress+deadlift+squat <= 545 {
			fmt.Println("Вы соответствуете требованиям для получения 3 разряда")
		} else if benchPress+deadlift+squat >= 545 && benchPress+deadlift+squat <= 617.5 {
			fmt.Println("Вы соответствуете требованиям для получения 2 разряда")
		} else if benchPress+deadlift+squat >= 617.5 && benchPress+deadlift+squat <= 735 {
			fmt.Println("Вы соответствуете требованиям для получения 1 разряда")
		} else if benchPress+deadlift+squat >= 735 && benchPress+deadlift+squat <= 890 {
			fmt.Println("Вы соответствуете требованиям для получения KMC разряда")
		} else if benchPress+deadlift+squat >= 890 && benchPress+deadlift+squat <= 1035 {
			fmt.Println("Вы соответствуете требованиям для получения MC разряда")
		} else if benchPress+deadlift+squat >= 1035 {
			fmt.Println("Вы соответствуете требованиям для получения MCMK разряда")
		} else {
			fmt.Println("Вы не соответствуете требованиям ни для одного разряда")
		}
	} else {
		fmt.Println("Вы не соответствуете требованиям ни для одного разряда")
	}
}
