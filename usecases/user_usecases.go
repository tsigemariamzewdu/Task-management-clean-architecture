package usecases

import (
	"context"
	"errors"
	domain "task_management/Domain"
	repositories "task_management/Repositories"
	infrastruture "task_management/infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCase struct {
	UserRepo repositories.UserRepository
	PasswordService infrastruture.PasswordServiceImpl
	JWTService infrastruture.JWTServiceImpl

}

//register use case
func (uc *UserUseCase) Register (ctx context.Context,input *domain.RegisterUserInput)(*domain.User,error){
	ctx,cancel:= context.WithTimeout(ctx ,10*time.Second)
	defer cancel()
	//first count to check if the user already exits or not

	count,err := uc.UserRepo.CountByUsername(ctx,input.Username)
	if err!=nil{
		return nil,errors.New("error while checking existing user")
	}
	if count>0{
		return nil,errors.New("username already exists")
	}
	//check number of total users and if 0 make the first user and admin
	totalUsers,err:= uc.UserRepo.CountAll(ctx)
	if err !=nil{
		return nil,errors.New("error checking total users")

	}
	//hash the password
	hashedPassword,err:=uc.PasswordService.HashPassword(input.Password)
	if err !=nil{
		return nil,errors.New("failed to hash password")

	}
	//set the role as user first then check if it the first user and if so make it admin
	role:=domain.RoleUser
	if totalUsers==0{
		role=domain.RoleAdmin
	}
	newUser := &domain.User{
		ID: primitive.NewObjectID(),
		Username: input.Username,
		Password: hashedPassword,
		Role: role,
	}
	err= uc.UserRepo.CreateUser(ctx,newUser)
	if err !=nil{
		return nil,errors.New("failed to add user")
	}
	return newUser,nil
}

//login use case 

func (uc *UserUseCase) Login(ctx context.Context, input domain.RegisterUserInput) (string, *domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	//find username 
	user, err := uc.UserRepo.FindByUsername(ctx, input.Username)
	if err != nil {
		return "", nil, errors.New("invalid username or password")
	}
	
	//compare password 
	ok := uc.PasswordService.ComparePassword(user.Password, input.Password)
	if !ok {
		return "", nil, errors.New("invalid username or password")
	}
	//generate token
	token, err := uc.JWTService.GenerateToken(user.ID.Hex(), string(user.Role))
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	return token, user, nil
}
//promoteuser usecase 
func (uc * UserUseCase) PromoteUser(ctx context.Context,userID string)error{
	return uc.UserRepo.PromoteUser(ctx,userID)
}

