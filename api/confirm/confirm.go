package confirm

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

// Booking contains binded and validated data.
//自己去查一下gin自带的tag验证规则
//binding:"max=10,min=1"
//binding:"len=1" //大于1，多个条件使用逗号或者|隔开，逗号表示并且
//gtfield=xxx 大于其他值
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

//加载到内存,需要验证的时候,再从内存获取
func BookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		//年份比较  今年至今多少天比较
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			fmt.Println("xxxx")
			return false
		}
	}
	fmt.Println("xxxx1")
	return true
}

//接收时间
//curl -X GET "http://localhost:8000/api/bind/bookable?check_in=2020-03-01&check_out=2020-03-02&age=1"
//使用curl不能解析check_in参数中间的空格，后面想想办法，先这样
//json格式:
//curl -H "Content-Type:application/json" -X GET "http://localhost:8000/api/bind/bookable" -d '{"check_in":"2020-03-01","check_out":"2020-03-02"}'
func GetBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindQuery(&b); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "succeed"})
		fmt.Println(b)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
