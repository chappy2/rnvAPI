rnvAPI
=======

Go Implemenation for accessing the rnv Start.Info API. 

Full documentation and API token available on 
[OpenData Portal Rhein-Neckar-Verkehr GmbH](https://opendata.rnv-online.de/startinfo-api)
Usage
--------------
Get an API  [token](https://opendata.rnv-online.de/startinfo-api) and copy it in the tokenConf-example.json. Rename File to tokenConf.json.

Example
---
A few code lines from apitest.go:
```go
import (
	 rnvrequest "rnvApi/Request"
    "log"
   	//... 
)
func main(){
	//... 
    stationMonitorDocuExamplesTest()
    //... 
}
//... 
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

//... 
```
The output:
```
2016/04/14 11:05:37 REQ: http://rnv.the-agent-factory.de:8080/easygo2/api/regions/rnv/modules/stationmonitor/element?hafasID=2077&mode=DEP&time=2016-04-14+11:05
2016/04/14 11:05:37 RES: 200 OK HTTP/1.1
2016/04/14 11:05:37 &{12:30+0 OK Rheingönheim n/a KOM 3089122418 454AUS 2  RNV075_RNV075A 75 9 false false}
2016/04/14 11:05:37 &{12:50+0 OK Rheingönheim n/a KOM 3073122418 454AUS 2  RNV075_RNV075A 75 29 false false}
2016/04/14 11:05:37 &{13:10 OK Rheingönheim n/a KOM 3070422418 454REFAUS 2  RNV075_RNV075A 75 49 false false}
2016/04/14 11:05:37 &{13:30 OK Rheingönheim n/a KOM 3051222418 454REFAUS 2  RNV075_RNV075A 75 69 false false}
2016/04/14 11:05:37 &{13:50 OK Rheingönheim n/a KOM 3096422418 454REFAUS 2  RNV075_RNV075A 75 89 false false}
2016/04/14 11:05:37 &{14:10 OK Rheingönheim n/a KOM 3078422418 454REFAUS 2  RNV075_RNV075A 75 109 false false}
2016/04/14 11:05:37 &{14:30 OK Rheingönheim n/a KOM 3100022418 454REFAUS 2  RNV075_RNV075A 75 129 false false}
2016/04/14 11:05:37 &{14:50 OK Rheingönheim n/a KOM 3073522418 454REFAUS 2  RNV075_RNV075A 75 149 false false}
2016/04/14 11:05:37 &{15:10 OK Rheingönheim n/a KOM 3076522418 454REFAUS 2  RNV075_RNV075A 75 169 false false}
2016/04/14 11:05:37 &{15:20 OK Rheingönheim Bahnhof n/a KOM 2673220718 454REFAUS 2  RNV075_RNV075A 75 179 false false}
2016/04/14 11:05:37 REQ: http://rnv.the-agent-factory.de:8080/easygo2/api/regions/rnv/modules/stationmonitor/element?hafasID=2077&time=null
2016/04/14 11:05:37 RES: 200 OK HTTP/1.1
2016/04/14 11:05:37 &{12:30+0 OK Rheingönheim n/a KOM 3089122418 454AUS 2  RNV075_RNV075A 75 9 false false}
2016/04/14 11:05:37 &{12:50+0 OK Rheingönheim n/a KOM 3073122418 454AUS 2  RNV075_RNV075A 75 29 false false}
2016/04/14 11:05:37 &{13:10 OK Rheingönheim n/a KOM 3070422418 454REFAUS 2  RNV075_RNV075A 75 49 false false}
2016/04/14 11:05:37 &{13:30 OK Rheingönheim n/a KOM 3051222418 454REFAUS 2  RNV075_RNV075A 75 69 false false}
2016/04/14 11:05:37 &{13:50 OK Rheingönheim n/a KOM 3096422418 454REFAUS 2  RNV075_RNV075A 75 89 false false}
2016/04/14 11:05:37 &{14:10 OK Rheingönheim n/a KOM 3078422418 454REFAUS 2  RNV075_RNV075A 75 109 false false}
2016/04/14 11:05:37 &{14:30 OK Rheingönheim n/a KOM 3100022418 454REFAUS 2  RNV075_RNV075A 75 129 false false}
2016/04/14 11:05:37 &{14:50 OK Rheingönheim n/a KOM 3073522418 454REFAUS 2  RNV075_RNV075A 75 149 false false}
2016/04/14 11:05:37 &{15:10 OK Rheingönheim n/a KOM 3076522418 454REFAUS 2  RNV075_RNV075A 75 169 false false}
2016/04/14 11:05:37 &{15:20 OK Rheingönheim Bahnhof n/a KOM 2673220718 454REFAUS 2  RNV075_RNV075A 75 179 false false}
2016/04/14 11:05:37 REQ: http://rnv.the-agent-factory.de:8080/easygo2/api/regions/rnv/modules/stationmonitor/element?hafasID=1160&needPlatformDetail=true&poles=2%3B3&time=2016-04-14+11:05
2016/04/14 11:05:38 RES: 200 OK HTTP/1.1
2016/04/14 11:05:38 &{12:22+0 OK Unikl. - Neuenh. Feld n/a KOM 2880954101 454AUS 7  RNV032_RNV032A 32 2 false false}
2016/04/14 11:05:38 &{12:26+0 OK Altstadt Universitätsplatz n/a KOM 2887854101 454AUS 10  RNV032_RNV032B 32 5 false false}
2016/04/14 11:05:38 &{12:25+1 OK Ziegelhausen Köpfel n/a KOM 2852354101 454AUS 30  RNV033_RNV033B 33 5 false false}
2016/04/14 11:05:38 &{12:24+2 OK Bismarckplatz n/a STRAB 3021444200 454AUS 6  RNV021_RNV021A 21 6 false false}
2016/04/14 11:05:38 &{12:27+0 OK Mannheim-Weinheim n/a STRAB 3425733551 454AUS 25  RNV005_RNV005B 5 6 false true}
2016/04/14 11:05:38 &{12:28+0 OK Weinheim-Mannheim n/a STRAB 3402433651 454AUS 10  RNV005_RNV005A 5 8 false true}
2016/04/14 11:05:38 &{12:29+0 OK Rohrbach Süd n/a STRAB 2925744200 454AUS 8  RNV024_RNV024A 24 8 false false}
2016/04/14 11:05:38 &{12:29+0 OK Ziegelh. Heidebuckelweg n/a KOM 2548655101 454AUS 25  RNV034_RNV034A 34 9 false false}
2016/04/14 11:05:38 &{12:30+0 OK Hans-Thoma-Platz n/a STRAB 3014344200 454AUS 4  RNV021_RNV021B 21 10 false false}
2016/04/14 11:05:38 &{12:27+5 OK Handschuhsheim Nord n/a STRAB 2941444200 454AUS 13  RNV024_RNV024B 24 12 false false}

```


Other Modules
--------------
Currently the following solutions have no token support. The token is necessary since April 2016.

[rnv.js](https://github.com/silsha/rnv.js) (Node.js; by silsha)
silsha also wrote a documentation in [his wiki](https://github.com/silsha/rnv-api/wiki).

[rnv-python](https://github.com/Don42/rnv-python) (Python; by Don42)

License
-------
GPL v2, see LICENSE
