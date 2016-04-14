/*
    A few test calls.

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
package main

import (
  
    rnvrequest "rnvApi/Request"
    con "strconv"
    "log"
    "time"
)


func main(){
	
    t1 := time.Now()
	 
    updateTest()
    log.Println(testIfHafasIDIsEqualToElementID())
    
    linePackageTest()
    lineInfoTest()
    stationMonitorDocuExamplesTest()
    newsTest()
    canceledLineTest()
    infoTest()
    mapTest()
    t2 := time.Now()
    log.Printf("[ApiTest] Time needed: %v\n", t2.Sub(t1))
}
func mapTest(){
    for _,v:=range rnvrequest.SupportedThumbnailSizes{
        for _,f:=range rnvrequest.SupportedFormats{
            maps,err:=rnvrequest.GetMaps(rnvrequest.MapsParams{&v,&f})
             if err!=nil{
                log.Println(err)
                return
            }
            log.Println(*maps)
            for _,in := range *maps{
                log.Println(in)
            }
        }
    }
}
func infoTest(){
    info,err:=rnvrequest.GetStationInfo(rnvrequest.StationInfoParams{nil,nil,nil})
    if err!=nil{
        log.Println(err)
        return
    }
    for _,in := range *info{
        log.Println(in)
    }
    var Lines []string
    Lines=append(Lines,"2")
    HafasID :="2321"
    info3,err3:=rnvrequest.GetStationInfo(rnvrequest.StationInfoParams{&Lines,&HafasID,nil})
    if err3!=nil{
        log.Println(err3)
        return
    }
    for _,in := range *info3{
        log.Println(in)
    }
    info2,err2:=rnvrequest.GetLineInfo(rnvrequest.LineInfoParams{nil,nil,nil})
    if err2!=nil{
        log.Println(err2)
        return
    }
    for _,in := range *info2{
        log.Println(in)
    }
    HafasID2 :="2386"
    var Poles []string
    Poles=append(Lines,"7")
    info4,err4:=rnvrequest.GetLineInfo(rnvrequest.LineInfoParams{&HafasID2,&Poles,nil})
    if err4!=nil{
        log.Println(err4)
        return
    }
    for _,in := range *info4{
        log.Println(in)
    }
}
func canceledLineTest(){
     lines,err:=rnvrequest.GetLinesPackage()
    if err!=nil{
        log.Println(err)
        return
    }
    
    
    for _,line:=range lines{
        news,err:=rnvrequest.GetCanceledLine(rnvrequest.CanceledLineParams{&line.LineID,rnvrequest.GetTimeUnixNow()})
    if err!=nil{
        log.Println(err)
        return
    }
    log.Println(news)
    }
}
func newsTest(){
    //No news available on Testserver..
    news,err:=rnvrequest.GetNews()
    if err!=nil{
        log.Println(err)
        return
    }
    log.Println(news)
    for _,n:= range news{
        log.Println(n)
    }
    var params []string
    params=append(params,"1")
    params=append(params,"2")    
    params=append(params,"5")
    params=append(params,"8")
    params=append(params,"76")
    news2,err2:=rnvrequest.GetTickerForLines(params)
    if err2!=nil{
        log.Println(err2)
        return
    }
    log.Println(news2)
    for _,n:= range news2{
        log.Println(n)
    }
}

func stationMonitorDocuExamplesTest(){
    id2077:="2077"
    id1160:="1160"
    nullTime:="null"
    dep:="DEP"
    poles:="2;3"
    needTrue:="true"
    
   time:=rnvrequest.GetTimeFormatNow()
    test1,err1:=rnvrequest.GetStationMonitor(rnvrequest.StationMonitorParams{&id2077,time,nil,&dep,nil})
    if err1!=nil{
        log.Println(err1)
        return
    }
    for _,line:=range test1.ListOfDepartures{
            log.Println(line)
    }
    test2,err2:=rnvrequest.GetStationMonitor(rnvrequest.StationMonitorParams{&id2077,&nullTime,nil,nil,nil})
    if err2!=nil{
        log.Println(err2)
        return
    }
    for _,line:=range test2.ListOfDepartures{
            log.Println(line)
    }
    test3,err3:=rnvrequest.GetStationMonitor(rnvrequest.StationMonitorParams{&id1160,time,&poles,nil,&needTrue})
    if err3!=nil{
        log.Println(err3)
        return
    }
    for _,line:=range test3.ListOfDepartures{
            log.Println(line)
    }
}
func lineInfoTest(){
    lines,err:=rnvrequest.GeltLineInfos()
    if err!=nil{
        log.Println(err)
        return
    }
    for _,line:=range lines{
        log.Println(line)
    }
}
func linePackageTest(){
    lines,err:=rnvrequest.GetLinesPackage()
    if err!=nil{
        log.Println(err)
        return
    }
    for _,line:=range lines{
        log.Println(line)
    }
}

func updateTest(){
    //Todo create time object and get TimeFormat
    time:="2013-07-08+13:40" 
    updates,err:=rnvrequest.GetUpdates(rnvrequest.UpdateParams{"1",&time,&time,&time})
	if err!=nil{
        log.Println(err)
        return
    }
    log.Println(updates)
    time2:=rnvrequest.GetTimeFormatNow()
    updates2,err2:=rnvrequest.GetUpdates(rnvrequest.UpdateParams{"1",time2,time2,time2})
	if err2!=nil{
        log.Println(err2)
        return
    }
    log.Println(updates2)
    
}

func testIfHafasIDIsEqualToElementID()bool{
	
    stationsBody,err:=rnvrequest.GetStations("1")//On Testserver is only one Stationspackage
    result:=true
	if err!=nil{
        log.Println(err)
        return true
    }
    nullTime:="null"
    for _,station := range stationsBody.Stations{  
        hafas,_:=con.Atoi(station.HafasID)
        if(station.ElementID!=hafas){
            log.Println( station.HafasID)
            log.Println(station.ElementID)
            result=false
        }
        stationsMonitorBody,_:=rnvrequest.GetStationMonitor(rnvrequest.StationMonitorParams{&station.HafasID,&nullTime,nil,nil,nil})
        
        log.Println("******************************************************************************************")
        log.Println(stationsMonitorBody.Ticker)
        for _,line:=range stationsMonitorBody.ListOfDepartures{
            log.Println(line)
            var paramsfive string="5"
            linesMonitor,_:=rnvrequest.GetLines(rnvrequest.LinesParams{&station.HafasID,&line.LineLabel,&paramsfive,&line.KindOfTour,&line.TourID,&line.Time})
            log.Println("_____________________________________________________________________________________")
            log.Println(linesMonitor)
        }
    }
    return result
}


