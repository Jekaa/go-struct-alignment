# go-struct-alignment

Реализация проекта с объяснением оптимизации структур в Go

Пример вывода программы:
```
Размеры структур:
BadStruct:   24 байт
GoodStruct:  16 байт

Использование памяти для мап:
map[int]bool:     3788792 байт
map[int]struct{}: 2898872 байт

Горутина получила сигнал
```

---

### Объяснение оптимизаций

**1. Выравнивание полей в структурах**

В Go компилятор автоматически добавляет padding (отступы) между полями структуры для выравнивания данных в памяти. Это требуется для эффективного доступа к памяти процессором.

Пример плохой структуры:
```go
type BadStruct struct {
	a bool   // 1 байт
	b int64  // 8 байт
	c bool   // 1 байт
}
```
Размер структуры:
- 1 (bool) + 7 (padding) + 8 (int64) + 1 (bool) + 7 (padding) = 24 байта

Оптимизированная версия:
```go
type GoodStruct struct {
	b int64  // 8 байт
	a bool   // 1 байт
	c bool   // 1 байт
}
```
Размер структуры:
- 8 (int64) + 1 (bool) + 1 (bool) + 6 (padding) = 16 байт

**Оптимизация:**
- Размещаем поля по убыванию размера
- Группируем маленькие поля вместе
- Уменьшаем общий размер структуры на 33%

**2. Использование пустой структуры struct{}**

Пустая структура `struct{}`:
- Не занимает места в памяти (0 байт)
- Используется как заглушка, когда нужно только наличие значения

**Пример с map:**
```go
// Плохой вариант: 1 байт на значение
m1 := make(map[int]bool) 

// Хороший вариант: 0 байт на значение
m2 := make(map[int]struct{})
```
При 1 млн элементов:
- map[int]bool: ~3.7 MB
- map[int]struct{}: ~2.8 MB

**Другие применения:**
- Сигналы в каналах
- Реализация множеств
- Заглушки для методов

**3. Преимущества оптимизаций**
- Уменьшение потребления памяти
- Более эффективное использование кэша процессора
- Уменьшение нагрузки на сборщик мусора
- Повышение производительности при работе с большими объемами данных

Эти оптимизации особенно важны для:
- Высоконагруженных систем
- Приложений с большими объемами данных
- Систем с ограниченными ресурсами
