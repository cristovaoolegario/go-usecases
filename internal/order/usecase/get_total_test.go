package usecase

import (
	"database/sql"

	"github.com/cristovaoolegario/go-usecases/internal/order/entity"
	"github.com/cristovaoolegario/go-usecases/internal/order/infra/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type GetTotalUseCaseTestSuite struct {
	suite.Suite
	OrderRepository entity.OrderRepositoryInterface
	Db              *sql.DB
}

func (suite *GetTotalUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = database.NewOrderRepository(db)
}

func (suite *GetTotalUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func (suite *GetTotalUseCaseTestSuite) TestGetTotal() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	err = suite.OrderRepository.Save(order)
	suite.NoError(err)

	getTotalUseCase := NewGetTotalUseCase(suite.OrderRepository)

	output, err := getTotalUseCase.Execute()
	suite.NoError(err)

	suite.Equal(1, output.Total)

}
