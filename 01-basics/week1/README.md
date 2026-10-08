# Неделя 1. Основы на практике

Темп: 1–1,5 часа в день, 7 дней. Каждый день: немного теории по ссылкам, затем
пакет `dayN/`, где тесты уже написаны, а вы пишете решение вместо `// TODO`.

День закрыт, когда:

1. `go test ./01-basics/week1/dayN/` проходит;
2. `gofmt -l .` ничего не выводит, `go vet ./...` молчит;
3. вы можете ответить на вопросы дня вслух, не подглядывая;
4. изменения закоммичены и запушены.

Если застряли больше чем на 20 минут, смотрите вывод `go test -v`, пишите `fmt.Println`
прямо в тесте или спрашивайте в проекте. Подглядывать в готовые решения в интернете не стоит.

---

## День 1. Инструменты и первый тест

**Теория:** [How to Write Go Code](https://go.dev/doc/code), [Add a test](https://go.dev/doc/tutorial/add-a-test).
Разберитесь с `go mod init`, `go run`, `go build`, `go test`, `gofmt`, `go vet`. Настройте редактор с gopls
(VS Code с расширением Go или GoLand), включите форматирование при сохранении.

**Задание** ([`day1/`](day1/)): склонируйте репозиторий, запустите `go test ./...` и убедитесь,
что все тесты красные. Затем реализуйте `Greet(name string) string`.

**Готово, когда:**
- тесты `day1` зелёные;
- `go test -cover ./01-basics/week1/day1/` показывает 100%;
- можете ответить: что такое модуль и пакет; почему `Greet` с большой буквы; чем `go run` отличается от `go build`.

## День 2. Переменные, константы, iota, типы

**Теория:** [Effective Go: Constants](https://go.dev/doc/effective_go#constants),
[Go by Example: Constants](https://gobyexample.com/constants), нулевые значения и явное приведение типов.

**Задание** ([`day2/`](day2/)): методы `Weekday.String` и `IsWeekend`; константы `KB`, `MB`, `GB`
через `iota` и сдвиг `<<`; `FormatBytes`; `CelsiusToFahrenheit`.

**Готово, когда:**
- тесты `day2` зелёные;
- можете ответить: какие нулевые значения у `int`, `string`, `bool`, указателя; почему `float64(n) / 1024`
  и `n / 1024` дают разное; зачем `fmt.Println(Monday)` печатает «понедельник».

## День 3. Управление потоком: for, switch, defer

**Теория:** [Effective Go: Control structures](https://go.dev/doc/effective_go#control-structures),
[Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover), `range` по числу (Go 1.22+).

**Задание** ([`day3/`](day3/)): `FizzBuzz`, `CollatzSteps` (с ошибкой `ErrNotPositive`),
`Grade` через `switch` без выражения, `DeferOrder` (только через `defer`).

**Готово, когда:**
- тесты `day3` зелёные;
- можете ответить: в каком порядке выполняются `defer`; когда вычисляются аргументы отложенного вызова;
  как `defer` может поменять именованный результат функции.

## День 4. Функции и замыкания

**Теория:** [Go by Example: Variadic Functions](https://gobyexample.com/variadic-functions),
[Closures](https://gobyexample.com/closures), [Multiple Return Values](https://gobyexample.com/multiple-return-values).

**Задание** ([`day4/`](day4/)): `Sum`, `MinMax`, `Map`, `Filter`, `NewCounter`, `Compose`.

**Готово, когда:**
- тесты `day4` зелёные;
- можете ответить: что такое замыкание и где живёт переменная счётчика; как передать слайс в
  вариадическую функцию; зачем возвращать `ok bool`, а не, скажем, `-1`.

## День 5. Массивы и слайсы

**Теория:** [Go Slices: usage and internals](https://go.dev/blog/slices-intro),
[Arrays, slices (and strings): The mechanics of append](https://go.dev/blog/slices), пакет [`slices`](https://pkg.go.dev/slices).

**Задание** ([`day5/`](day5/)): `Reverse` на месте, `Unique` без изменения входа, `Chunk`
(тест `TestChunkIndependent` ловит ловушку общего базового массива), `RemoveAt`.

**Готово, когда:**
- тесты `day5` зелёные;
- можете без запуска предсказать вывод и объяснить его:

```go
a := []int{1, 2, 3, 4}
b := a[:2]
b = append(b, 99)
fmt.Println(a, len(b), cap(b))
```

## День 6. Map, строки, руны

**Теория:** [Go maps in action](https://go.dev/blog/maps), [Strings, bytes, runes and characters](https://go.dev/blog/strings),
пакеты [`strings`](https://pkg.go.dev/strings), [`unicode`](https://pkg.go.dev/unicode), [`maps`](https://pkg.go.dev/maps).

**Задание** ([`day6/`](day6/)): `RuneLen`, `ReverseString`, `IsPalindrome`, `WordCount`, `TopWords`.
`WordCount` и `TopWords` — заготовка для CLI `wordfreq` на следующей неделе.

**Готово, когда:**
- тесты `day6` зелёные;
- можете ответить: почему `len("Привет")` равно 12; что даёт `for i, r := range s` по сравнению с `s[i]`;
  почему порядок обхода map каждый раз разный и как получить стабильный.

## День 7. Указатели и итог недели

**Теория:** [Tour: Pointers](https://go.dev/tour/moretypes/1), [Tour: Pointer receivers](https://go.dev/tour/methods/4),
[FAQ: Should I define methods on values or pointers?](https://go.dev/doc/faq#methods_on_values_or_pointers).

**Задание** ([`day7/`](day7/)): `Swap`, `ResetAge`, `NewAccount` и методы `Deposit`, `Withdraw`,
функция `Transfer`, которая при ошибке не меняет ни один баланс.
Сравните `ResetAgeCopy` (уже написана) с вашей `ResetAge`.

**Готово, когда:**
- тесты `day7` зелёные, и `go test -race ./...` по всей неделе зелёный;
- в Summary последнего запуска CI на GitHub все 7 пакетов отмечены как решённые;
- можете ответить: что будет при вызове метода с указателем-получателем на `nil`; почему `Balance` в копейках
  и `int64`, а не `float64`.
