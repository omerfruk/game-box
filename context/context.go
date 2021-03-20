package context

import (
	"errors"
	"github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	turkish "github.com/go-playground/validator/v10/translations/tr"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type AppCtx struct {
	*fiber.Ctx
	Db *gorm.DB
}

func (c *AppCtx) ErrorResponse(code int, msg string) error {
	model := &ResponseModel{
		HataVarMi: true,
		Message:   msg,
	}

	return c.Status(code).JSON(model)
}

//girilen degerleri cekip kontrolunu burada saglayacagiz
func (c *AppCtx) BodyParserAndValidation(model interface{}) error {
	//girilen degerleri aliyoruz
	if err := c.BodyParser(&model); err != nil {
		return errors.New("hatali model:" + err.Error())
	}
	//degerlerin kontrolu burda yapiliyor
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("labelname"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	//hatalarin turkceye cevrilmesini burda yapiyoruz
	tr := tr.New()
	uni := ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")
	turkish.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(model)
	if err != nil {
		msg := ""
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			msg += e.Translate(trans) + "\n"
		}
		return errors.New(msg)
	}
	return nil
}
