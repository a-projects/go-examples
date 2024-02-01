package main

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

type Data struct {
	Desc string
}

type PrimarySrv struct {
	Id   int
	Data *Data
}

type SecondarySrv struct {
	Id   int
	Data *Data
}

type Pool struct {
	PrimarySrv   *PrimarySrv
	SecondarySrv *SecondarySrv
}

func main() {
	// Поставим задачу: с помощью рефлексии добавить данные для всех не пустых Srv
	// Создадим экземпляр, причём PrimarySrv будет nil и данные необходимо добавить только в Secondary
	pool := &Pool{
		PrimarySrv: nil,
		SecondarySrv: &SecondarySrv{
			Data: nil,
		},
	}

	// Перебрать все поля в Pool
	value := reflect.ValueOf(pool)
	for i := 0; i < reflect.Indirect(value).NumField(); i++ {
		srvField := reflect.Indirect(value).Field(i)
		// Выбрать поля являющие указателем, но не nil
		if srvField.Kind() == reflect.Ptr && srvField.IsValid() && !srvField.IsNil() {
			typeField := reflect.Indirect(value).Type().Field(i)
			fmt.Println(typeField)

			// Получить экземляр Srv
			iSrv := srvField.Interface()
			value := reflect.ValueOf(iSrv)

			// Поиск можно осуществлять и перебором полей,
			// но в данном примере ипользуется встроенный FieldByName

			// Найти поле Id
			idField := value.Elem().FieldByName("Id")
			if idField.IsValid() {
				// Присвоить полю Id значение
				idField.SetInt(2)
			}

			// Найти поле Data
			dataField := value.Elem().FieldByName("Data")
			if dataField.IsValid() {
				data := &Data{
					Desc: "Это вспомогательный сервер",
				}
				// Присвоить полю Data значение указателя
				dataField.Set(reflect.ValueOf(data))
			}
		}
	}

	spew.Dump(pool)

	// результат:
	// {SecondarySrv  *main.SecondarySrv  8 [1] false}
	// (*main.Pool)(0xc00008a330)({
	//  PrimarySrv: (*main.PrimarySrv)(<nil>),
	//  SecondarySrv: (*main.SecondarySrv)(0xc00008a340)({
	//   Id: (int) 2,
	//   Data: (*main.Data)(0xc00008a350)({
	//    Desc: (string) (len=50) "Это вспомогательный сервер"
	//   })
	//  })
	// })
}
