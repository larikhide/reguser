package routergin

import (
	"fmt"
	"net/http"

	"github.com/larikhide/reguser/api/handler"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/google/uuid"
)

type RouterGin struct {
	*gin.Engine
	hs *handler.Handlers
}

func GinAuthMW(c *gin.Context) {
	if u, p, ok := c.Request.BasicAuth(); !ok || !(u == "admin" && p == "admin") {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unautorized"))
		return
	}
	c.Next()
}

func NewRouterGin(hs *handler.Handlers) *RouterGin {
	r := gin.Default()
	ret := &RouterGin{
		hs: hs,
	}

	r.Use(GinAuthMW)

	r.POST("/create", ret.CreateUser)
	r.GET("/read/:id", ret.ReadUser)
	r.DELETE("/delete/:id", ret.DeleteUser)
	r.GET("/search/:q", ret.SearchUser)

	ret.Engine = r
	return ret
}

type User handler.User

func (rt *RouterGin) CreateUser(c *gin.Context) {
	ru := User{}
	if err := c.ShouldBindJSON(&ru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := rt.hs.CreateUser(c.Request.Context(), handler.User(ru))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (rt *RouterGin) ReadUser(c *gin.Context) {
	sid := c.Param("id")

	uid, err := uuid.Parse(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := rt.hs.ReadUser(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (rt *RouterGin) DeleteUser(c *gin.Context) {
	sid := c.Param("id")

	uid, err := uuid.Parse(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := rt.hs.DeleteUser(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (rt *RouterGin) SearchUser(c *gin.Context) {
	q := c.Param("id")
	w := c.Writer
	fmt.Fprintln(w, "[")
	comma := false
	err := rt.hs.SearchUser(c.Request.Context(), q, func(u handler.User) error {
		if comma {
			fmt.Fprintln(w, ",")
		} else {
			comma = true
		}
		(render.JSON{Data: u}).Render(w)
		w.Flush()
		return nil
	})
	if err != nil {
		if comma {
			fmt.Fprint(w, ",")
		}
		(render.JSON{Data: err}).Render(w)
	}
	fmt.Fprintln(w, "]")
}
