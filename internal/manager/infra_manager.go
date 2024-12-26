package manager

import (
	"fmt"

	"github.com/AlifiChiganjati/go-clean/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	InfraManager interface {
		Conn() *gorm.DB
	}

	infraManager struct {
		db  *gorm.DB
		cfg *config.Config
	}
)

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}
	return conn, nil
}

func (i *infraManager) openConn() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	i.db = db
	return nil
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}
