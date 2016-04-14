/*
    Provide Parameter and Request Function for monitoring a Station.

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
//var UI_SOURCE_LINE="LINE"
type StationMonitorParams struct{
    //$-1 not needed anymore (testserver)
    //hafasID=116$-1&transportFilter=4&time=null&uiSource=LINE old?
    HafasId *string
    Time *string    
    //TransportFilter *string old?
    //UiSource *string old?
    Poles *string
    Mode *string
    NeedPlatformDetail *string
}

func GetStationMonitor(params StationMonitorParams) (*rnvresponse.BodyStationMonitor,error){
    par :=make(map[string][]string)
    var addTimeManually bool
       par["hafasID"]=[]string{*params.HafasId}
    if(*params.Time=="null"){
       par["time"]=[]string{*params.Time}  
        addTimeManually=false
    } else if(params.Time==nil){
       par["time"]=[]string{"null"}  
        addTimeManually=false
    } else{
        addTimeManually=true
    }
    if(params.Mode!=nil){
        par["mode"]=[]string{*params.Mode}
    }
    if(params.Poles!=nil){
        par["poles"]=[]string{*params.Poles}
    }
    if(params.NeedPlatformDetail!=nil){
        par["needPlatformDetail"]=[]string{*params.NeedPlatformDetail}
    }
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+DEPARTURES_URL,values)
    var urlString string
    if addTimeManually {        
        urlString=urlV.String()+"&time="+*params.Time
    } else {
        urlString=urlV.String()
    }
    content, err := getContent(urlString)
	if err != nil {
		return nil, err
	}
	var record rnvresponse.BodyStationMonitor
	err = json.Unmarshal(content, &record)
	if err != nil {
		return nil, err
	}
    
	return &record, err
}