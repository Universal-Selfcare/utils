package data

import (
	"time"

	"github.com/Universal-Selfcare/utils/validator"
)

// TODO it might make sense to refactor this into blocks which are more often used/updated independently
type MedicalInformation struct {
	ID     int64 `gorm:"primaryKey"     json:"id"`
	UserID int64 `gorm:"not null;index" json:"user_id"`

	// Basic information
	Height            uint   `gorm:"not null"           json:"height"`
	Weight            uint   `gorm:"not null"           json:"weight"`
	Diagnosis         string `gorm:"type:text;not null" json:"diagnosis"`
	DiagnosisSeverity string `gorm:"type:text;not null" json:"diagnosis_severity"`
	CurrentPriority   string `gorm:"type:text;not null" json:"current_priority"` // "What is most important to you today?"
	Gender            string `gorm:"type:text;not null" json:"gender"`

	OralAntibiotics                 bool `gorm:"default:false" json:"oral_antibiotics"`
	FrequentHydroLotions            bool `gorm:"default:false" json:"frequent_hydro_lotions"`
	MetalsOrMagnesiumPowder         bool `gorm:"default:false" json:"metals_or_magnesium_powder"`
	UnfilteredTapWater              bool `gorm:"default:false" json:"unfiltered_tap_water"`
	PesticidesFromFarm              bool `gorm:"default:false" json:"pesticides_from_farm"`
	TwoOrMoreHoursScreenTime        bool `gorm:"default:false" json:"two_or_more_hours_screen_time"`
	DentalOrBodyXRays               bool `gorm:"default:false" json:"dental_or_body_x_rays"`
	FrequentWirelessDevice          bool `gorm:"default:false" json:"frequent_wireless_device"`
	WaterLeakageInBasement          bool `gorm:"default:false" json:"water_leakage_in_basement"`
	MustyMildewSmell                bool `gorm:"default:false" json:"musty_mildew_smell"`
	FrequentDeodorantWithNailPolish bool `gorm:"default:false" json:"frequent_deodorant_with_nail_polish"`
	CannedFoodsThermalReceipts      bool `gorm:"default:false" json:"canned_foods_thermal_receipts"`
	ContactWithBuildingMaterials    bool `gorm:"default:false" json:"contact_with_building_materials"`
	DailyUsePlasticUtensils         bool `gorm:"default:false" json:"daily_use_plastic_utensils"`
	FrequentMealsShellfishLargeFish bool `gorm:"default:false" json:"frequent_meals_shellfish_large_fish"`

	TraumaOrNightmares            bool `gorm:"default:false" json:"trauma_or_nightmares"`
	ScreamsOrShrieks              bool `gorm:"default:false" json:"screams_or_shrieks"`
	MoodSwings                    bool `gorm:"default:false" json:"mood_swings"`
	Irritability                  bool `gorm:"default:false" json:"irritability"`
	BrainFog                      bool `gorm:"default:false" json:"brain_fog"`
	DifficultyConcentrating       bool `gorm:"default:false" json:"difficulty_concentrating"`
	AnxietyDarkThoughts           bool `gorm:"default:false" json:"anxiety_dark_thoughts"`
	AttentionDeficitHyperactivity bool `gorm:"default:false" json:"attention_deficit_hyperactivity"`
	BipolarDisorder               bool `gorm:"default:false" json:"bipolar_disorder"`
	Schizophrenia                 bool `gorm:"default:false" json:"schizophrenia"`
	SensoryIntegrationDisorder    bool `gorm:"default:false" json:"sensory_integration_disorder"`
	Autism                        bool `gorm:"default:false" json:"autism"`

	// Body Symptoms
	HairIsThinning              bool `gorm:"default:false" json:"hair_is_thinning"`
	BleedingGums                bool `gorm:"default:false" json:"bleeding_gums"`
	Gingivitis                  bool `gorm:"default:false" json:"gingivitis"`
	CoatedTongue                bool `gorm:"default:false" json:"coated_tongue"`
	Stammering                  bool `gorm:"default:false" json:"stammering"`
	DizzinessSpinning           bool `gorm:"default:false" json:"dizziness_spinning"`
	LimitedSpeech               bool `gorm:"default:false" json:"limited_speech"`
	AnswersbyRepeatingSchedulal bool `gorm:"default:false" json:"answers_by_repeating_schedual"`
	PoorEyeContact              bool `gorm:"default:false" json:"poor_eye_contact"`
	DifficultyFallingAsleep     bool `gorm:"default:false" json:"difficulty_falling_asleep"`
	WakeUpMiddleOfNight         bool `gorm:"default:false" json:"wake_up_middle_of_night"`
	ChronicCough                bool `gorm:"default:false" json:"chronic_cough"`
	ChronicRunnyNose            bool `gorm:"default:false" json:"chronic_runny_nose"`
	AbnormalEarlyDevelopment    bool `gorm:"default:false" json:"abnormal_early_development"`
	PainfulPeriods              bool `gorm:"default:false" json:"painful_periods"`
	HeadachesOrMigraines        bool `gorm:"default:false" json:"headaches_or_migraines"`
	HeartPalpitations           bool `gorm:"default:false" json:"heart_palpitations"`
	FrequentlyCatchesInfections bool `gorm:"default:false" json:"frequently_catches_infections"`
	SinusCongestion             bool `gorm:"default:false" json:"sinus_congestion"`
	ChronicEarAche              bool `gorm:"default:false" json:"chronic_ear_ache"`
	TinglingInHandsOrFeet       bool `gorm:"default:false" json:"tingling_in_hands_or_feet"`
	SexualDysfunction           bool `gorm:"default:false" json:"sexual_dysfunction"`
	MuscleCrampsOrTwitch        bool `gorm:"default:false" json:"muscle_cramps_or_twitch"`
	AthletesFoot                bool `gorm:"default:false" json:"athletes_foot"`
	JockItch                    bool `gorm:"default:false" json:"jock_itch"`
	FungalNailInfections        bool `gorm:"default:false" json:"fungal_nail_infections"`
	ChronicAcheOrPain           bool `gorm:"default:false" json:"chronic_ache_or_pain"`

	Eczema           bool `gorm:"default:false" json:"eczema"`
	Acne             bool `gorm:"default:false" json:"acne"`
	Psoriasis        bool `gorm:"default:false" json:"psoriasis"`
	DrySkin          bool `gorm:"default:false" json:"dry_skin"`
	Rash             bool `gorm:"default:false" json:"rash"`
	Burning          bool `gorm:"default:false" json:"burning"`
	Hives            bool `gorm:"default:false" json:"hives"`
	ItchyEar         bool `gorm:"default:false" json:"itchy_ear"`
	ItchyScalpNation bool `gorm:"default:false" json:"itchy_scalp_nation"`
	ItchyGenitalArea bool `gorm:"default:false" json:"itchy_genital_area"`
	TinyBumpsOnCheek bool `gorm:"default:false" json:"tiny_bumps_on_cheek"`

	BadBreath                   bool `gorm:"default:false" json:"bad_breath"`
	CavitiesDentalHealth        bool `gorm:"default:false" json:"cavities_dental_health"`
	BleedingGumsGI              bool `gorm:"default:false" json:"bleeding_gums_gi"`
	CoatedTongueGI              bool `gorm:"default:false" json:"coated_tongue_gi"`
	BloatingInStomach           bool `gorm:"default:false" json:"bloating_in_stomach"`
	MoreThan2BowlsDaily         bool `gorm:"default:false" json:"more_than_2_bowls_daily"`
	Diarrhea                    bool `gorm:"default:false" json:"diarrhea"`
	Constipation                bool `gorm:"default:false" json:"constipation"`
	FrequentUrinationBedWetting bool `gorm:"default:false" json:"frequent_urination_bed_wetting"`
	StoolWithUndigestedFood     bool `gorm:"default:false" json:"stool_with_undigested_food"`
	BladderInfection            bool `gorm:"default:false" json:"bladder_infection"`

	IrritableBowelSyndrome             bool   `gorm:"default:false" json:"irritable_bowel_syndrome"`
	UlcerativeColitis                  bool   `gorm:"default:false" json:"ulcerative_colitis"`
	GastritisOrPepticUlcer             bool   `gorm:"default:false" json:"gastritis_or_peptic_ulcer"`
	GERD                               bool   `gorm:"default:false" json:"gerd"`
	CeliacDisease                      bool   `gorm:"default:false" json:"celiac_disease"`
	HeartDisease                       bool   `gorm:"default:false" json:"heart_disease"`
	ElevatedOrLowCholesterol           bool   `gorm:"default:false" json:"elevated_or_low_cholesterol"`
	HighBloodPressure                  bool   `gorm:"default:false" json:"high_blood_pressure"`
	POTSDysautonomia                   bool   `gorm:"default:false" json:"pots_dysautonomia"`
	RheumaticFever                     bool   `gorm:"default:false" json:"rheumatic_fever"`
	MitralValveProlapse                bool   `gorm:"default:false" json:"mitral_valve_prolapse"`
	Type1Diabetes                      bool   `gorm:"default:false" json:"type_1_diabetes"`
	Type2Diabetes                      bool   `gorm:"default:false" json:"type_2_diabetes"`
	Hypoglycemia                       bool   `gorm:"default:false" json:"hypoglycemia"`
	InsulinResistanceOrPrediabetes     bool   `gorm:"default:false" json:"insulin_resistance_or_prediabetes"`
	Hypothyroidism                     bool   `gorm:"default:false" json:"hypothyroidism"`
	Hyperthyroidism                    bool   `gorm:"default:false" json:"hyperthyroidism"`
	EndocrineProblems                  bool   `gorm:"default:false" json:"endocrine_problems"`
	WeightGain                         bool   `gorm:"default:false" json:"weight_gain"`
	WeightLoss                         bool   `gorm:"default:false" json:"weight_loss"`
	WeightFluctuations                 bool   `gorm:"default:false" json:"weight_fluctuations"`
	OtherEatingDisorder                bool   `gorm:"default:false" json:"other_eating_disorder"`
	MitochondrialDysfunction           bool   `gorm:"default:false" json:"mitochondrial_dysfunction"`
	FolateDeficiency                   bool   `gorm:"default:false" json:"folate_deficiency"`
	FattyAcidOxidationDefect           bool   `gorm:"default:false" json:"fatty_acid_oxidation_defect"`
	KidneyStones                       bool   `gorm:"default:false" json:"kidney_stones"`
	UrinaryTractInfections             bool   `gorm:"default:false" json:"urinary_tract_infections"`
	YeastInfections                    bool   `gorm:"default:false" json:"yeast_infections"`
	Arthritis                          bool   `gorm:"default:false" json:"arthritis"`
	Fibromyalgia                       bool   `gorm:"default:false" json:"fibromyalgia"`
	ChronicPain                        bool   `gorm:"default:false" json:"chronic_pain"`
	ChronicFatigueSyndrome             bool   `gorm:"default:false" json:"chronic_fatigue_syndrome"`
	AutoimmuneDisease                  string `gorm:"type:text"     json:"autoimmune_disease"` // Free text field
	RheumatoidArthritis                bool   `gorm:"default:false" json:"rheumatoid_arthritis"`
	Lupus                              bool   `gorm:"default:false" json:"lupus"`
	ImmuneDeficiencyDisease            bool   `gorm:"default:false" json:"immune_deficiency_disease"`
	PoorImmuneFunction                 bool   `gorm:"default:false" json:"poor_immune_function"`
	FoodAllergies                      bool   `gorm:"default:false" json:"food_allergies"`
	EnvironmentalAllergies             bool   `gorm:"default:false" json:"environmental_allergies"`
	MultipleChemicalSensitivities      bool   `gorm:"default:false" json:"multiple_chemical_sensitivities"`
	LatexAllergy                       bool   `gorm:"default:false" json:"latex_allergy"`
	FrequentEarInfections              bool   `gorm:"default:false" json:"frequent_ear_infections"`
	FrequentSinusInfections            bool   `gorm:"default:false" json:"frequent_sinus_infections"`
	FrequentUpperRespiratoryInfections bool   `gorm:"default:false" json:"frequent_upper_respiratory_infections"`
	Bronchitis                         bool   `gorm:"default:false" json:"bronchitis"`
	SleepApnea                         bool   `gorm:"default:false" json:"sleep_apnea"`
	TiredALotOfTheTime                 bool   `gorm:"default:false" json:"tired_a_lot_of_the_time"`
	CantFallAsleep                     bool   `gorm:"default:false" json:"cant_fall_asleep"`
	NeurologicalSymptoms               bool   `gorm:"default:false" json:"neurological_symptoms"`
	SensitivityToStimuli               bool   `gorm:"default:false" json:"sensitivity_to_stimuli"`
	BullsEyeRash                       bool   `gorm:"default:false" json:"bulls_eye_rash"`
	SweatingHeadacheCognitive          bool   `gorm:"default:false" json:"sweating_headache_cognitive"`

	OtherConditions       string `gorm:"type:text"     json:"other_conditions"`        // Free text field
	FoodRelatedConditions string `gorm:"type:text"     json:"food_related_conditions"` // Free text field
	AppendixRemoved       bool   `gorm:"default:false" json:"appendix_removed"`        // Yes/No field

	HealthTriggers string `gorm:"type:text" json:"health_triggers"`
	DesiredChanges string `gorm:"type:text" json:"desired_changes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type MedicalInformationStore interface {
	CreateMedicalInformation(userIntake *MedicalInformation) (*MedicalInformation, error)
	GetMedicalInformation(id int64) (*MedicalInformation, error)
	GetMedicalInformationByUserID(userID int64) (*MedicalInformation, error)
	UpdateMedicalInformation(userIntake *MedicalInformation) error
	DeleteMedicalInformation(id int64) error
}

func ValidateMedicalInformation(
	v *validator.Validator,
	info *MedicalInformation,
	allowPartial bool,
) {
  v.Check(info.Height > 0, "height", "must be provided")
}
