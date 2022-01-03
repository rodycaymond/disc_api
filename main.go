package main

import (
    "example.com/services"
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Welcome to the Disc Golf Api")
    r := gin.Default()
    r.GET("/all", func(c *gin.Context){
      c.JSON(200, services.AllDiscs())
    })
    r.GET("/very-understable", func(c *gin.Context){
        res := services.VeryUnderstable()
        c.JSON(200, res)
    } )
    r.GET("/understable", func(c *gin.Context){
      res := services.Understable()
      fmt.Println("in main layer")
      fmt.Println(res)
      c.JSON(200, res)
    } )
    r.GET("/stable", func(c *gin.Context){
      res := services.Stable()
      fmt.Println("in main layer")
      fmt.Println(res)
      c.JSON(200, res)
    } )
    r.GET("/overstable", func(c *gin.Context){
      res := services.OverStable()
      fmt.Println("in main layer")
      fmt.Println(res)
      c.JSON(200, res)
    } )
    r.GET("/very-overstable", func(c *gin.Context){
      res := services.VeryOverStable()
      fmt.Println("in main layer")
      fmt.Println(res)
      c.JSON(200, res)
    } )
    r.GET("/putters", func(c *gin.Context){
        res := services.Putters()
        c.JSON(200, res)
    })
    r.GET("/brand/:brandName", func(c *gin.Context){
      brandName := c.Param("brandName")
      res := services.ByBrand(brandName)
      c.JSON(200, res)
    })
    r.GET("putter/brand/:brandName", func(c *gin.Context){
        brandName := c.Param("brandName")
        res := services.ByPutterBrand(brandName)
        c.JSON(200, res)
    })
    r.GET("/disc-name/:name", func(c *gin.Context){
        discName := c.Param("name")
        res := services.ByName(discName)
        c.JSON(200, res)
    })
    r.GET("putter/disc-name/:name", func(c *gin.Context){
        discName := c.Param("name")
        res := services.ByPutterName(discName)
        c.JSON(200, res)
    })
    r.Run()
}