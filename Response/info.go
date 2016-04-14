/*
    Response structs for info calls.

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
package Response

type StationInfo struct{
    Title string `json:"title"`
    Text string `json:"text"`
    ID int `json:"id"`
    LineID string `json:"lineId"`
    StationIds []string `json:"stationIds"`
    StationNames []string `json:"stationNames"`    
    Url string `json:"url"`    
    Author string `json:"author"`
    Created int64 `json:"created"`
    ValidFrom int64 `json:"validFrom"`
    ValidTo int64 `json:"validTo"`
    DisplayFrom int64 `json:"displayFrom"`
    DisplayTo int64 `json:"displayTo"`   
    
}

type LineInfo struct{
    Title string `json:"title"`
    Text string `json:"text"`    
    ID int `json:"id"`
    StationID string `json:"stationId"`
    StationName string `json:"stationName"`
    Poles []string `json:"poles"`    
    Author string `json:"author"`
    Created int64 `json:"created"`
    ValidFrom int64 `json:"validFrom"`
    ValidTo int64 `json:"validTo"`  
    
}