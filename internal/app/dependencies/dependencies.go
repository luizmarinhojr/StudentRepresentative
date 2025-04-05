package dependencies

import (
	"database/sql"

	"github.com/luizmarinhojr/StudentRepresentative/internal/app/repository"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/service"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase"
	"github.com/luizmarinhojr/StudentRepresentative/internal/app/usecase/validation"
	"github.com/luizmarinhojr/StudentRepresentative/internal/http/gin/handler"
)

type Dependencies struct {
	StudentHandler *handler.StudentHandler
	UserHandler    *handler.UserHandler
	ClassHandler   *handler.ClassHandler
}

func Inject(db *sql.DB) *Dependencies {
	// REPOSITORIES
	studentRepository := repository.NewStudentRepository(db)
	userRepository := repository.NewUserRepository(db)
	classRepository := repository.NewClassRepository(db)

	// SERVICES
	passwordService := service.NewPasswordService()

	// VALIDATORS
	studentRegisterValidator := []validation.StudentRegisterValidator{
		validation.NewStudentDuplicationByRegister(*studentRepository),
	}

	userRegisterValidator := []validation.UserRegisterValidator{
		validation.NewUserIsStudentExists(*studentRepository),
		validation.NewStudentHaveUser(*studentRepository),
		validation.NewUserDuplicationByEmail(*userRepository),
		validation.NewUserValidationByLastName(*studentRepository),
	}

	// USECASES
	studentUseCase := usecase.NewStudentUseCase(*studentRepository, studentRegisterValidator...)
	userUseCase := usecase.NewUserUseCase(*userRepository, *studentRepository, *passwordService, studentRegisterValidator, userRegisterValidator...)
	classUseCase := usecase.NewClassUseCase(*classRepository)

	// HANDLERS
	studentHanler := handler.NewStudentController(*studentUseCase)
	userHandler := handler.NewUserHandler(*userUseCase)
	classHandler := handler.NewClassHandler(*classUseCase)

	return &Dependencies{
		StudentHandler: studentHanler,
		UserHandler:    userHandler,
		ClassHandler:   classHandler,
	}
}
