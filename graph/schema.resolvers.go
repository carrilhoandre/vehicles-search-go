package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/carrilhoandre/webmotors-search-go/graph/generated"
	"github.com/carrilhoandre/webmotors-search-go/graph/model"
	elastic "github.com/olivere/elastic/v7"
)

func (r *mutationResolver) CreateMake(ctx context.Context, input model.NewQuery) (*model.Vehicle, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Vehicles(ctx context.Context, text *string) ([]*model.Vehicle, error) {
	var vehicles []*model.Vehicle
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("[esclient] Cant connect", err)
	}
	searchSource := elastic.NewSearchSource()
	q := elastic.NewMoreLikeThisQuery().LikeText(*text).Field("MakeName", "ModelName").MinTermFreq(1).MinDocFreq(1)
	searchSource.Query(q)

	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2, queryJs)
	}
	//fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	searchService := esclient.Search().Index("vehicle").SearchSource(searchSource)
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return vehicles, nil
	}

	for _, hit := range searchResult.Hits.Hits {
		var vehicle model.Vehicle
		err := json.Unmarshal(hit.Source, &vehicle)
		if err != nil {
			fmt.Println("[Getting Vehicles][Unmarshal] Err=", err)
		}

		vehicles = append(vehicles, &model.Vehicle{MakeName: vehicle.MakeName, ModelName: vehicle.ModelName})
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	}
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	return vehicles, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	return client, err

}
