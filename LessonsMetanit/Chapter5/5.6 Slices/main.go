package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println(" СРЕЗЫ ")
	// Срезы представляют собой динамическую структуру данных, в которые можно
	// добавлять и удалять элементы
	// Slice определяется также как массив, за тем исключением, что у него не определяется
	// количество элементов
	// общий синтаксис определения
	//
	// var name_variable []type

	//var users []string

	// срез также можно инициализировать значениями
	var users []string = []string{"Tom", "Alice", "Kate"}
	//Или так
	users2 := []string{"Tom", "Alice", "Kate"}
	fmt.Println(users)
	fmt.Println(users2)

	// Оператор среза и создание среза из последовательностей
	// Аналогичен операторам индексов и диапозонов в C#
	// List<int> a = new(){1,2,3,4,5,}
	// var b = a[0..2]
	// только в общем случае пишется в через ":"
	// s := [5]string{"Tom", "Sam", "Ivan", "Anatol", "SomePerson"}
	// i, j := 1, 4
	// s1 := s[i:j] // Срез s1 будет содержать элементы массива s, с первого по третий,
	// то есть с i по j-1
	// то есть видим, что последний элемент не включается в срез
	s := [5]string{"Tom", "Sam", "Ivan", "Anatol", "SomePerson"}
	i, j := 1, 4
	s1 := s[i:j]
	// Оператор среза создаёт из последовательности или другого среза новый срез
	fmt.Println(s1)

	// Необязательно указывать начальный(i) и конечный(j) индекс среза,
	// если не указан начальный индекс, то он поумолчанию будет начинаться с 0
	// если не указан конечный индекс, то он поумолчанию будет начинаться с len(slice)
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users_1 := initialUsers[2:6] // С 3-его по 6-ой
	users_2 := initialUsers[:4]  // С 1-его по 4-ой
	users_3 := initialUsers[3:]  // С 2-его по 8-ой
	fmt.Println(users_1)
	fmt.Println(users_2)
	fmt.Println(users_3)

	//Создание среза из строки
	str := "Hello world"
	slice := str[6:]
	fmt.Println(slice)

	//Создание среза из другого среза
	var array [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	slice1 := array[1:6]
	fmt.Println("slice from array:", slice1) // slice from array: [2 3 4 5 6]

	slice2 := slice1[:3]
	fmt.Println("slice from array:", slice2) // slice from array: [2 3 4]

	//Обращение к элементам среза и перебор среза
	var users3 []string = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users3[2]) // Kate
	users3[2] = "Katherine"

	for _, v := range users3 {
		fmt.Println(v)
	}

	//ВНУТРЕННЯЯ РЕАЛИЗАЦИЯ СРЕЗА
	//Обычно срез состоит из указателя на сегмент массива, длины сегмента массива и максимума,
	//до которого может быть расширен срез:

	// Указатель
	// В реальности срезы сами не хранят данные, они ссылаются на данные откуда были срезаны,
	// то есть на базовую последовательность
	// Это говорит о том, что они ведут себя как ссылочные типы в C#
	// стоит изменить данные, на которые ссылаются slice, то всё
	// данные, которые он выдаёт изменятся
	// Например:
	array1 := [6]int{1, 2, 3, 4, 5, 6} // Начальный массив
	fmt.Println("Начальный массив: ", array1)

	slice_to_array1 := array1[1:5]                   //[2, 3, 4, 5]
	fmt.Println("Начальный срез: ", slice_to_array1) //[2, 3, 4, 5]

	//Изменяем второй элемент массива
	array1[1] = 3128947
	fmt.Println("Что в массиве array1: ", array1, "И что в срезе slice_to_array1: ", slice_to_array1)

	//Изменение среза затрагивает источник данных
	slice_to_array1[1] = -9000
	fmt.Println("Что в массиве array1: ", array1, "И что в срезе slice_to_array1: ", slice_to_array1)

	//Длина - len(slice) // возвращает количество элементов среза, которые указывают на какие-либо данные
	fmt.Println("len(slice_to_array1) = ", len(slice_to_array1))

	// Также у срезов имеется свойство capacity - емкость, оно возвращает общее количество
	// указателей в срезе, емкость связана непосредственно со срезом,
	// в то время как длина с массивом среза
	// чтобы узнать емкость среза применяется функция
	// cap(slice)

	slice3 := array1[1:4]
	fmt.Println("cap(slice3) =", cap(slice3), "len(slice3) =", cap(slice3))

	// СОЗДАНИЕ СРЕЗА С ПОМОЩЬЮ ФУНКЦИИ MAKE
	// Чтобы создать чисто срез не из последовательности
	// используется функция make()
	// имя_среза := make([] тип_элементов_среза, длина_среза, емкость_среза)
	var users4 []string = make([]string, 3, 16) // длина среза 3
	users4[0] = "Sam"
	users4[1] = "Thomas"
	users4[2] = "Bob"

	fmt.Println(users4)

	//Двумерные срезы

	//Срез, содержащий другие срезы
	slice2D := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{5, 6},
	}

	fmt.Println(slice2D)
	//Срез, содержащий массив int из двух элементов
	slice_array2D := [][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(slice_array2D)

	// Так как срез динамическая структура, то туда можно добавить элемент
	// Это осуществляется с помощью функции
	// append(slice, adding_element),
	// которая возвращает срез с добавленным элементом
	users5 := []string{"Tom", "Alice", "Kate"}
	fmt.Println(users5)
	users5 = append(users5, "Bob") // Добавление происходит в конец среза

	//Удаление элемента также происходит с помощью append()

	//Хотим удалить четвертый элемент
	users5 = []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	n := 3
	fmt.Println(users5)

	// То есть берем данные среза до третьего элемента не включительно,
	// далее в него добавляем все элементы после третьего элемента
	users5 = append(users5[:n], users5[n+1:]...)
	fmt.Println(users5)

	//КОПИРОВАНИЕ СРЕЗА
	// Копирование среза происходит с помощью функции copy()
	// func copy(destination, source []T) int
	// Где destination - срез в который копируется данные из среза source
	// source - срез из которого копируем элементы
	// Возвращает число скопированных элементов
	slice_1 := []int{1, 2, 3, 4, 5, 6}
	slice_2 := []int{}
	fmt.Println(copy(slice_2, slice_1))
	fmt.Println(slice_2)
	//0, так как скопировано 0 элементов
	// так как slice_2 ноль элементов, то в него ничего не будет скопировано
	// так как там нет места для копирования

	slice_2 = make([]int, 3)
	// здесь длин slice_2 теперь 3, поэтому скопировано 3 элемента
	fmt.Println(copy(slice_2, slice_1)) //3
	fmt.Println(slice_2)

	//Ещё одна ситуация, при которой количество элементов в
	//срезе достижения больше, чем в срезе из которого копируем
	// элементы не были заменины копией, будут иметь значение,
	// которое до этого хранили, или значение поумолчанию
	slice_2 = make([]int, 9)
	fmt.Println(copy(slice_2, slice_1)) //6
	fmt.Println(slice_2)

	//Для сортировки элементов в срезе имеется пакет sort
	// Для среза int-ов sort.Ints(ints)
	// Для строк sort.Strings(strings)
	// Для float64 sort.Float64s(float32s)
	// Для float32 sort.Float32s(float32s)

	users6 := []string{"Tom", "Bob", "Sam"}
	sort.Strings(users6)
	fmt.Println(users6)

	// ПОИСК
	// Чтобы найти индекс определенного элемента в срезе используются функции
	// Для этого пакет sort предоставляет
	// соответственно функции sort.SearchInts(),
	// sort.SearchFloat64s() и sort.SearchStrings.
	// Эти функции выполняют двоичный поиск в отсортированном срезе в порядке
	// возрастания.

	// Если элемент найден в массиве, то возвратиться индекс, в котором
	// он стоит, если не найден, то индекс, по которому его нужно поместить в срез
	// чтобы его можно было найти бинарным поиском
	intSlice := []int{11, 22, 33, 44, 55, 66}
	//Ищем число 33
	fmt.Println(sort.SearchInts(intSlice, 33))
	//Ищем число 45
	fmt.Println(sort.SearchInts(intSlice, 33))

	stringSlice := []string{"Tom", "Bob", "Sam"}
	//Ищем строку Sam
	fmt.Println(sort.SearchStrings(stringSlice, "Sam")) // 2, после автосортировки
	//Ищем строку Alice
	fmt.Println(sort.SearchStrings(stringSlice, "Sam")) // 0

	// Можно реверсировать порядок элементов в срезе
	// Только предварительно нужно обязательно его обернуть в
	// соответствующий тип
	// IntSlice/Float64Slice/StringSlice с помощью соответственно функций
	// sort.IntSlice/sort.Float64Slice/sort.StringSlice
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice) // [66 55 44 33 22 11]

	sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
	fmt.Println(stringSlice)

	// Сравнение двух срезов происходит с помощью функции
	// reflect.DeepEqual(slice1, slice2)
	// Два среза равны если:
	// 1) если они хранят данные одного типа
	// 2) содержат одни и те же элементы

	slice4 := []int{1, 2, 3, 4}
	slice5 := []string{"Tom", "Bob", "Sam"}
	slice6 := []int{1, 2, 3}
	slice7 := []int{1, 2, 3, 4}

	fmt.Println("slice4 == slice5", reflect.DeepEqual(slice4, slice5)) // false
	fmt.Println("slice4 == slice6", reflect.DeepEqual(slice4, slice6)) // false
	fmt.Println("slice4 == slice7", reflect.DeepEqual(slice4, slice7)) // false

}
