package middlware

import (
	"avitotestgo2024/internal/auth"
	"avitotestgo2024/internal/entitys"
	"avitotestgo2024/internal/service"
	"io"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Service service.ServiceInterface
	Auth    auth.AuthInterface
	Logger  *slog.Logger
	adres   string
}

func NewServer(host string, port string, auth auth.AuthInterface, service service.ServiceInterface, logger *slog.Logger) *Server {
	s := Server{Service: service, Auth: auth, Logger: logger}
	BannerGet := func(c *gin.Context) {
		if ok := s.CheckPermission(c, 2); !ok {
			return
		}
		var tag_id, feature_id, limit, offset int
		var err error
		tag_id_str := c.Query("tag_id")
		if len(tag_id_str) > 0 {
			tag_id, err = strconv.Atoi(tag_id_str)
			if err != nil || tag_id < 0 {
				c.Status(400)
				c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный tag_id"}))
				return
			}
		} else {
			tag_id = -1
		}
		feature_id_str := c.Query("feature_id")
		feature_id, err = strconv.Atoi(feature_id_str)
		if len(feature_id_str) > 0 {
			if err != nil || feature_id < 0 {
				c.Status(400)
				c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный feature_id"}))
				return
			}
		} else {
			feature_id = -1
		}
		limit_str := c.Query("limit")
		if len(limit_str) > 0 {
			limit, err = strconv.Atoi(limit_str)
			if err != nil || limit <= 0 {
				c.Status(400)
				c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный limit"}))
				return
			}
		} else {
			limit = 0
		}
		offset_str := c.Query("offset")
		if len(offset_str) > 0 {
			offset, err = strconv.Atoi(offset_str)
			if err != nil || offset < 0 {
				c.Status(400)
				c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный offset"}))
				return
			}
		} else {
			offset = 0
		}
		ans, err := s.Service.GetAllBanners(tag_id, feature_id, limit, offset)
		if err != nil {
			c.Status(500)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()}))
			return
		}
		c.Status(200)
		c.Header("Content-Type", "application/json; charset=UTF-8")
		c.Writer.Write(ans)
	}
	BannerIdDelete := func(c *gin.Context) {
		if ok := s.CheckPermission(c, 2); !ok {
			return
		}
		id_str := c.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil || id < 0 {
			c.Status(400)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный id"}))
			return
		}
		err = s.Service.Delete(id)
		if err.Error() == "404" {
			c.Status(404)
			return
		}
		if err != nil {
			c.Status(500)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()}))
			return
		}
		c.Status(204)
	}
	BannerIdPatch := func(c *gin.Context) {
		if ok := s.CheckPermission(c, 2); !ok {
			return
		}
		id_str := c.Param("id")
		id, err := strconv.Atoi(id_str)
		if err != nil || id < 0 {
			c.Status(400)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный id"}))
			return
		}

		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(500)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()}))
			return
		}
		err = s.Service.UpdateBanner(id, data)
		if err != nil && err.Error() == "404" {
			c.Status(404)
			return
		}
		if err != nil {
			c.Status(500)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()}))
			return
		}
		c.Status(200)
	}
	BannerPost := func(c *gin.Context) {
		if ok := s.CheckPermission(c, 2); !ok {
			return
		}
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(500)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()}))
			return
		}
		ans, err := s.Service.CreateBanner(data)
		if err != nil {
			c.Status(500)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()}))
			return
		}
		c.Header("Content-Type", "application/json; charset=UTF-8")
		c.Status(201)
		c.Writer.Write(ans)
	}
	UserBannerGet := func(c *gin.Context) {
		if ok := s.CheckPermission(c, 3); !ok {
			return
		}
		tag_id_str := c.Query("tag_id")
		tag_id, err := strconv.Atoi(tag_id_str)
		if err != nil || tag_id < 0 {
			c.Status(400)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный tag_id"}))
			return
		}
		feature_id_str := c.Query("feature_id")
		feature_id, err := strconv.Atoi(feature_id_str)
		if err != nil || feature_id < 0 {
			c.Status(400)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный feature_id"}))
			return
		}
		use_last_revision_str := c.Query("use_last_revision")
		use_last_revision := false
		switch use_last_revision_str {
		case "true":
			use_last_revision = true
			break
		case "false":
			break
		case "":
			break
		default:
			c.Status(400)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: "не корректный use_last_revision"}))
			return
		}
		data, err := s.Service.GetUserBanner(tag_id, feature_id, use_last_revision, c.Request.Header.Get("token"))
		if err != nil {
			c.Status(500)
			c.Writer.Write(s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()}))
			return
		}
		if len(data) == 0 {
			c.Status(404)
			return
		}
		c.Header("Content-Type", "application/json; charset=UTF-8")
		c.Status(200)
		c.Writer.Write(data)
	}
	ping := func(c *gin.Context) {
		c.Status(200)
	}
	g := gin.New()

	g.GET("/user_banner", UserBannerGet)
	g.GET("/banner", BannerGet)
	g.POST("/banner", BannerPost)
	g.PATCH("/banner/:id", BannerIdPatch)
	g.DELETE("/banner/:id", BannerIdDelete)
	g.GET("/ping", ping)
	s.Router = g
	s.adres = host + ":" + port
	return &s
}

func (s *Server) CheckPermission(c *gin.Context, permission int) bool {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.Status(401)
		return false
	}
	if permission == 3 {
		ok, err := s.Auth.HasPermission(token, 1)
		if !ok && err.Error() == "401" {
			c.Status(401)
			return false
		} else {
			return true
		}
	}
	ok, err := s.Auth.HasPermission(token, permission)
	if ok {
		return true
	}
	if err.Error() == "401" {
		c.Status(401)
		return false
	}
	if err.Error() == "403" {
		c.Status(403)
		return false
	}
	c.Status(500)
	data := s.Service.CovertErrorToBytes(entitys.Error{Message: err.Error()})
	c.Writer.Write(data)
	return false
}

func (s *Server) Run() {
	s.Router.Run(s.adres)
}
