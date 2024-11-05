package config

import (
	"os"

	"github.com/vinicius457/BancoDeDados2/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
    host = "postgres_db"
    port = "5432"
    user = "postgres"
    password = "postgres"
    dbname = "restaurante"
)

func ConnectPostgresql() (*gorm.DB, error) {
	// logger := GetLogger(("postgresql"))
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	// if host == "" || port == "" || user == "" || password == "" || dbname == "" {
	// 	errDotEnv :=  errors.New("Missing environment variables for database connection")
    //     logger.Warnf("postgresql connection warning: %v", errDotEnv)
    //     logger.Info("Initializing default database connection...")
        // host := "postgres_db"
        // port := "5432"
        // user := "postgres"
        // password := "postgres"
        // dbname := "restaurante"
    //}
	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!= nil {
        logger.Errorf("postgresql connection error: %v", err)
        return nil, err
    }

	err = db.AutoMigrate(&model.Cliente{}, &model.Fornecedor{}, &model.Ingrediente{}, &model.Prato{}, &model.Venda{}, &model.Usos{})
	if err!= nil {
        logger.Errorf("postgresql migration error: %v", err)
        return nil, err
    }

    err = initDatabase(db)
    if err!= nil {
        logger.Warnf("Warning... something wrong with initializing database: %v", err)
        //return nil, err
    }

	return db, nil
}

func initDatabase(db *gorm.DB) error {
    sqlFile := "config/Restaurante.sql"

    sqlBytes, err := os.ReadFile(sqlFile)
    if err != nil {
        logger.Errorf("Error reading SQL file for init database: %v", err)
        return err
    }

    // Executa o comando e verifica se ocorreu erro dizendo qual comando falhou
    commands := string(sqlBytes)
    if result := db.Exec(commands); result.Error != nil {
        logger.Errorf("Error executing SQL command: %v",result.Error)
        return result.Error
    }
    return nil
}