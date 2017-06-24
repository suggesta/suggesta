package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-shosa/shosa/response"
	"github.com/labstack/echo"
	"github.com/suggesta/suggesta/apis/cognitive"
)

type Request struct {
	URL string `json:"url"`
}

type ResultScores struct {
	Scores cognitive.Scores `json:"scores"`
}

func EmotionLatest(c echo.Context) (err error) {
	res := cognitive.Scores{
		Anger:     8.817463E-06,
		Contempt:  0.00624216069,
		Disgust:   0.000121028206,
		Fear:      1.05626214E-06,
		Happiness: 0.828075,
		Neutral:   0.1579598,
		Sadness:   0.00755788572,
		Surprise:  3.428282E-05,
	}
	result := ResultScores{Scores: res}
	return c.JSON(200, result)
}

// Image is a handler
func Image(c echo.Context) (err error) {
	var er *response.ErrorResponse

	var rq Request
	if err = c.Bind(&rq); err != nil {
		er = response.NewErrorResponse(c, response.BadRequest)
		er.ErrorInformation.DeveloperMessage = err.Error()
		return er.JSON()
	}

	log.Print(rq.URL)
	res, code, err := cognitive.EmotionImageURL(rq.URL)
	if err != nil {
		er = response.NewErrorResponse(c, response.InternalServerError)
		er.ErrorInformation.Message = "failed to access API"
		er.ErrorInformation.DeveloperMessage = err.Error()
		return er.JSON()
	}
	if code != 200 {
		er = response.NewErrorResponse(c, response.InternalServerError)
		er.ErrorInformation.Message = fmt.Sprintf("failed to access API. code:%d", code)
		er.ErrorInformation.DeveloperMessage = err.Error()
		return er.JSON()
	}
	if len(res) == 0 {
		er = response.NewErrorResponse(c, response.InternalServerError)
		er.ErrorInformation.DeveloperMessage = "success to access API. Can't get response"
		return er.JSON()
	}
	scores := res[0].Scores
	return c.JSON(http.StatusOK, scores)
}
