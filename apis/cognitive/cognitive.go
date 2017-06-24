package cognitive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type Param struct {
	URL string `json:"url"`
}

type ResultEmotion struct {
	FaceRectangle FaceRectangle `json:"faceRectangle"`
	Scores        Scores        `json:"scores"`
}

type FaceRectangle struct {
	Height int `json:"height"`
	Left   int `json:"left"`
	Top    int `json:"top"`
	Width  int `json:"width"`
}

type Scores struct {
	Anger     float64 `json:"anger"`
	Contempt  float64 `json:"contempt"`
	Disgust   float64 `json:"disguest"`
	Fear      float64 `json:"fear"`
	Happiness float64 `json:"happiness"`
	Neutral   float64 `json:"neutral"`
	Sadness   float64 `json:"sadness"`
	Surprise  float64 `json:"surprise"`
}

// BaseURL is url
const BaseURL = "https://westus.api.cognitive.microsoft.com"

// EmotionImageBinary is request to cognitive emotion API
func EmotionImageBinary(f *os.File) (request interface{}, err error) {
	// if os.Getenv("MS_SUBSCRIPTION_KEY") == "" {
	// 	return nil, fmt.Errorf("access key is not set")
	// }

	reqURL := BaseURL + "/emotion/v1.0/recognize"
	req, _ := http.NewRequest("POST", reqURL, f)
	req.Header.Set("Content-Type", "application/octet-stream")
	// req.Header.Set("Ocp-Apim-Subscription-Key", os.Getenv("MS_SUBSCRIPTION_KEY"))
	req.Header.Set("Ocp-Apim-Subscription-Key", "d1773fe0c0844d01a212a95b8807aa27")

	client := new(http.Client)
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &request)
	return request, err
}

// EmotionImageURL is request to cognitive emotion API
func EmotionImageURL(url string) (result []ResultEmotion, code int, err error) {
	// if os.Getenv("MS_SUBSCRIPTION_KEY") == "" {
	// 	return result, 500, fmt.Errorf("access key is not set")
	// }

	param := Param{
		URL: url,
	}

	reqURL := BaseURL + "/emotion/v1.0/recognize"
	log.Print(reqURL)
	log.Print(param)
	r := gorequest.New()
	resp, body, errs := r.Post(reqURL).
		Send(param).
		Set("Content-Type", "application/json").
		// Set("Ocp-Apim-Subscription-Key", os.Getenv("MS_SUBSCRIPTION_KEY")).
		Set("Ocp-Apim-Subscription-Key", "d1773fe0c0844d01a212a95b8807aa27").
		End()
	if errs != nil {
		return result, resp.StatusCode, errs[0]
	}
	log.Print(body)

	if resp.StatusCode != 200 {
		return result, resp.StatusCode, fmt.Errorf("Failed to request cognitive emotion API")
	}

	d := json.NewDecoder(strings.NewReader(body))
	err = d.Decode(&result)
	return result, resp.StatusCode, err
}
