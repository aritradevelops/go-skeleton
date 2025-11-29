package http

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"skeleton-test/internal/handlers"
	"skeleton-test/internal/translation"
	"skeleton-test/internal/validation"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func errorHandler(c *fiber.Ctx, e error) error {
	if e != nil {
		// handle validation error
		if errs, ok := e.(validation.ValidationErrors); ok {
			c.Status(http.StatusBadRequest)
			localizedErrors := fiber.Map{}
			for _, err := range errs {
				err.Message = translation.Localize(c, fmt.Sprintf("validation.%s", err.Code), map[string]any{
					"Param": err.Param,
					"Field": Labelize(err.Field),
				})
				localizedErrors[err.Field] = err
			}
			return c.JSON(handlers.NewErrorResponse(translation.Localize(c, fmt.Sprintf("errors.%d", http.StatusUnprocessableEntity)), localizedErrors))
		}

		// handle database error
		var pgErr *pgconn.PgError
		if errors.As(e, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				c.Status(http.StatusConflict)
				err := validation.ValidationError{
					Message: translation.Localize(c, "validation.unique", map[string]any{
						"Field": Labelize(getUniqueViolationField(pgErr.Detail)),
					}),
				}
				return c.JSON(handlers.NewErrorResponse(translation.Localize(c, fmt.Sprintf("errors.%d", http.StatusConflict)), fiber.Map{
					getUniqueViolationField(pgErr.Detail): err,
				}))
			}
		}

		c.Status(http.StatusInternalServerError)
		return c.JSON(handlers.NewErrorResponse(translation.Localize(c, fmt.Sprintf("errors.%d", http.StatusInternalServerError)), e))
	}
	return nil
}

func Labelize(field string) string {
	formatted := []string{}
	for _, part := range strings.Split(field, "_") {
		formatted = append(formatted, strings.Title(part))
	}
	fmt.Println("field", field, "label", strings.Join(formatted, " "))
	return strings.Join(formatted, " ")
}

func getUniqueViolationField(detail string) string {
	// example: Key (email)=(test@gmail.com) already exists.
	re := regexp.MustCompile(`\(([^)]+)\)=\(([^)]+)\)`)
	matches := re.FindStringSubmatch(detail)

	if len(matches) == 3 {
		column := matches[1] // "email"
		// value := matches[2]  // "test@gmail.com"
		return column
	}
	return ""
}
