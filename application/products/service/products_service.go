package service

import (
	"context"
	"jastip-core/application/products/repository"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/valyala/fasthttp"
)

type productsService struct {
	repo repository.ProductsRepositoryContract
}

func NewProductsService(repo repository.ProductsRepositoryContract) *productsService {
	return &productsService{
		repo: repo,
	}
}

func (s *productsService) Create(ctx context.Context, poolData *domain.Config, data domain.ProductData, r *fasthttp.Request) (err domain.ErrorData) {
	file, _ := r.MultipartForm()

	fileHeader, compressImg, errData := validationImage(file)
	if errData.Code != 0 {
		return
	}
	err = saveProduct(ctx, poolData, *fileHeader, compressImg, data, s.repo)
	return
}

func (s *productsService) GetList(poolData *domain.Config, userID int, params domain.Params) (totalPage int, currentPage int, total int64, limit int, result []domain.ProductResp, err domain.ErrorData) {
	result, currentPage, limit, total, err = getList(poolData, userID, params, s.repo)
	if err.Code != 0 {
		return
	}

	totalPage = int(handler.CalculateTotalPages(total, int64(params.Limit)))
	return
}

func (s *productsService) Update(poolData *domain.Config, id int, userID int, update map[string]any) (err domain.ErrorData) {
	err = updateProducts(poolData, s.repo, id, userID, update)
	return
}

func (s *productsService) Delete(poolData *domain.Config, id int, userId int) (err domain.ErrorData) {
	err = deleteProducts(poolData, s.repo, id, userId)
	return
}
