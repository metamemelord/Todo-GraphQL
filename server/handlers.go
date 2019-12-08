package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"todo-graph/graphql"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handlerSignature struct {
	path    string
	method  string
	handler func(*gin.Context)
}

func getGraphGetHandler(log *logrus.Logger) func(*gin.Context) {
	return func(g *gin.Context) {
		query := g.Query("query")
		if query == "" {
			g.AbortWithError(400, errors.New("Request query is empty"))
		}
		result := graphql.ExecuteQuery(query)
		if !result.HasErrors() {
			g.Header("content-type", "application/json")
			err := json.NewEncoder(g.Writer).Encode(result.Data)
			if err != nil {
				log.Error(err)
			}
		} else {
			errs := []error{}
			for _, err := range result.Errors {
				errs = append(errs, err.OriginalError())
			}
			g.AbortWithStatusJSON(400, errs)
		}
	}
}

func getGraphPostHandler(log *logrus.Logger) func(*gin.Context) {
	return func(g *gin.Context) {
		query := g.Query("query")
		if query == "" {
			q, err := ioutil.ReadAll(g.Request.Body)
			if err != nil {
				log.Error("Error while parsing body")
				return
			}
			query = string(q)
		}
		if query == "" {
			g.AbortWithError(400, errors.New("Request query is empty"))
		}
		result := graphql.ExecuteQuery(query)
		if !result.HasErrors() {
			g.Header("content-type", "application/json")
			err := json.NewEncoder(g.Writer).Encode(result.Data)
			if err != nil {
				log.Error(err)
			}
		} else {
			errs := []error{}
			for _, err := range result.Errors {
				errs = append(errs, err.OriginalError())
			}
			g.AbortWithStatusJSON(400, errs)
		}
	}
}

func GetAllHandlers(l *logrus.Logger) []*handlerSignature {
	return []*handlerSignature{
		&handlerSignature{"/graphql", "GET", getGraphGetHandler(l)},
		&handlerSignature{"/graphql", "POST", getGraphPostHandler(l)},
	}
}
