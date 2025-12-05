package main

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func main() {
	// Тест 1: Простая кириллица
	test1 := "ПРИВЕТ"
	fmt.Printf("ToMorse(%q) = %q\n", test1, morse.ToMorse(test1))

	// Тест 2: Проблемный текст
	test2 := "ЩИНДЧТМФЙВФНИХЦХЫУЭБЙЫЩФЦЙАФТ"
	fmt.Printf("ToMorse(%q) = %q\n", test2, morse.ToMorse(test2))

	// Тест 3: Первые несколько символов
	for _, ch := range test2[:5] {
		// Попробуем через RuneToMorse если она экспортирована
		fmt.Printf("Char %q\n", ch)
	}
}
