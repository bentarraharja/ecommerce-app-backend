package handler

import (
	"MyEcommerce/features/cart"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartService cart.CartServiceInterface
}

func New(cs cart.CartServiceInterface) *CartHandler {
	return &CartHandler{
		cartService: cs,
	}
}

func (handler *CartHandler) CreateCart(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing product id", nil))
	}

	errInsert := handler.cartService.Create(userIdLogin, productID)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *CartHandler) UpdateCart(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	cartID, err := strconv.Atoi(c.Param("cart_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing cart id", nil))
	}

	updateCart := CartRequest{}
	errBind := c.Bind(&updateCart)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	cartCore := RequestToCore(updateCart)

	errUpdate := handler.cartService.UpdateCart(userIdLogin, cartID, cartCore)
	if errUpdate != nil {
		if errUpdate.Error() == "you do not have permission to edit this product" {
			return c.JSON(http.StatusForbidden, responses.WebResponse("you do not have permission to edit this product", nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error updating data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *CartHandler) DeleteProductCart(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	cartID, err := strconv.Atoi(c.Param("cart_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing cart id", nil))
	}

	errDelete := handler.cartService.DeleteCart(userIdLogin, cartID)
	if errDelete != nil {
		if errDelete.Error() == "you do not have permission to delete this product" {
			return c.JSON(http.StatusForbidden, responses.WebResponse("you do not have permission to delete this product", nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error deleting data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *CartHandler) GetProductCart(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	results, errSelect := handler.cartService.Get(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var cartResult = CoreToResponseList(results)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", cartResult))
}
