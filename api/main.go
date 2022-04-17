package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))

	fmt.Println("helllllllo")
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!222222")
}

// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// )

// func add(x int, y int) int {
// 	return x + y
// }
// func add2(x, y int) int {
// 	return x + y
// }
// func swap(s, y string) (string, string) {
// 	return s, y
// }

// func main() {
// 	var q bool
// 	var w string
// 	var e int
// 	var x, y int = 3, 4
// 	var z float64 = math.Sqrt(float64(x*x + y*y))
// 	fmt.Println(rand.Intn(100))         //擬似乱数
// 	fmt.Println(math.Sqrt(5))           //平方根
// 	fmt.Println(math.Pi)                //円周率
// 	fmt.Println(add(42, 13))            //関数
// 	fmt.Println(add2(15, 15))           //関数,省略
// 	fmt.Println(swap("hello", "world")) //関数、２値返し
// 	fmt.Println(q)                      // false
// 	fmt.Println(w)                      // ""
// 	fmt.Println(e)                      // 0
// 	fmt.Println(z)                      // 0

// }
