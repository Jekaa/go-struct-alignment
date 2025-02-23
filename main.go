package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// Неоптимизированная структура с плохим выравниванием
type BadStruct struct {
	a bool  // 1 байт
	b int64 // 8 байт
	c bool  // 1 байт
}

// Оптимизированная структура с правильным выравниванием
type GoodStruct struct {
	b int64 // 8 байт
	a bool  // 1 байт
	c bool  // 1 байт
}

func measureMemoryUsage(initFunc func()) uint64 {
	runtime.GC() // Очистка памяти перед измерением
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	before := memStats.HeapAlloc

	initFunc() // Выполнение функции с тестовыми данными

	runtime.ReadMemStats(&memStats)
	return memStats.HeapAlloc - before
}

func main() {
	bad := BadStruct{}
	good := GoodStruct{}

	fmt.Println("Размеры структур:")
	fmt.Printf("BadStruct:  %3d байт\n", unsafe.Sizeof(bad))
	fmt.Printf("GoodStruct: %3d байт\n\n", unsafe.Sizeof(good))

	const elements = 100_000 // Количество элементов для теста мап

	// Тестирование памяти для map с bool
	boolMapMemory := measureMemoryUsage(func() {
		m := make(map[int]bool, elements)
		for i := 0; i < elements; i++ {
			m[i] = true
		}
	})

	// Тестирование памяти для map с struct{}
	structMapMemory := measureMemoryUsage(func() {
		m := make(map[int]struct{}, elements)
		for i := 0; i < elements; i++ {
			m[i] = struct{}{}
		}
	})

	fmt.Println("Использование памяти для мап:")
	fmt.Printf("map[int]bool:    %6d байт\n", boolMapMemory)
	fmt.Printf("map[int]struct{}: %6d байт\n\n", structMapMemory)

	// Пример использования struct{} в канале
	signal := make(chan struct{})
	go func() {
		fmt.Println("Горутина получила сигнал")
		signal <- struct{}{}
	}()
	<-signal
	close(signal)
}
