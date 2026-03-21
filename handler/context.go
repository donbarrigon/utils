package handler

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/donbarrigon/utils/auth"
	"github.com/donbarrigon/utils/herror"
	"github.com/donbarrigon/utils/lang"
	"github.com/donbarrigon/utils/str"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// type Auth interface {
// 	Can(permission string) error
// 	HasRole(role string) error
// 	UserID() bson.ObjectID
// }

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	handler *Handler
	Auth    *auth.Session
}

func NewContext(w http.ResponseWriter, r *http.Request, h *Handler) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		handler: h,
	}
}

func (c *Context) Lang() string {
	return c.Request.Header.Get("Accept-Language")
}

func (c *Context) GetBodyJson(request any) herror.Error {
	defer c.Request.Body.Close()
	decoder := json.NewDecoder(c.Request.Body)
	if e := decoder.Decode(request); e != nil {
		return herror.BadRequestMsg(e, "The request body could not be read or contains malformed JSON.")
	}
	return nil
}

func (c *Context) GetBodyMsgpack(request any) herror.Error {
	defer c.Request.Body.Close()
	decoder := msgpack.NewDecoder(c.Request.Body)
	if e := decoder.Decode(request); e != nil {
		return herror.BadRequestMsg(e, "The request body could not be read or contains malformed MessagePack data.")
	}
	return nil
}

func (c *Context) GetBody(request any) herror.Error {
	defer c.Request.Body.Close()
	if c.Request.Header.Get("Content-Type") == "application/json" {
		return c.GetBodyJson(request)
	}
	return c.GetBodyMsgpack(request)
}

func (c *Context) Get(param string) string {
	return c.Request.URL.Query().Get(param)
}

func (c *Context) ResponseJSON(status int, data any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)

	if err := json.NewEncoder(c.Writer).Encode(data); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.WriteHeader(500)
		c.Writer.Write([]byte(lang.T(c.Lang(), `{"message": "Error", "error": "Could not encode the response"}`, nil)))
	}
}

func (c *Context) ResponseMsgpack(status int, data any) {
	c.Writer.Header().Set("Content-Type", "application/msgpack")
	c.Writer.WriteHeader(status)

	if err := msgpack.NewEncoder(c.Writer).Encode(data); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(`{"message": "Error", "error": "Could not encode the response"}`))
	}
}

func (c *Context) Response(status int, data any) {
	if c.Request.Header.Get("accept") == "application/json" {
		c.ResponseJSON(status, data)
		return
	}
	c.ResponseMsgpack(status, data)
}

func (c *Context) ResponseError(e herror.Error) {
	e.Translate(c.Lang())
	if c.Request.Header.Get("accept") == "application/json" {
		e.WriteJSON(c.Writer)
		return
	}
	e.WriteProto(c.Writer)
}

func (c *Context) ResponseNotFound() {
	c.ResponseError(herror.NotFoundMsg(nil,
		lang.T(c.Lang(), "The resource [:method :path] does not exist", str.Placeholder{
			{Key: "method", Value: c.Request.Method},
			{Key: "path", Value: c.Request.URL.Path},
		}),
	))
}

func (c *Context) ResponseOk(data any) {
	if c.Request.Header.Get("accept") == "application/json" {
		c.ResponseJSON(http.StatusOK, data)
		return
	}
	c.ResponseMsgpack(http.StatusOK, data)
}

func (c *Context) ResponseCreated(data any) {
	if c.Request.Header.Get("accept") == "application/json" {
		c.ResponseJSON(http.StatusCreated, data)
		return
	}
	c.ResponseMsgpack(http.StatusCreated, data)
}

func (c *Context) ResponseNoContent() {
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (c *Context) ResponseCSV(fileName string, data any, comma ...rune) {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		e := herror.InternalServerErrorMsg(
			errors.New("CSV serialization failed: data must be a slice of structs, got unsupported type"),
			"The file could not be generated. Please try again later or contact support if the issue persists.",
		)
		c.ResponseError(e)
		return
	}

	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	del := ';'
	if len(comma) > 0 {
		del = comma[0]
	}
	writer.Comma = del

	if val.Len() == 0 {
		e := herror.NotFoundMsg(nil, "The requested data could not be found or does not exist.")
		c.ResponseError(e)
		return
	}

	first := val.Index(0)
	elemType := first.Type()
	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	var headers []string
	var fields []int

	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		tag = strings.Split(tag, ",")[0]
		headers = append(headers, tag)
		fields = append(fields, i)
	}
	writer.Write(headers)

	for i := 0; i < val.Len(); i++ {
		var record []string
		elem := val.Index(i)
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}

		for _, j := range fields {
			fieldVal := elem.Field(j)

			if fieldVal.Type() == reflect.TypeOf(bson.ObjectID{}) {
				objID := fieldVal.Interface().(bson.ObjectID)
				record = append(record, objID.Hex()) // sin comillas manuales
				continue
			}

			switch fieldVal.Kind() {
			case reflect.String:
				record = append(record, fieldVal.String())
			case reflect.Int, reflect.Int64:
				record = append(record, fmt.Sprintf("%d", fieldVal.Int()))
			case reflect.Float64:
				record = append(record, fmt.Sprintf("%f", fieldVal.Float()))
			case reflect.Bool:
				record = append(record, fmt.Sprintf("%t", fieldVal.Bool()))
			case reflect.Struct:
				if t, ok := fieldVal.Interface().(time.Time); ok {
					record = append(record, t.Format(time.RFC3339))
				} else {
					jsonVal, _ := json.Marshal(fieldVal.Interface())
					record = append(record, string(jsonVal))
				}
			case reflect.Slice, reflect.Map, reflect.Array:
				jsonVal, _ := json.Marshal(fieldVal.Interface())
				record = append(record, string(jsonVal))
			default:
				record = append(record, fmt.Sprintf("%v", fieldVal.Interface()))
			}
		}
		writer.Write(record)
	}
	writer.Flush()

	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename="+fileName+".csv")
	c.Writer.Write(buffer.Bytes())
}
