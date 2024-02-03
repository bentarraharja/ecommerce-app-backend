package handler

import (
	"MyEcommerce/features/product"
	"MyEcommerce/utils/cloudinary"
	"MyEcommerce/utils/middlewares"
	"MyEcommerce/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
	cld            cloudinary.CloudinaryUploaderInterface
}

func New(ps product.ProductServiceInterface, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) *ProductHandler {
	return &ProductHandler{
		productService: ps,
		cld:            cloudinaryUploader,
	}
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	if newProduct.Name == "" {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("nama produk tidak boleh kosong", nil))
	}

	if newProduct.Price <= 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("harga produk harus lebih besar dari 0", nil))
	}

	fileHeader, err := c.FormFile("photo_product")
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
	}

	imageURL, err := handler.cld.UploadImage(fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image", nil))
	}

	productCore := RequestToCore(newProduct, imageURL, uint(userIdLogin))
	productCore.PhotoProduct = imageURL

	errInsert := handler.productService.Create(userIdLogin, productCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error create product", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert product", nil))
}

func (handler *ProductHandler) GetAllProduct(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	category := c.QueryParam("category")

	products, totalPage, err := handler.productService.GetAll(page, limit, category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error get data", nil))
	}

	productResponses := CoreToResponseListGetAllProduct(products)

	return c.JSON(http.StatusOK, responses.WebResponsePagi("success get data", productResponses, totalPage))
}

func (handler *ProductHandler) GetProductById(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing product id", nil))
	}

	result, errSelect := handler.productService.GetById(productID)
	if errSelect != nil {
		return c.JSON(http.StatusNotFound, responses.WebResponse("product not found", nil))
	}

	var productResult = CoreToResponse(*result)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", productResult))
}

func (handler *ProductHandler) UpdateProductById(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing product id", nil))
	}

	updateProduct := ProductRequest{}
	errBind := c.Bind(&updateProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	fileHeader, err := c.FormFile("photo_product")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the file", nil))
	}

	var imageURL string
	if fileHeader != nil {
		imageURL, err = handler.cld.UploadImage(fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error uploading the image", nil))
		}
	}

	productCore := RequestToCore(updateProduct, imageURL, uint(userIdLogin))
	productCore.ID = uint(productID)
	productCore.PhotoProduct = imageURL

	errUpdate := handler.productService.Update(userIdLogin, productCore)
	if errUpdate != nil {
		if errUpdate.Error() == "you do not have permission to edit this product" {
			return c.JSON(http.StatusForbidden, responses.WebResponse("you do not have permission to edit this product", nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error updating data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *ProductHandler) DeleteProductById(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing product id", nil))
	}

	errDelete := handler.productService.Delete(userIdLogin, productID)
	if errDelete != nil {
		if errDelete.Error() == "you do not have permission to delete this product" {
			return c.JSON(http.StatusForbidden, responses.WebResponse("you do not have permission to delete this product", nil))
		}
		if errDelete.Error() == "product not found" {
			return c.JSON(http.StatusNotFound, responses.WebResponse("product not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error deleting data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *ProductHandler) GetProductByUserId(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	products, err := handler.productService.GetByUserId(userIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data", nil))
	}

	productResponses := CoreToResponseListGetAllProduct(products)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data", productResponses))
}

func (handler *ProductHandler) SearchProduct(c echo.Context) error {
	query := c.QueryParam("search")
	if query == "" {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("query parameter is required", nil))
	}

	products, err := handler.productService.Search(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data", nil))
	}

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, responses.WebResponse("The provided search query is not valid. Please provide a valid search term.", nil))
	}

	productResponses := CoreToResponseListGetAllProduct(products)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data", productResponses))
}
