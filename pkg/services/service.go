package services

import (
	"context"
	"net/http"

	"github.com/stebinsabu13/note_taking_microservice/note_srv/pkg/db"
	"github.com/stebinsabu13/note_taking_microservice/note_srv/pkg/models"
	"github.com/stebinsabu13/note_taking_microservice/note_srv/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedNoteServiceServer
}

func (s *Server) ListAllNote(ctx context.Context, req *pb.ListAllNoteRequest) (*pb.ListAllNoteResponse, error) {
	var notes []models.Note
	if err := s.H.DB.Model(&models.Note{}).Where("userid=?", req.Id).Scan(&notes).Error; err != nil {
		return &pb.ListAllNoteResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	var res []*pb.Notes
	for i := 0; i < len(notes); i++ {
		data := &pb.Notes{
			Id:   uint32(notes[i].Id),
			Note: notes[i].Note,
		}
		res = append(res, data)
	}
	return &pb.ListAllNoteResponse{
		Status: http.StatusOK,
		Notes:  res,
	}, nil
}

func (s *Server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	var note models.Note
	note.Userid = req.Uid
	note.Note = req.Note
	if err := s.H.DB.Create(&note).Error; err != nil {
		return &pb.CreateNoteResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}
	return &pb.CreateNoteResponse{
		Status: http.StatusOK,
		Id:     uint32(note.Id),
	}, nil
}

func (s *Server) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	if err := s.H.DB.Where("userid=$1 and id=$2", req.Uid, req.Id).Delete(&models.Note{}).Error; err != nil {
		return &pb.DeleteNoteResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}
	return &pb.DeleteNoteResponse{
		Status: http.StatusOK,
	}, nil
}
