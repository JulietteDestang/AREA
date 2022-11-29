/** @file covidJob.go
 * @brief This file contain all the functions to handle the actions and reactions of the Covid API
 * @author Juliette Destang
 * 
 */

package jobs

// @cond

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"errors"
	"strconv"
	"os"

	"github.com/tidwall/gjson"
)

// @endcond

/** @brief Retrieves all the data concerning covid cases from the covid API
 * @param [] byte, error
 */
func GetCovidData() ([] byte, error) {

	url := "https://covid-193.p.rapidapi.com/statistics?country=france"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "54fb216729msh1db59bd41d901b7p12938ajsn6b6525d7a1c2")
	req.Header.Add("X-RapidAPI-Host", "covid-193.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("covid api down")
		return nil, myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil 
}

/** @brief Returns true if the covid cases are over than the params passed as argument
 * @param params string
 * @return bool
 */
func CovidCaseIsOverN(params string) (bool) {

	covidData, Err := GetCovidData()
	if (Err != nil) {
		fmt.Fprintln(os.Stderr, Err)
		return false
	}
	data := gjson.GetBytes(covidData, "response.0.cases.active")
	floatCase, _ := strconv.ParseFloat(data.String(), 64)
	compareCaseNb, _ := strconv.ParseFloat(params, 64)
	if (floatCase > compareCaseNb) {
		return true
	} else {
		return false
	}
}

/** @brief Returns true if the covid critical cases are over than the params passed as argument
 * @param params string
 * @return bool
 */
func CovidCriticalCaseIsOverN(params string) (bool) {
	covidData, Err := GetCovidData()
	if (Err != nil) {
		fmt.Fprintln(os.Stderr, Err)
		return false
	}
	data := gjson.GetBytes(covidData, "response.0.cases.critical")
	floatCase, _ := strconv.ParseFloat(data.String(), 64)
	compareCaseNb, _ := strconv.ParseFloat(params, 64)
	if (floatCase > compareCaseNb) {
		return true
	} else {
		return false
	}
}