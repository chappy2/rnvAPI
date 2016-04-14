/*
    Provide Parameter and Request Function Ticker and News.

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
    "encoding/json"
    rnvresponse "rnvApi/Response"
)
// Same Response Object News for Ticker. The Lines must be given as
// Query Parameter, in this implementation all Lines as String in an array.
func GetTickerForLines(lines []string)([]rnvresponse.News,error){
    var values string="?lines="
    for _,l := range lines{
        values=values+l+";"
    }
    valuesTrimmed  := values[:len(values)-1]
    urlV:=getUrl(SITE_URL+MODUL_URL+TICKER_URL)
    content, err := getContent(urlV.String()+valuesTrimmed)
	if err != nil {
		return nil, err
	}
	var records []rnvresponse.News
	err = json.Unmarshal(content, &records)
	if err != nil {
		return nil, err
	} 
	return records, err
}

func GetNews()([]rnvresponse.News,error){
    urlV:=getUrl(SITE_URL+MODUL_URL+NEWS_URL)
    content, err := getContent(urlV.String())
	if err != nil {
		return nil, err
	}
	var records []rnvresponse.News
	err = json.Unmarshal(content, &records)
	if err != nil {
		return nil, err
	} 
	return records, err
}