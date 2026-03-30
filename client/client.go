package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetStudent(ctx, &pb.StudentRequest{
		Id: 101,
	})

	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	resList, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	log.Println("Student List:")
	for _, s := range resList.Student {
		log.Printf("ID: %d | Name: %s | Major: %s | Email: %s | Phone: %s",
			s.Id, s.Name, s.Major, s.Email, s.Phone)
	}

	log.Printf("Student Info:")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
	log.Printf("Phone: %s", res.Phone)
}
