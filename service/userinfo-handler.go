package service

import (
    "net/http"
    "strconv"

    "MySql/entities"

    "github.com/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        if len(req.Form["username"][0]) == 0 {
            formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
            return
        }
        u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
        u.DepartName = req.Form["departname"][0]
        entities.Save(u)
        formatter.JSON(w, http.StatusOK, u)
    }
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        if len(req.Form["userid"][0]) != 0 {
            i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

            u := entities.FindByID(int(i))
            formatter.JSON(w, http.StatusBadRequest, u)
            return
        }
        ulist := entities.FindAll()
        formatter.JSON(w, http.StatusOK, ulist)
    }
}