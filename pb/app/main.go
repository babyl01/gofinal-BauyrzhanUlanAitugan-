package main

import (
	"awesomeProject2/films/pkg/models"
	"awesomeProject2/pb/proto"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"log"
	"net"
	"os"
	"os/signal"
)

type FilmServiceServer struct {
	proto.UnimplementedFilmServiceServer
}

var database *sql.DB

func main() {
	listener, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("Unable to listen on port :4000: %v", err)
	}

	s := grpc.NewServer()
	srv := &FilmServiceServer{}
	proto.RegisterFilmServiceServer(s, srv)

	dsn := flag.String("dsn", "root:root@/cinema?parseTime=true", "MySQL DSN")

	flag.Parse()

	database = connect(*dsn)

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Println("Server successfully started on port :4000")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)
	<-c
	s.Stop()
	listener.Close()

	defer database.Close()

}

func (s *FilmServiceServer) ListFilms(req *proto.ListFilmReq, stream proto.FilmService_ListFilmsServer) error {
	stmt := `SELECT id, name, price FROM films`

	rows, err := database.Query(stmt)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		data := &models.Film{}
		err := rows.Scan(&data.ID, &data.Name, &data.Price)

		stream.Send(&proto.ListFilmRes{
			Film: &proto.Film{
				Id:    int64(data.ID),
				Name:  data.Name,
				Price: int64(data.Price),
			},
		})

		if err != nil {
			return err
		}

	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
