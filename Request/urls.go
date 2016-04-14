/*
    Manage the api urls.

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
    "net/url"
)


const PROTOCOL="http"
const API_DOMAIN = "rnv.the-agent-factory.de:8080"
const PROTCOL_AND_API_COMAIN="http://rnv.the-agent-factory.de:8080"
const SITE_URL = "/easygo2/api/"
const MODUL_URL="regions/rnv/modules/"
//Linienpakete
const LINE_PACKAGE_URL=LINE_URL+"/allJourney"
const STATIONS_URL = "stations/packages/"
const UPDATE_URL="update"
//Haltestellenmonitor
const DEPARTURES_URL = "stationmonitor/element"
//Linienverlauf
const LINE_URL = "lines"
const LINE_ALL_URL = LINE_URL + "/all"
//Nachrichten
const NEWS_URL = "news"
const NEWS_COUNT_URL = "news/numberOfNewEntries/0" //alt? nicht ind oku
//Ticker
const TICKER_URL = "ticker"
const TICKER_COUNT_URL = "ticker/numberOfNewEntries/0" //alt? nicht in doku
//Entfallende Linien
const CANCELED_LINE_URL="canceled/line"
//Info
const INFO_STATION_URL="info/station"
const INFO_LINE_URL="info/journey"
//Netzpläne
const MAPS_URL="maps"

func getValues(params map[string][]string)(*url.Values){
    val:=url.Values{}
    for k,v :=range params{
        val.Set(k,v[0]) 
    }
    return &val
}
func getUrl(path string)(*url.URL){
    return &url.URL{
			Host:   API_DOMAIN,
			Scheme: PROTOCOL,
            Path:path,
		}
}
func getUrlWithValues(path string,params *url.Values)(*url.URL){
    return &url.URL{
			Host:   API_DOMAIN,
			Scheme: PROTOCOL,
            Path:path,
        RawQuery:params.Encode(),
		}
}
