package task

// Description:
// Например,  число 2 записывается как II в римской транслитерации, просто путем сложения двух единиц.
// 12Число записывается как  XII, что просто означает X + II. Число 27 записывается как XXVII, что означает XX + V + II.

// Римские цифры обычно записываются от наибольшей к наименьшей слева направо. Однако цифра четыре записывается не так IIII.
// Вместо этого число четыре записывается как IV. Поскольку единица стоит перед пятеркой, мы вычитаем ее, получая четыре.
// Тот же принцип применим к числу девять, которое записывается как IX. Существует шесть случаев, когда используется вычитание:
// I может быть размещен перед V(5) и X(10), чтобы получить 4 и 9.
// X может быть размещен перед L(50) и C(100), чтобы получить 40 и 90.
// C может быть размещен перед D(500) и M(1000), чтобы получить 400 и 900.

func RomanToInt(s string) int {
	var res int
	charInt := map[string]int{
		"I":        1,
		"V":        5,
		"X":        10,
		"L":        50,
		"C":        100,
		"D":        500,
		"M":        1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	lenght := len(s)
	i:=0
	for i < lenght{
		var temp string
		num := string(s[i])
		if "I" == num || "X" == num || "C" == num{
			if i + 1 < lenght{
				temp = num + string(s[i+1])
				_, ok := charInt[temp]
				if ok{
					num = temp
					i++
				}
			}
		}
		res += charInt[num]
		i++
	}
	
	return res
}

func RomanToInt2(s string) int{
	values := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
	result := 0
    n := len(s)
	for i := 0; i < n; i++{
		current := values[s[i]]
		if i+1 < n && current < values[s[i+1]]{
			result -= current
		} else{
			result += current
		}
	}

	return result
}