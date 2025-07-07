package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	// apiresp "gitlab.com/devkit3/apiresponse"
// )

// const (
// 	InternalServerError = "internal server error"
// 	BadRequest          = "bad request"
// )

// type Response interface {
// 	ErrorResponse(c *gin.Context, err error, code int, message string, isUserFacing bool)
// 	ResultResponse(c *gin.Context, clientMsg string, datatype apiresp.ResponseDataType, data any)
// }

// type ResultOk struct {
// 	Status   string `json:"status"` // ok
// 	Response struct {
// 		Message string      `json:"message"`
// 		Type    string      `json:"type"`           // [AVALIABLE]: object, array, empty
// 		Data    interface{} `json:"data,omitempty"` // [AVALIABLE]: object, array of objects, empty
// 	} `json:"response"`
// }

// type ResultError struct {
// 	Status   string `json:"status"` // error
// 	Response struct {
// 		Code    int    `json:"code"` // [RULE]: must be one of codes from table (Check DEV.PAGE)
// 		Message string `json:"message"`
// 	} `json:"response"`
// }

// func (h *Handler) ErrorResponse(c *gin.Context, err error, code int, message string, isUserFacing bool) {
// 	errText := "empty error"
// 	if err != nil {
// 		errText = err.Error()
// 	}

// 	h.logger.Error(fmt.Sprintf("(Client Answer) %s : (Dev Message) %s", message, errText))

// 	errMessage := message
// 	if err != nil && isUserFacing {
// 		errMessage = fmt.Sprintf("%s, %s", message, err.Error())
// 	}

// 	httpError := apiresp.NewError(code, errMessage)
// 	httpResult := apiresp.NewResult(apiresp.ERROR, httpError)

// 	c.Header("Content-Type", "application/json")
// 	c.Writer.WriteHeader(code)

// 	if err := json.NewEncoder(c.Writer).Encode(httpResult); err != nil {
// 		h.logger.Error("invalid write resultError", "(error)", err.Error())
// 	}
// }

// func (h *Handler) ResultResponse(c *gin.Context, clientMsg string, datatype apiresp.ResponseDataType, data any) {
// 	httpResponse := apiresp.NewResponse(clientMsg, datatype, data)
// 	httpResult := apiresp.NewResult(apiresp.OK, httpResponse)

// 	c.Header("Content-Type", "application/json")
// 	c.Writer.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(c.Writer).Encode(httpResult); err != nil {
// 		h.logger.Error("invalid write resultOk", "(error)", err.Error())
// 	}
// }
