package greeter

import (
	"af-go-frame/core/transport/rest"
	"af-goframe-demo/dddDemo/domain/greeter"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GreeterService struct {
	uc     *greeter.GreeterUsecase
	logger *zap.Logger
}

func NewGreeterService(uc *greeter.GreeterUsecase, logger *zap.Logger) *GreeterService {

	return &GreeterService{uc: uc, logger: logger}

}

// SayHello godoc
// @Summary      sayhello example
// @Description  do sayhello
// @Tags         Greeter
// @Accept       json
// @Produce      json
// @Success      200  {string} json
// @Failure      400  {string} string
// @Failure      404  {string} string
// @Failure      500  {string} string
// @Router       /greeter/sayhello [get]
func (s *GreeterService) SayHello(c *gin.Context) {
	foo := c.DefaultQuery("foo", "def")
	g, err := s.uc.CreateGreeter(c, &greeter.Greeter{Hello: foo})

	if err != nil {
		//res := http.StatusInternalServerError
		s.logger.Error("controller/greeter/greeter.go line 30")

		rest.ResErrJson(c,err,"",gin.H{"rep": "err"})
	} else {

		rest.ResOkJson(c,"",gin.H{"rep": g.Hello})
	}
}

type greeterExample struct {
	Name string `json:"greeter_name"`
	ID   string `json:"greeter_id"`
}

var examples []*greeterExample

// 提交 列表

// PostExample godoc
// @Summary      PostExample Summary
// @Description  PostExample Description
// @Accept       json
// @Produce      json
// @Tags         Greeter1
// @Param        message  body      greeterExample  true  "Greeter Info"
// @Success      200      {string}  string         "success"
// @Failure      500      {string}  string         "fail"
// @Router       /greeter/post [post]
func (s *GreeterService) PostExample(c *gin.Context) {
	var e = &greeterExample{}

	if err := c.ShouldBind(e); err != nil {
		fmt.Println(e)
		return
	}
	//s.uc.CreateGreeter
	examples = append(examples, e)
	fmt.Println(examples)
}

// ListExample godoc
// @Summary      ListExample Summary
// @Description  ListExample Description
// @Accept       json
// @Produce      json
// @Tags         Greeter2
// @Success      200      {array}  greeterExample
// @Failure      500      {string}  string         "fail"
// @Router       /greeter/list [get]
func (s *GreeterService) ListExample(c *gin.Context) {
	c.JSON(200, examples)
}
