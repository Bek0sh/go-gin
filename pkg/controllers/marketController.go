package controllers

import interfaces "project2/pkg/service/iservice"

type MarketController struct {
	service interfaces.MarketServiceInterface
}

func NewMarketController(service interfaces.MarketServiceInterface) *MarketController {
	return &MarketController{service: service}
}
