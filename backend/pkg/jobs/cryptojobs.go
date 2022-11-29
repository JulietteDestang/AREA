package jobs

import(
	"net/http"
	"errors"
	"io/ioutil"
	"fmt"
	"strconv"

	"github.com/tidwall/gjson"
	"AREA/pkg/utils"
	"os"
)

// const axios = require("axios");

// const options = {
//   method: 'GET',
//   url: 'https://crypto-forex.p.rapidapi.com/crypto',
//   params: {target: 'btc', base: 'usd'},
//   headers: {
//     'X-RapidAPI-Key': '54fb216729msh1db59bd41d901b7p12938ajsn6b6525d7a1c2',
//     'X-RapidAPI-Host': 'crypto-forex.p.rapidapi.com'
//   }
// };

// axios.request(options).then(function (response) {
// 	console.log(response.data);
// }).catch(function (error) {
// 	console.error(error);
// });

func GetCryptodata(target string, base string) (string, error) {

	url := fmt.Sprintf("https://crypto-forex.p.rapidapi.com/currencies?target=%s&base=%s", target, base)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "54fb216729msh1db59bd41d901b7p12938ajsn6b6525d7a1c2")
	req.Header.Add("X-RapidAPI-Host", "crypto-forex.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("heads-tails api down")
		return "", myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	result := gjson.GetBytes(body, "outcome").String()

	fmt.Println(string(body))

	return result, nil 
}

func CryptoIsOverN(params string) bool{
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 3 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return false
	}
	result, _ := GetCryptodata(paramsArr[0], paramsArr[1])
	newparam, _ :=strconv.ParseFloat(paramsArr[2], 64)
	newresult, _ :=strconv.ParseFloat(result, 64)
	if (newresult > newparam) {
		return true
	} else {
		return false
	}
}

func CryptoIsUnderN(params string) bool{
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 3 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return false
	}
	result, _ := GetCryptodata(paramsArr[0], paramsArr[1])
	newparam, _ :=strconv.ParseFloat(paramsArr[2], 64)
	newresult, _ :=strconv.ParseFloat(result, 64)
	if (newresult < newparam) {
		return true
	} else {
		return false
	}
}