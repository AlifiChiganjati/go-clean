package manager

import (
	"fmt"

	"github.com/AlifiChiganjati/go-clean/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	Connect() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	connect := &infraManager{cfg: cfg}
	if err := connect.openConnect(); err != nil {
		return nil, err
	}
	return connect, nil
}

func (i *infraManager) openConnect() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		i.cfg.Host,
		i.cfg.Port,
		i.cfg.User,
		i.cfg.Password,
		i.cfg.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	i.db = db
	return nil
}

func (i *infraManager) Connect() *gorm.DB {
	return i.db
}
