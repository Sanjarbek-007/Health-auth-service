package postgres

// import (
// 	pb "auth-service/genproto/user"
// 	"auth-service/logger"
// 	"auth-service/storage"
// 	"context"
// 	"reflect"
// 	"testing"
// )

// func ReservationRepo(t *testing.T) storage.IStorage {
// 	db, err := ConnectionDb()
// 	log := logger.NewLogger()
// 	if err != nil {
// 		t.Error("ERROR : ", err)
// 		return nil
// 	}
// 	userRepo := NewPostgresStorage(db, log)
// 	return userRepo
// }

// func TestCreateUser(t *testing.T) {
// 	reservationRepo := ReservationRepo(t)

//     request := pb.RegisterReq{
//         Username:     "admin",
//         Email:        "admin@example.com",
//         Password:     "admin123",
//         Fullname:     "Admin User",
//         NativeLanguage: "English",
//         Role:          "admin",
//     }

//     res, err := reservationRepo.User().CreateUser(context.Background(), &request)
//     if err!= nil {
//         t.Error("ERROR : ", err)
//         return
//     }

//     case1 := pb.RegisterRes{Id: res.Id}

//     if!reflect.DeepEqual(res, &case1) {
//         t.Error("Result : ", res, "Expected : ", &case1)
//         return
//     }
// }