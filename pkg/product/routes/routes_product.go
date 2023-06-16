package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/product/pb"
	"github.com/xuri/excelize/v2"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

type PaginationsRequestBody struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := CreateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func ListProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	page, _ := strconv.ParseInt(ctx.Query("page"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Query("limit"), 10, 64)

	res, err := c.ListProduk(context.Background(), &pb.ListproductsRequest{
		Page:  page,
		Limit: limit,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func DownloadProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	type DownloadBody struct {
		TotalData int64  `json:"limit"`
		Direction string `json:"direction"`
	}
	body := DownloadBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var enumdirection pb.Direction
	switch body.Direction {
	case "ASC", "asc":
		enumdirection = pb.Direction_ASC
	case "DESC", "desc":
		enumdirection = pb.Direction_DESC
	default:
		enumdirection = pb.Direction_ASC
	}

	res, err := c.DownloadDataProduct(context.Background(), &pb.DownloadDataProductRequest{
		TotalData: body.TotalData,
		Direction: enumdirection,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	f := excelize.NewFile()
	f.SetCellValue("sheet1", "A1", "ID")
	f.SetCellValue("sheet1", "B1", "Name")
	f.SetCellValue("sheet1", "C1", "Stock")
	f.SetCellValue("sheet1", "D1", "Price")

	row := 2
	for _, data := range res.Data {
		f.SetCellValue("sheet1", fmt.Sprintf("A%d", row), data.Id)
		f.SetCellValue("sheet1", fmt.Sprintf("B%d", row), data.Name)
		f.SetCellValue("sheet1", fmt.Sprintf("C%d", row), data.Stock)
		f.SetCellValue("sheet1", fmt.Sprintf("D%d", row), data.Price)
		row++
	}

	currentTime := time.Now()
	date := currentTime.Format("2006-01-02T15:04:05")

	filename := "List_ProductD" + date + ".xlsx"
	if err := f.SaveAs("./" + filename); err != nil {
		ctx.String(http.StatusInternalServerError, "failed save file: %v", err)
		return
	}

	data, _ := f.WriteToBuffer()
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)

	// Mengirimkan file Excel sebagai respons
	ctx.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data.Bytes())

	err = os.Remove("./" + filename)
	if err != nil {
		log.Println("failed remove file:", err)
	}
}
