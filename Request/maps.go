/*
    Provide Parameter and Request Function maps.

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
var SupportedThumbnailSizes  = []string{"32","64","128","256","512"} //default is 128
var SupportedFormats =[]string{"pdf","png"}//default is png

type MapsParams struct {
    ThumbnailSize *string
    Format *string
}

func GetMaps(params MapsParams)(*[]rnvresponse.Maps,error){
    par :=make(map[string][]string)
     
    if(params.ThumbnailSize!=nil){
        par["thumbnailSize"]=[]string{*params.ThumbnailSize} 
    } 
    if(params.Format!=nil){
        par["format"]=[]string{*params.Format} 
    }     
    
    values:=getValues(par)
    urlV:=getUrlWithValues(SITE_URL+MODUL_URL+MAPS_URL,values)
    
    content, err := getContent(urlV.String())
	if err != nil {
		return nil, err
	}
	var record []rnvresponse.Maps
	err = json.Unmarshal(content, &record)
	if err != nil {
		return nil, err
	}
    
	return &record, err
}