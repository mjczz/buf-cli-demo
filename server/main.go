package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	connect "connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	petv1 "github.com/mjczz/buf-cli-demo/gen/pet/v1"
	"github.com/mjczz/buf-cli-demo/gen/pet/v1/petv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := petv1connect.NewPetStoreServiceHandler(&petStoreServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

// petStoreServiceServer implements the PetStoreService API.
type petStoreServiceServer struct {
	petv1connect.UnimplementedPetStoreServiceHandler
}

func (s *petStoreServiceServer) GetPet(
	ctx context.Context,
	req *connect.Request[petv1.GetPetRequest],
) (*connect.Response[petv1.GetPetResponse], error) {
	petId := req.Msg.GetPetId()
	if err := protovalidate.Validate(req.Msg); err != nil {
		fmt.Println("validation failed:", err)
		return &connect.Response[petv1.GetPetResponse]{}, err
	}
	log.Printf("Got a request to get a %v", petId)
	return connect.NewResponse(&petv1.GetPetResponse{}), nil
}

func (s *petStoreServiceServer) DeletePet(
	ctx context.Context,
	req *connect.Request[petv1.DeletePetRequest],
) (*connect.Response[petv1.DeletePetResponse], error) {
	petId := req.Msg.GetPetId()
	if err := protovalidate.Validate(req.Msg); err != nil {
		fmt.Println("validation failed:", err)
		return &connect.Response[petv1.DeletePetResponse]{}, err
	}
	log.Printf("Got a request to delete a %v", petId)
	return connect.NewResponse(&petv1.DeletePetResponse{}), nil
}

// PutPet adds the pet associated with the given request into the PetStore.
func (s *petStoreServiceServer) PutPet(
	ctx context.Context,
	req *connect.Request[petv1.PutPetRequest],
) (*connect.Response[petv1.PutPetResponse], error) {
	name := req.Msg.GetName()
	petType := req.Msg.GetPetType()
	if err := protovalidate.Validate(req.Msg); err != nil {
		fmt.Println("validation failed:", err)
		return &connect.Response[petv1.PutPetResponse]{}, err
	}
	log.Printf("Got a request to create a %v named %s", petType, name)
	return connect.NewResponse(&petv1.PutPetResponse{}), nil
}
