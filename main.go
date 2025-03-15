package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Universal-Selfcare/utils/data"
	"github.com/Universal-Selfcare/utils/password"
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

// Test users to be created
var users = []struct {
	Username       string
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	Password       string
	IntakeComplete bool
}{
	{
		Username:       "john_doe",
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@example.com",
		PhoneNumber:    "10123456789",
		Password:       "password123",
		IntakeComplete: true,
	},
	{
		Username:       "jane_smith",
		FirstName:      "Jane",
		LastName:       "Smith",
		Email:          "jane.smith@example.com",
		PhoneNumber:    "10987654321",
		Password:       "securepass456",
		IntakeComplete: false,
	},
	{
		Username:       "robert_johnson",
		FirstName:      "Robert",
		LastName:       "Johnson",
		Email:          "robert.j@example.com",
		PhoneNumber:    "10567891234",
		Password:       "robert2025",
		IntakeComplete: true,
	},
	{
		Username:       "sarah_williams",
		FirstName:      "Sarah",
		LastName:       "Williams",
		Email:          "sarah.w@example.com",
		PhoneNumber:    "10345678912",
		Password:       "sarah_pass789",
		IntakeComplete: false,
	},
	{
		Username:       "michael_brown",
		FirstName:      "Michael",
		LastName:       "Brown",
		Email:          "michael.b@example.com",
		PhoneNumber:    "10678912345",
		Password:       "brownmike2025",
		IntakeComplete: true,
	},
}

// Sample medical information data
var sampleConditions = []string{
	"None", "Asthma", "Diabetes", "Hypertension", "Arthritis", "Allergies",
}

var sampleTriggers = []string{
	"Stress", "Lack of sleep", "Poor diet", "Dehydration", "Environmental factors",
}

var sampleChanges = []string{
	"Improve diet", "Exercise more", "Reduce stress", "Better sleep habits", "Stay hydrated",
}

var sampleAllergies = []string{
	"Peanuts", "Shellfish", "Pollen", "Dust", "Penicillin", "Latex", "Dairy",
}

var sampleReactions = []string{
	"Rash", "Swelling", "Difficulty breathing", "Itching", "Nausea",
}

var sampleMedications = []string{
	"Lisinopril", "Metformin", "Atorvastatin", "Levothyroxine", "Albuterol", "Omeprazole",
}

var sampleFoods = []string{
	"Chicken", "Rice", "Broccoli", "Salmon", "Yogurt", "Eggs", "Oatmeal", "Sweet potatoes",
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

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create users
	fmt.Println("Creating users with the following credentials:")
	fmt.Println("Username\t|\tPassword\t|\tIntake Complete")
	fmt.Println("------------------------------------------------------")

	for _, u := range users {
		// Hash the password
		hashedPassword, err := password.HashPassword(u.Password)
		if err != nil {
			log.Printf("Failed to hash password for %s: %v", u.Username, err)
			continue
		}

		// Create the user
		user := &data.User{
			UserName:           u.Username,
			FirstName:          u.FirstName,
			LastName:           u.LastName,
			Email:              u.Email,
			PhoneNumber:        u.PhoneNumber,
			Hash:               hashedPassword,
			UserIntakeComplete: u.IntakeComplete,
		}

		err = db.Create(user).Error
		if err != nil {
			log.Printf("Failed to create user %s: %v", u.Username, err)
			continue
		}

		fmt.Printf("%s\t|\t%s\t|\t%t\n", u.Username, u.Password, u.IntakeComplete)

		// If user has completed intake, add medical information
		if u.IntakeComplete {
			createMedicalInformation(db, user.ID)
			createRandomAllergies(db, user.ID)
			createRandomMedications(db, user.ID)
			createRandomFrequentFoods(db, user.ID)
			createEmergencyContact(db, user.ID)
		}
	}

	fmt.Println("\nDatabase seeding completed successfully!")
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

func createMedicalInformation(db *gorm.DB, userID int64) {
	medInfo := &data.MedicalInformation{
		UserID:            userID,
		Height:            fmt.Sprintf("%d", 150+rand.Intn(50)),
		Weight:            fmt.Sprintf("%d", 50+rand.Intn(70)),
		Diagnosis:         randomElement(sampleConditions),
		DiagnosisSeverity: randomElement([]string{"Mild", "Moderate", "Severe"}),
		CurrentPriority: randomElement(
			[]string{"Improve health", "Manage symptoms", "Preventive care"},
		),
		Gender:            randomElement([]string{"Male", "Female", "Other"}),
		State:             randomElement([]string{"CA", "NY", "TX", "FL", "IL"}),
		ContactPreference: randomElement([]string{"Phone", "Email"}),
		CaregivingPriority: randomElement(
			[]string{"Medication management", "Diet improvement", "Stress reduction"},
		),
		HealthTriggers:  randomElement(sampleTriggers),
		DesiredChanges:  randomElement(sampleChanges),
		OralAntibiotics: randomBool(),
		BrainFog:        randomBool(),
		ChronicPain:     randomBool(),
		Eczema:          randomBool(),
		Diarrhea:        randomBool(),
		HeartDisease:    randomBool(),
		OtherConditions: randomElement(sampleConditions),
	}

	err := db.Create(medInfo).Error
	if err != nil {
		log.Printf("Failed to create medical information for user %d: %v", userID, err)
	}
}

func createRandomAllergies(db *gorm.DB, userID int64) {
	// Create 1-3 random allergies
	numAllergies := 1 + rand.Intn(3)

	for i := 0; i < numAllergies; i++ {
		allergy := &data.Allergy{
			UserID:      userID,
			AllergyName: randomElement(sampleAllergies),
			Reaction:    randomElement(sampleReactions),
		}

		err := db.Create(allergy).Error
		if err != nil {
			log.Printf("Failed to create allergy for user %d: %v", userID, err)
		}
	}
}

func createRandomMedications(db *gorm.DB, userID int64) {
	// Create 1-2 random medications
	numMeds := 1 + rand.Intn(2)

	for i := 0; i < numMeds; i++ {
		current := randomBool()
		endDate := ""
		if !current {
			endDate = randomDate(2024, 2025)
		}

		medication := &data.Medication{
			UserID:    userID,
			Name:      randomElement(sampleMedications),
			Dosage:    randomElement([]string{"10mg", "25mg", "50mg", "100mg"}),
			StartDate: randomDate(2020, 2024),
			EndDate:   endDate,
			Current:   current,
			SideEffects: randomElement(
				[]string{"None", "Drowsiness", "Nausea", "Headache", "Dizziness"},
			),
		}

		err := db.Create(medication).Error
		if err != nil {
			log.Printf("Failed to create medication for user %d: %v", userID, err)
		}
	}
}

func createRandomFrequentFoods(db *gorm.DB, userID int64) {
	// Create 2-4 random frequent foods
	numFoods := 2 + rand.Intn(3)

	// Shuffle the foods array to get random selection
	foods := make([]string, len(sampleFoods))
	copy(foods, sampleFoods)
	rand.Shuffle(len(foods), func(i, j int) { foods[i], foods[j] = foods[j], foods[i] })

	for i := 0; i < numFoods && i < len(foods); i++ {
		food := &data.FrequentFood{
			UserID:   userID,
			FoodName: foods[i],
		}

		err := db.Create(food).Error
		if err != nil {
			log.Printf("Failed to create frequent food for user %d: %v", userID, err)
		}
	}
}

func createEmergencyContact(db *gorm.DB, userID int64) {
	contact := &data.EmergencyContact{
		UserID:      userID,
		FirstName:   randomElement([]string{"Alex", "Sam", "Jamie", "Chris", "Jordan"}),
		LastName:    randomElement([]string{"Johnson", "Williams", "Davis", "Miller", "Wilson"}),
		PhoneNumber: fmt.Sprintf("1%09d", rand.Intn(1000000000)),
		Email:       fmt.Sprintf("emergency%d@example.com", rand.Intn(1000)),
	}

	err := db.Create(contact).Error
	if err != nil {
		log.Printf("Failed to create emergency contact for user %d: %v", userID, err)
	}
}

// Helper functions

func randomElement(array []string) string {
	return array[rand.Intn(len(array))]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomDate(minYear, maxYear int) string {
	year := minYear + rand.Intn(maxYear-minYear+1)
	month := 1 + rand.Intn(12)
	day := 1 + rand.Intn(28)
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}
