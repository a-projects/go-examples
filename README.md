# go-examples
Examples in Go like duck typing, pointer on interface, channels, generators

Пример утиной типизации в Go, указателя на интерфейс, работа каналов и мьютексов, реализация генератора

## Утиная типизация
Известно, что Go использует «утиную типизацию». На самом деле это «структурная типизация» потому, что происходит во время компиляции, тогда как «утиная типизация» обычно происходит во время выполнения. И работает это так: если у структуры есть набор методов, соответствующих интерфейсу, вы можете использовать данную структуру везде, где нужен этот интерфейс, без явного определения.

В примере рассмотрен как сам механизм «структурной типизации», так и детали применения.

[ducktype-in-action](https://github.com/a-projects/go-examples/tree/main/ducktype-in-action)

## Комозиция, не наследование
В мире ООП объекты состоят из более мелких объектов, это называется комозицей. В Go используют композицию как через встраивание, так и явную. Ниже в примере рассмотрена явная композиция от «банды четырёх»

[composition-basic](https://github.com/a-projects/go-examples/tree/main/composition-basic)

## Встраивание и затенение
Встраивание (embedding) в Go - это композиция с дополнительным преимуществом "переадресации" (forwarding) интерфейса. Если проще, то структуры могут встраиваться в друг друга, причём встраиваются как поля так и методы. Методы встроенной структуры могут переопределяться (затеняться).
Надо заметить, что встраивание не является наследованием, Вы нес сможете присвоить переменную типа Manager переменной типа Employee, хотя
Employee будет встроенным в Manager. Также встроенная структура не сможет воспользоваться методами внешнего структуры. Для функционала полиморфизма и динамической диспетчеризации нужно использовать интерфейсы

[forwarding-and-shadowing](https://github.com/a-projects/go-examples/tree/main/forwarding-and-shadowing)

## Тип интерфейса в качестве значения и указателя
В Go интерфейсы определяются как набор методов. Если тип реализует все методы интерфейса, то говорят, что он удовлетворяет интерфейсу. Это означает, что тип может использоваться везде, где ожидается интерфейс. Это и есть утиная типизация, в данном примере рассмотрим функции, которые принимают тип интерфейса в качестве значения, а также которые принимают интерфейс в качестве указателя.

[interface-value-pointer](https://github.com/a-projects/go-examples/tree/main/interface-value-pointer)

## Указатель на интерфейс
Данный подход используется крайне редко и по сути это «указатель на указатель». Это используется в рефликсии, чтобы понимать какой интерфейс, а не структура был передан, по сути это маскировка типа, что и продемонстрировано в примере.

[pointer-on-interface](https://github.com/a-projects/go-examples/tree/main/pointer-on-interface)

## Горутины и каналы
Горутины или go-подпрограммы являются ключевой концепцией в модели конкурентности Go. По сути в рантайме выполнения Go есть планировщик, который принимает решение о выполнении нескольких подпрограмм в рамках одного потока или нескольких потоков ОС. Ключевое же предназначение каналов - это организация общения между горутинами. В примере продемонстрировано использование горутин с каналами.

#### Запуск горутин с ожиданием окончания их выполнения

[goroutine-with-wait](https://github.com/a-projects/go-examples/tree/main/goroutine-with-wait)

#### Передача данных между горутинами с помощью каналов

[goroutine-with-channel](https://github.com/a-projects/go-examples/tree/main/goroutine-with-channel)

#### Передача каналов в виде возвращающего значения функции

[function-return-channel](https://github.com/a-projects/go-examples/tree/main/function-return-channel)

#### Прекращение выполнения горутины с помощью функции отмены

[goroutine-with-cancel](https://github.com/a-projects/go-examples/tree/main/goroutine-with-cancel)

## Конкурентность с мьютексами
В Go при работе с данными предпочтительно использовать каналы. Так легче понять как происходит движение данных в программе. Но иногда каналы усложняют код, например когда дело касается состояния экземпляра. В таком случае и пригодяться мьютексы.

[mutex-in-action](https://github.com/a-projects/go-examples/tree/main/mutex-in-action)
