/*
    Provide Parameter and Request Function for availability of new line or Ssation packages.

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
const timeFormatDate="2006-01-02"
const timeFormatTime="15:04"
const timeFormat="2006-01-02+15:04"
//Mon Jan 2 15:04:05 MST 2006
const timeSeparator="$"
//update?regionID=1&time=2014-07-08+13:40$2014-07-08+13:40$2011-11-11+11:11
//http://rnv-test.the-agent-factory.de:8080/easygo2/api/update?regionId=1&time=2013-07-08+13:40$2013-07-08+13:40$2013-07-08+13:40
type UpdateParams struct{
    RegionID string
    //Time in Pattern yyyy-MM-dd+HH:mm separated by $
    LastUpdateStationsTime *string
    LastUpdateLinesTime *string
    UnusedTime *string
}


func GetUpdates(params UpdateParams) (*[]rnvresponse.Update,error){
    content, err := getContent(getUrlForUpdate(params))
	if err != nil {
		return nil, err
	}
	var records []rnvresponse.Update
	err = json.Unmarshal(content, &records)
	if err != nil {
		return nil, err
	}
    
	return &records, err
}
func getUrlForUpdate(params UpdateParams) (string){
    return PROTCOL_AND_API_COMAIN+SITE_URL+UPDATE_URL+"?regionID="+params.RegionID+"&time="+*params.LastUpdateStationsTime+timeSeparator+*params.LastUpdateLinesTime+timeSeparator+*params.UnusedTime
}