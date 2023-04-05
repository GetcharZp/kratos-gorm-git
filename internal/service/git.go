package service

import (
	"github.com/asim/git-http-backend/server"
	"kratos-gorm-git/helper"
	"kratos-gorm-git/models"
	"net/http"
	"os"
	"strings"
)

func GitHttpBackend(w http.ResponseWriter, r *http.Request) {
	// 1. 获取仓库的基础信息
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/git")
	repoPaths := strings.Split(r.URL.Path, "/")
	if len(repoPaths) < 3 {
		w.Write([]byte("error request"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	repoPath := repoPaths[1] + string(os.PathSeparator) + repoPaths[2]
	rb := new(models.RepoBasic)
	ub := new(models.UserBasic)
	err := models.DB.Model(new(models.RepoBasic)).Where("path = ?", repoPath).First(rb).Error
	if err != nil {
		w.Write([]byte("sys error"))
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	// 2. clone || push
	if repoPaths[3] == "git-upload-pack" || repoPaths[3] == "git-receive-pack" {
		if rb.Type == 0 {
			// 获取用户的基础信息
			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.Header().Set("WWW-Authenticate", `Basic realm="Git Repository"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			username, password, _ := r.BasicAuth()
			err = models.DB.Model(new(models.UserBasic)).Where("username = ? AND password = ?", username, helper.GetMd5(password)).First(ub).Error
			if err != nil {
				w.Write([]byte("sys error"))
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			// 获取授权信息
			var cnt int64
			err = models.DB.Model(new(models.RepoUser)).Where("rid = ? AND uid = ?", rb.ID, ub.ID).Count(&cnt).Error
			if err != nil {
				w.Write([]byte("sys error"))
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			if cnt == 0 {
				w.Write([]byte("no auth"))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
	}
	server.Handler().ServeHTTP(w, r)
}
