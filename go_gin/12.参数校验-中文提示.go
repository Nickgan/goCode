package main

import (
	"fmt"
	"goCode/go_gin/resp"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/*
*
* 参考文档： https://www.fengfengzhidao.com/special/3/41
 */

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
}

func main() {

	gin.SetMode(gin.ReleaseMode) //精简日志输出
	r := gin.Default()

	// 4：json参数
	r.POST("/user/add/json/validata", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required,min=3,max=5"`
			Age  int    `json:"age"`
		}

		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println("error============>", err.Error())
			// 校验失败，返回 400 错误
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(user)
		resp.OkWithMsg(c, "操作成功")
	})

	r.Run(":8080")
}
