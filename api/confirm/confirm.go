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
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02 15:04:05"`
	CheckOut time.Time `form:"check_out" binding:"required" time_format:"2006-01-02"`
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
			return false
		}
	}
	return true
}

//接收时间
//http://localhost:8000/api/bind/bookable?check_in=2020-03-01 11:12:12&check_out=2020-03-02
func GetBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindQuery(&b); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "succeed"})
		fmt.Println(b)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
