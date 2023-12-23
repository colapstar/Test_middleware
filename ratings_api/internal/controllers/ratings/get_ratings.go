package ratings

import (
	"middleware/example3/internal/helpers"
	"middleware/example3/internal/models"
	"middleware/example3/internal/services/ratings"
	"net/http"

	"github.com/sirupsen/logrus"
)

// test
// GetRatings
// @Tags         ratings
// @Summary      Get ratings.
// @Description  Get ratings.
// @Success      200            {array}  models.Rating
// @Failure      500            "Something went wrong"
// @Router       /ratings [get]
func GetRatings(w http.ResponseWriter, r *http.Request) {
	ratings, err := ratings.GetAllRatings()
	if err != nil {
		logrus.Errorf("error: %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			helpers.RespondWithFormat(w, r, customError)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			helpers.RespondWithFormat(w, r, "Something went wrong")
		}
		return
	}

	helpers.RespondWithFormat(w, r, ratings)
}
