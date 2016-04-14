/*
    Manages Token and utility functions.

    Copyright (C) 2016  Sabrina Sachs

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package Request

import (
	
    "os"
	"io/ioutil"
	"net/http"
    "encoding/json"    
    con "strconv"
    "sync"
    "time"
    "log"
)

//The API Token is stored in tokenConf.json. The struct Token is a singleton.
type token struct{
    UserToken string `json:"token"`
}
var apiToken *token
var once sync.Once

func GetLinesSeparated(lines []string)(string){
    var values string=""
    for _,l := range lines{
        values=values+l+";"
    }
    valuesTrimmed  := values[:len(values)-1]
    return valuesTrimmed
}
// Utility Functions for Time Parameters
func GetTimeFormat(t time.Time)(*string){    
    result:=t.Format(timeFormat)
    log.Println(t)
    return &result
}
func GetTimeFormatNow()(*string){
    return GetTimeFormat(time.Now())
}
func GetTimeUnixNow()(*string){
    return GetTimeUnix(time.Now())
}
func GetTimeUnix(t time.Time)(*string){
    unixTime:=t.Unix()
    result:=con.FormatInt(unixTime, 10)
    return &result
}

//Functions for api token
func getToken() string {
    once.Do(func() {
        apiToken = readToken()
    })
    apiToken=readToken()
    return apiToken.UserToken
}
func readToken() (*token){
	
    file, errFile := os.Open("tokenConf.json")
    decoder := json.NewDecoder(file)
    t := token{}
    err := decoder.Decode(&t)
    if err!=nil || errFile!=nil{
        panic("Cannot read Token. Please Check your tokenConf.json")
    }
    return &t
}
// This Function is called for all API calls. 
// It always requests json objekts. The Token is added to the Request Header.
func getContent(url string) ([]byte,error) {

	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
    log.Println("REQ: "+url)
    req.Header.Add("Accept", "application/json")
    req.Header.Add("RNV_API_TOKEN", getToken())
	log.Println(getToken())
	log.Println(req.Header)

	

    client := http.Client{}

	resp, err := client.Do(req)
	log.Println(resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}
    log.Println("RES: "+resp.Status+" "+resp.Proto)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}




