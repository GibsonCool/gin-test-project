package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*

 */

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

func main() {
	en := en2.New()
	zh := zh2.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()

	r := gin.Default()

	/*
		curl -X GET "http://127.0.0.1:8080/testing"

		curl -X GET "http://127.0.0.1:8080/testing?age=10&name=ggtest&address=beijing"

		curl -X GET "http://127.0.0.1:8080/testing?age=10&name=ggtest&address=beijing&locale=en"
	*/
	r.GET("/testing", handlerTesting)
	r.POST("/testing", handlerTesting)
	if err := r.Run(); err != nil {
		log.Println(err.Error())
	}
}

func handlerTesting(c *gin.Context) {
	// 取出参数中的local 的值,可以将这部分抽离到中间件中
	locale := c.DefaultQuery("locale", "zh")
	trans, _ := Uni.GetTranslator(locale)
	switch locale {
	case "zh":
		_ = zh_translations.RegisterDefaultTranslations(Validate, trans)
		break
	case "en":
		_ = en_translations.RegisterDefaultTranslations(Validate, trans)
		break
	case "zh_tw":
		_ = zh_tw_translations.RegisterDefaultTranslations(Validate, trans)
		break
	default:
		_ = zh_translations.RegisterDefaultTranslations(Validate, trans)
		break
	}

	/*
		自定义错误内容，覆盖原内容，其实各种语言翻译也是调用这个方法来实现的翻译
	*/
	Validate.RegisterTranslation("gt", trans,
		func(ut ut.Translator) error {
			return ut.Add("gt", "自定义翻译错误内容：{0} 年龄字段必须大于 {1}", true) // see universal-translator for details
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			var digits uint64
			if idx := strings.Index(fe.Param(), "."); idx != -1 {
				digits = uint64(len(fe.Param()[idx+1:]))
			}

			f64, _ := strconv.ParseFloat(fe.Param(), 64)
			t, _ := ut.T("gt", fe.Field(), ut.FmtNumber(f64, digits))
			return t
		})

	var person Person
	// 直接绑定解析 数据
	_ = c.ShouldBind(&person)
	// 使用我们上面通过 local 参数获取到的翻译验证去翻译错误
	e := Validate.Struct(person)
	sliceErrs := []string{}
	if e != nil {
		errors := e.(validator.ValidationErrors)
		for _, err := range errors {
			sliceErrs = append(sliceErrs, err.Translate(trans))
		}
		c.String(http.StatusOK, "person bind err:%#v", sliceErrs)
	} else {
		c.String(http.StatusOK, "%v", person)
	}

}
