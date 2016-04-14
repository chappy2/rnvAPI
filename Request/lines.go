/*
    Provide Parameter and Request Function for Linepackage, Canceled Liens.

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
type LinesParams struct{

    HafasID *string
    LineID *string
    StopIndex *string
    TourType *string
    TourID *string
    Time *string  
}
type CanceledLineParams struct{
    LineID *string
    /*Hier wird ein Unix-
Zeitstempel (die vergangenen Millisekunden
seit dem 1. Januar 1970 00:00 Uhr) erwartet.
*/
    DepartureTime *string
}
func GetCanceledLine(params CanceledLineParams)([]rnvresponse.CanceledLine,error){
    par :=make(map[string][]string)
    par["line"]=[]string{*params.LineID} 
    par["departureTime"]=[]string{*params.DepartureTime} 
     
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+CANCELED_LINE_URL,values)
   
    
    content, err := getContent(urlV.String())
	if err != nil {
		return nil, err
	}
	var record []rnvresponse.CanceledLine
	err = json.Unmarshal(content, &record)
	if err != nil {
		return nil, err
	}
    
	return record, err
}
func GeltLineInfos()([]rnvresponse.Line,error){
    urlV:=getUrl(SITE_URL+MODUL_URL+LINE_ALL_URL)
    content, err := getContent(urlV.String())
	if err != nil {
		return nil, err
	}
	var records []rnvresponse.Line
	err = json.Unmarshal(content, &records)
	if err != nil {
		return nil, err
	} 
	return records, err
}
func GetLinesPackage()([]rnvresponse.Departures,error){
    urlV:=getUrl(SITE_URL+MODUL_URL+LINE_PACKAGE_URL)
    content, err := getContent(urlV.String())
	if err != nil {
		return nil, err
	}
	var records []rnvresponse.Departures
	err = json.Unmarshal(content, &records)
	if err != nil {
		return nil, err
	} 
	return records, err
}
// for Fahrtverlauf
func GetLines(params LinesParams) (*rnvresponse.Lines,error){

    par :=make(map[string][]string)
    par["hafasID"]=[]string{*params.HafasID} 
    par["lineID"]=[]string{*params.LineID}
    par["stopIndex"]=[]string{*params.StopIndex}
    par["tourType"]=[]string{*params.TourType}
    par["tourID"]=[]string{*params.TourID}
    var addTimeManually bool
    if(*params.Time=="null"){
       par["time"]=[]string{*params.Time}  
        addTimeManually=false
    } else if(params.Time==nil){
       par["time"]=[]string{"null"}  
        addTimeManually=false
    } else{
        addTimeManually=true
    }    
    
    
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+LINE_URL,values)
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
	var record rnvresponse.Lines
	err = json.Unmarshal(content, &record)
	if err != nil {
		return nil, err
	}
    
	return &record, err
}