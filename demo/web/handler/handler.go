package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yydsorg/goyyds/tool/env"
	"log"
	"net/http"
)

type GetInfoReq struct {
	Id uint64 `json:"id" form:"id" binding:"required"`
}

type GetInfoRsp struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

func GetInfo(c *gin.Context) {
	req := new(GetInfoReq)
	rsp := new(GetInfoRsp)

	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	fail := func(code int64, msg string) {
		rsp.Message = msg
		rsp.Code = -1
	}

	if err := c.ShouldBind(req); err != nil {
		fail(-1, "参数错误")
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		fail(-1, "获取文件失败")
		return
	}

	if file != nil {
		if err := c.SaveUploadedFile(file, env.GetPublicPath()); err != nil {
			fail(-1, "fail")
			return
		}
	}

	log.Println(req.Id)

	rsp.Message = "success"
	rsp.Code = 200

}
