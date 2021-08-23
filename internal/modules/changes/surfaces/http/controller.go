package http

import (
	"diffme.dev/diffme-api/internal/core"
	errors2 "diffme.dev/diffme-api/internal/core/errors"
	domain "diffme.dev/diffme-api/internal/modules/changes"
	"github.com/gofiber/fiber/v2"
	"log"
)

type ChangeController struct {
	changeUseCases domain.ChangeUseCases
}

func (e *ChangeController) GetChanges(c *fiber.Ctx) error {
	query := new(domain.QueryChangesRequest)

	if err := c.QueryParser(query); err != nil {
		return err
	}

	errors := core.ValidateStruct(query)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	log.Printf("Query %+v", *query)

	changes, err := e.changeUseCases.GetChanges(domain.QueryChangesRequest{
		Limit: query.Limit,
		Sort: query.Sort,
		Before: query.Before,
		After: query.After,
	})

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	response := struct {
		Changes []domain.Change `json:"changes"`
	}{
		Changes: changes,
	}

	return c.JSON(response)
}

func (e *ChangeController) SearchChanges(c *fiber.Ctx) error {

	search := new(domain.SearchRequest)

	if err := c.QueryParser(search); err != nil {
		return err
	}

	errors := core.ValidateStruct(search)

	if len(errors) > 0 {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid json.")
	}

	log.Printf("Search Query: %+v", *search)

	searchChanges, err := e.changeUseCases.SearchChange(*search)

	if err != nil {
		apiError := errors2.ApiError{
			Message: err.Error(),
			StatusCode: fiber.StatusBadRequest,
		}

		return c.Status(apiError.StatusCode).JSON(apiError)
	}

	response := struct {
		Changes []domain.SearchChange `json:"changes"`
	}{
		Changes: searchChanges,
	}

	return c.JSON(response)
}

func (e *ChangeController) GetChangeByReferenceID(c *fiber.Ctx) error {

	referenceId := c.Params("id")

	changes, err := e.changeUseCases.FetchChangeForReferenceId(referenceId)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return c.JSON(changes)
}