/*
    Response structs for news and ticker calls.

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

type News struct{
    RegionID int `json:"regionID"`
    Title string `json:"title"`
    ValidFrom int64 `json:"validFrom"`
    dateAsString string `json:"tdateAsString"`
    ValidTo int64 `json:"validTo"`
    Lines string `json:"lines"`
    seperatedLines []string `json:"seperatedLines"`
    ImgUrl string `json:"imgUrl"`
    TextAsHtml string `json:"textAsHtml"`
    IsOldNews bool `json:"idOldNews"`
    ThumbUrl string `json:"thumbUrl"`
    FurtherLink string `json:"furtherLink"`
    ElementID int `json:"elmentID"`
}