/*
    Response structs for line calls.

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

type Lines struct{
    Ticker string `json:"ticker"`
    PredictedTimeList []string `json:"predictedTimeList"`//2014-10-23 20:04
    StationIDs []string `json:"stationsIDs"`
    LineIDs []string `json:"lineIDs"`//Bad Dürkheim Bahnhof
    TimeList []string `json:"timeList"`//2014-10-23 20:04+0
    ValidFromIndex int `json:"validFromIndex"`//-1
}


//Line Object
type Line struct{
    LineID string `json:"lineID"`
    LineType string `json:"lineType"`
    Hexcolor string `json:"hecolor"`
    IconName string `json:"iconName"`
}
type CanceledLine struct{
    LineID string `json:"lineId"`
    ValidFrom string `json:"validFrom"`
    ValidTo string `json:"validTo"`
    OperationDay string `json:"operationDay"`
    canceled bool `json:"canceled"`
    deleted bool `json:"deleted"`
}
//Used for stationMonitor and Line Package
// LineJourney Object
type Departures struct {
    Time string `json:"time"`
    Status string `json:"status"`
    Direction string `json:"direction"`
    Platform string `json:"platform"`
    Transportation string `json:"transportation"`
    TourID string `json:"tourId"`
    KindOfTour string `json:"kindOfTour"`
    PositionInTour string `json:"positionInTour"`
    StatusNote string `json:"statusNote"`
    LineID string `json:"lineId"`
    LineLabel string `json:"lineLabel"`
    DifferenceTime string `json:"differenceTime"`
    ForeignLine string `json:"foreignLine"`
    NewsAvailable string `json:"newsAvailable"`
}