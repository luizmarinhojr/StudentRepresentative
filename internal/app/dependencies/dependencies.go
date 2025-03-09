package dependencies

import (
	"database/sql"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/service"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase/validator"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/handler"
)

type Dependencies struct {
	StudentRepository        *repository.StudentRepository
	UserRepository           *repository.UserRepository
	PasswordService          *service.PasswordService
	StudentRegisterValidator []validator.StudentRegisterValidator
	UserRegisterValidator    []validator.UserRegisterValidator
	StudentUseCase           *usecase.StudentUseCase
	UserUseCase              *usecase.UserUseCase
	StudentHandler           *handler.StudentHandler
	UserHandler              *handler.UserHandler
}

func Inject(db *sql.DB) *Dependencies {
	// REPOSITORIES
	studentRepository := repository.NewStudentRepository(db)
	userRepository := repository.NewUserRepository(db)

	// SERVICES
	passwordService := service.NewPasswordService()

	// VALIDATORS
	studentRegisterValidator := []validator.StudentRegisterValidator{validator.NewStudentDuplicationByRegister(*studentRepository)}
	userRegisterValidator := []validator.UserRegisterValidator{validator.NewUserIsStudentExists(*studentRepository),
		validator.NewStudentHaveUser(*studentRepository), validator.NewUserDuplicationByEmail(*userRepository)}

	// USECASES
	studentUseCase := usecase.NewStudentUseCase(*studentRepository, studentRegisterValidator...)
	userUseCase := usecase.NewUserUseCase(*userRepository, *studentRepository, *passwordService, studentRegisterValidator, userRegisterValidator...)

	// HANDLERS
	studentHanler := handler.NewStudentController(*studentUseCase)
	userHandler := handler.NewUserHandler(*userUseCase)

	return &Dependencies{
		StudentRepository:        studentRepository,
		UserRepository:           userRepository,
		PasswordService:          passwordService,
		StudentRegisterValidator: studentRegisterValidator,
		UserRegisterValidator:    userRegisterValidator,
		StudentUseCase:           studentUseCase,
		UserUseCase:              userUseCase,
		StudentHandler:           studentHanler,
		UserHandler:              userHandler,
	}
}
