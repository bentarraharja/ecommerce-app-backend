package handler

import (
	"MyEcommerce/features/order"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService order.OrderServiceInterface
}

func New(os order.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: os,
	}
}

func (handler *OrderHandler) CreateOrder(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newOrder := OrderRequest{}
	errBind := c.Bind(&newOrder)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data order not valid", nil))
	}

	orderCore := RequestToCoreOrder(newOrder)
	payment, errInsert := handler.orderService.CreateOrder(userIdLogin, newOrder.CartIDs, orderCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert order", nil))
	}

	result := CoreToResponse(payment)

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", result))
}

func (handler *OrderHandler) GetOrderUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	results, errSelect := handler.orderService.GetOrderUser(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var orderResult []GetOrderUserResponse
	for _, result := range results {
		orderResult = append(orderResult, CoreToResponseOrderUser(result.Order, []order.OrderItemCore{result})...)
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", orderResult))

}

func (handler *OrderHandler) GetOrderAdmin(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	

	results, totalPage, errSelect := handler.orderService.GetOrderAdmin(userIdLogin, page, limit)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	response := CoreToResponseOrderAdmin(results)
	return c.JSON(http.StatusOK, responses.WebResponsePagi("success read data", response, totalPage))
}

func (handler *OrderHandler) CancleOrderById(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	orderId := c.Param("id")

	updateOrderStatus := CancleOrderRequest{}
	errBind := c.Bind(&updateOrderStatus)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	orderCore := CancleRequestToCoreOrder(updateOrderStatus)
	errCancle := handler.orderService.CancleOrder(userIdLogin, orderId, orderCore)
	if errCancle != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error cancle order", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success cancle order", nil))
}

func (handler *OrderHandler) WebhoocksNotification(c echo.Context) error {
	var reqNotif = WebhoocksRequest{}
	errBind := c.Bind(&reqNotif)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	orderCore := WebhoocksRequestToCore(reqNotif)
	err := handler.orderService.WebhoocksService(orderCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error Notif "+err.Error(), nil))
	}

	log.Println("transaction success")
	return c.JSON(http.StatusOK, responses.WebResponse("transaction success", nil))
}
