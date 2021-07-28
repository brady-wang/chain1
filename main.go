package main

import (
	"fmt"
	"github.com/brady-wang/gee/gee"
	"log"
	"net/http"
	"strconv"
	"test3/Block"
	"test3/BlockChain"
)
var first *Block.Block

var preBlock *Block.Block

func main() {
	go func() {
		first = Block.CreateFirstBlock()
		preBlock = first
		BlockChain.AddChain(first)
		fmt.Printf("%#v\n",first)
	}()

	r := gee.New()
	r.GET("/chain",chainList)
	r.GET("/next",nextChain)
	log.Fatal(r.Run(":8080"))
}

func nextChain(c *gee.Context) {
	var bmpStr = c.Query("bmp")
	bmp,_ := strconv.Atoi(bmpStr)
	next := Block.NextBlock(bmp,*preBlock)
	BlockChain.AddChain(next)
	preBlock = next
	c.JSON(http.StatusOK,next)

}

func chainList(c *gee.Context) {
	chain := BlockChain.Show()
	fmt.Printf("区块长度%d",len(chain))
	c.JSON(http.StatusOK,chain)
}



