package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-shosa/shosa/response"
	"github.com/labstack/echo"
	"github.com/suggesta/suggesta/apis/cognitive"
	"github.com/suggesta/suggesta/apis/database"
)

type Request struct {
	URL string `json:"url"`
}

type Emotion struct {
	Anger     float64 `json:"anger"`
	Contempt  float64 `json:"contempt"`
	Disgust   float64 `json:"disguest"`
	Fear      float64 `json:"fear"`
	Happiness float64 `json:"happiness"`
	Neutral   float64 `json:"neutral"`
	Sadness   float64 `json:"sadness"`
	Surprise  float64 `json:"surprise"`
	CreatedAt int64   `json:"created_at"`
}

// type ResultScores struct {
// 	Scores cognitive.Scores `json:"scores"`
// }

type ResultScores struct {
	Anger     float64 `json:"anger"`
	Contempt  float64 `json:"contempt"`
	Disgust   float64 `json:"disguest"`
	Fear      float64 `json:"fear"`
	Happiness float64 `json:"happiness"`
	Neutral   float64 `json:"neutral"`
	Sadness   float64 `json:"sadness"`
	Surprise  float64 `json:"surprise"`
	Max       string  `json:"max"`
}

type ResultScoresSummary struct {
	Summary cognitive.Scores `json:"summary"`
}

type ResultCalendar struct {
	Event   string `json:"event"`
	StartAt int64  `json:"start_at"`
	EndAt   int64  `json:"end_at"`
}

func Calendar(c echo.Context) (err error) {
	result := ResultCalendar{
		Event:   "外出中",
		StartAt: 1498352400,
		EndAt:   1498363200,
	}
	return c.JSON(200, result)
}

func EmotionIndex(c echo.Context) (err error) {
	result, err := emotionIndex()
	if err != nil {
		er := response.NewErrorResponse(c, response.InternalServerError)
		er.ErrorInformation.DeveloperMessage = err.Error()
		return er.JSON()
	}
	return c.JSON(200, result)
}

func EmotionSummary(c echo.Context) (err error) {
	var er *response.ErrorResponse
	result := cognitive.Scores{}
	rs, err := emotionIndex()
	if err != nil {
		er = response.NewErrorResponse(c, response.InternalServerError)
		er.ErrorInformation.DeveloperMessage = err.Error()
		return er.JSON()
	}
	if len(rs) == 0 {
		return c.JSON(200, ResultScoresSummary{Summary: result})
	}
	cnt := float64(len(rs))
	for _, r := range rs {
		result.Anger = result.Anger + r.Anger
		result.Contempt = result.Contempt + r.Contempt
		result.Disgust = result.Disgust + r.Disgust
		result.Fear = result.Fear + r.Fear
		result.Happiness = result.Happiness + r.Happiness
		result.Neutral = result.Neutral + r.Neutral
		result.Sadness = result.Sadness + r.Sadness
		result.Surprise = result.Surprise + r.Surprise
	}
	if result.Anger > 0 {
		result.Anger = result.Anger / cnt
	}
	if result.Contempt > 0 {
		result.Contempt = result.Contempt / cnt
	}
	if result.Disgust > 0 {
		result.Disgust = result.Disgust / cnt
	}
	if result.Fear > 0 {
		result.Fear = result.Fear / cnt
	}
	if result.Happiness > 0 {
		result.Happiness = result.Happiness / cnt
	}
	if result.Neutral > 0 {
		result.Neutral = result.Neutral / cnt
	}
	if result.Sadness > 0 {
		result.Sadness = result.Sadness / cnt
	}
	if result.Surprise > 0 {
		result.Surprise = result.Surprise / cnt
	}
	return c.JSON(200, ResultScoresSummary{Summary: result})
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

	db, err := database.Connect()
	if err != nil {
		er = response.NewErrorResponse(c, response.InternalServerError)
		er.ErrorInformation.Message = "failed to access database"
		er.ErrorInformation.DeveloperMessage = err.Error()
		return er.JSON()
	}
	defer db.Close()

	db.AutoMigrate(&Emotion{})

	db.Create(&Emotion{
		Anger:     scores.Anger,
		Contempt:  scores.Contempt,
		Disgust:   scores.Disgust,
		Fear:      scores.Fear,
		Happiness: scores.Happiness,
		Neutral:   scores.Neutral,
		Sadness:   scores.Sadness,
		Surprise:  scores.Surprise,
		CreatedAt: time.Now().Unix(),
	})

	result := ResultScores{
		Anger:     scores.Anger,
		Contempt:  scores.Contempt,
		Disgust:   scores.Disgust,
		Fear:      scores.Fear,
		Happiness: scores.Happiness,
		Neutral:   scores.Neutral,
		Sadness:   scores.Sadness,
		Surprise:  scores.Surprise,
	}

	result.Max = "Neutral"
	if result.Anger > result.Neutral {
		result.Max = "Anger"
	} else if result.Contempt > result.Neutral {
		result.Max = "Contempt"
	} else if result.Disgust > result.Neutral {
		result.Max = "Disgust"
	} else if result.Fear > result.Neutral {
		result.Max = "Fear"
	} else if result.Happiness > result.Neutral {
		result.Max = "Happiness"
	} else if result.Sadness > result.Neutral {
		result.Max = "Sadness"
	} else if result.Surprise > result.Neutral {
		result.Max = "Surprise"
	}

	return c.JSON(http.StatusOK, result)
}

func emotionIndex() (result []cognitive.Scores, err error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var emotions []Emotion
	db.Order("created_at desc").Limit(5).Find(&emotions)

	for _, em := range emotions {
		result = append(result, cognitive.Scores{
			Anger:     em.Anger,
			Contempt:  em.Contempt,
			Disgust:   em.Disgust,
			Fear:      em.Fear,
			Happiness: em.Happiness,
			Neutral:   em.Neutral,
			Sadness:   em.Sadness,
			Surprise:  em.Surprise,
		})
	}
	if len(result) == 0 {
		return []cognitive.Scores{}, nil
	}
	return result, nil
}
