package main

import (
	"flag"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Universal-Selfcare/utils/data"
)

type config struct {
	port  int
	env   string
	debug bool
	db    struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
	limiter struct {
		enabled bool
		rps     float64
		burst   int
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	cors struct {
		trustedOrigins []string
	}
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")
	flag.BoolVar(&cfg.debug, "debug", false, "Enable debug mode for stack traces on panic")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "NONE", "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(
		&cfg.db.maxIdleTime,
		"db-max-idle-time",
		15*time.Minute,
		"PostgreSQL max connection idle time",
	)

	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "a7420fc0883489", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "e75ffd0a3aa5ec", "SMTP password")
	flag.StringVar(
		&cfg.smtp.sender,
		"smtp-sender",
		"Greenlight <no-reply@greenlight.alexedwards.net>",
		"SMTP sender",
	)

	flag.Func(
		"cors-trusted-origins",
		"Trusted CORS origins (space separated)",
		func(val string) error {
			cfg.cors.trustedOrigins = strings.Fields(val)
			return nil
		},
	)

	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		return
	}

	db.AutoMigrate(
		data.User{},
		data.Token{},
		data.MedicalInformation{},
		data.Caregiver{},
		data.Allergy{},
		data.MedicalEvent{},
		data.FrequentFood{},
		data.Medication{},
		data.EmergencyContact{},
		data.DietarySupplement{},
	)
}

func openDB(cfg config) (*gorm.DB, error) {
	dsn := cfg.db.dsn // dsn == connection string

	postgresDB := postgres.Open(dsn)
	db, err := gorm.Open(postgresDB, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.db.maxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.db.maxIdleConns)
	sqlDB.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	return db, nil
}
