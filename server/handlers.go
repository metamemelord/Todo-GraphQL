package server

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type handlerSignature struct {
	path    string
	method  string
	handler func(*gin.Context)
}

func getGraphGetHandler(log *log.Logger) func(*gin.Context) {
	log.Println("Regitering a GET handler...")
	return func(g *gin.Context) {
		log.Println("YAIIIIIIIIIIIIIII")
		query := g.Query("query")
		if query == "" {
			g.AbortWithError(400, errors.New("Request query is empty"))
		}
	}
}

func getGraphPostHandler(log *log.Logger) func(*gin.Context) {
	log.Println("Regitering a POST handler...")
	return func(g *gin.Context) {
		query := g.Query("query")
		if query == "" {
			q, err := ioutil.ReadAll(g.Request.Body)
			if err != nil {
				log.Println("Error while parsing body")
				return
			}
			query = string(q)
		}
		if query == "" {
			g.AbortWithError(400, errors.New("Request query is empty"))
		}
	}
}

func GetAllHandlers(l *log.Logger) []*handlerSignature {
	return []*handlerSignature{
		&handlerSignature{"/graphql", "GET", getGraphGetHandler(l)},
		&handlerSignature{"/graphql", "POST", getGraphPostHandler(l)},
	}
}
