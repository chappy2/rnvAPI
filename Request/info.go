/*
    Provide Parameter and Request Function for Line and Stations Infos.

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
type StationInfoParams struct{

    Lines *[]string
    HafasID *string
    DepartureTime *string
}
type LineInfoParams struct{

    HafasID *string
    Poles *[]string
    DepartureTime *string
}

func GetLineInfo(params LineInfoParams)(*[]rnvresponse.LineInfo,error){
    par :=make(map[string][]string)
     if(params.Poles!=nil){         
         par["poles"]=[]string{GetLinesSeparated(*params.Poles)}        
    } 
    if(params.HafasID!=nil){
        par["hafasID"]=[]string{*params.HafasID} 
    } 
    if(params.DepartureTime!=nil){
        par["departureTime"]=[]string{*params.DepartureTime} 
    }     
    
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+INFO_LINE_URL,values)
    
    content, err := getContent(urlV.String())
	if err != nil {
		return nil, err
	}
	var record []rnvresponse.LineInfo
	err = json.Unmarshal(content, &record)
	if err != nil {
		return nil, err
	}
    
	return &record, err
}
func GetStationInfo(params StationInfoParams)(*[]rnvresponse.StationInfo,error){
    par :=make(map[string][]string)
     if(params.Lines!=nil){
         
         par["lines"]=[]string{GetLinesSeparated(*params.Lines)}
        
    } 
    if(params.HafasID!=nil){
        par["hafasID"]=[]string{*params.HafasID} 
    } 
    if(params.DepartureTime!=nil){
        par["departureTime"]=[]string{*params.DepartureTime} 
    }     
    
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+INFO_STATION_URL,values)
    
    content, err := getContent(urlV.String())
	if err != nil {
		return nil, err
	}
	var record []rnvresponse.StationInfo
	err = json.Unmarshal(content, &record)
	if err != nil {
		return nil, err
	}
    
	return &record, err
}