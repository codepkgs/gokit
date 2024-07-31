package estrings

func GenLowerChars() string {
	var lowerChars = make([]rune, 0, 26)
	for i := '\x61'; i < '\x61'+26; i++ {
		lowerChars = append(lowerChars, i)
	}
	return string(lowerChars)
}

func GenUpperChars() string {
	var upperChars = make([]rune, 0, 26)
	for i := '\x41'; i < '\x41'+26; i++ {
		upperChars = append(upperChars, i)
	}
	return string(upperChars)
}

func GenNumbers() string {
	var numbers = make([]rune, 0, 10)
	for i := '\x30'; i < '\x30'+10; i++ {
		numbers = append(numbers, i)
	}
	return string(numbers)
}

func GenSymbols() string {
	var symbols = make([]rune, 0, 32)
	for i := '\x21'; i < '\x21'+32; i++ {
		if i >= '\x30' && i <= '\x39' {
			continue
		}
		symbols = append(symbols, i)
	}
	return string(symbols)
}

func GenAlphaChars() string {
	return GenLowerChars() + GenUpperChars()
}

func GenAlphaNumChars() string {
	return GenAlphaChars() + GenNumbers()
}
