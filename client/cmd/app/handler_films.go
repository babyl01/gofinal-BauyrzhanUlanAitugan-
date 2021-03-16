package main

import (
	"awesomeProject2/pb/proto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"

	"log"
	"net/http"
	"os"
)

func (app *App) getFilmList(w http.ResponseWriter, r *http.Request) {
	cc, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := proto.NewFilmServiceClient(cc)

	req := &proto.ListFilmReq{}

	stream, err := c.ListFilms(context.Background(), req)
	if err != nil {
		log.Fatalf("error rpc: %v", err)
	}
	var films []*proto.Film
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		films = append(films, res.GetFilm())
	}

	json.NewEncoder(w).Encode(films)
}

func (app *App) getFilm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	url := fmt.Sprintf("%s%s%s", app.apis.films, "id=", id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Fprint(w, string(contents))

}

func (app *App) setFilm(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	price := r.URL.Query().Get("price")

	url := fmt.Sprintf("%s%s", app.apis.films, fmt.Sprintf("add?id=%s&name=%s&price=%s", id, name, price))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Fprint(w, string(contents))
}
