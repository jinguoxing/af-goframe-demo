package greeter

import (

    "af-goframe-demo/dddDemo/domain/greeter"
    "github.com/gin-gonic/gin"
    "net/http"
)

type GreeterService struct {

    uc *greeter.GreeterUsecase

}


func NewGreeterService(uc *greeter.GreeterUsecase) *GreeterService{

    return &GreeterService{uc:uc}

}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(c *gin.Context)  {
    foo := c.DefaultQuery("foo","def")
    g, err := s.uc.CreateGreeter(c, &greeter.Greeter{Hello: foo})

    if err != nil {
       s := http.StatusInternalServerError
       c.JSON(s,gin.H{g})
    }else{
       s := http.StatusOK
       c.JSON(s,gin.H{g})
    }

}
