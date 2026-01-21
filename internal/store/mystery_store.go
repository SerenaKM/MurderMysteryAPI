package store

import (
	"database/sql"
)

type CrimeScene struct {
	CrimeSceneID int    `json:"crimeSceneId"`
	Location     string `json:"location"`
	Description  string `json:"description"`
	KeyObjects   string `json:"keyObjects"`
}

type Suspect struct {
	SuspectID            int    `json:"suspectId"`
	Name                 string `json:"name"`
	Occupation           string `json:"occupation"`
	Backstory            string `json:"backstory"`
	PossibleMotive       string `json:"possibleMotive"`
	RelationshipToVictim string `json:"relationshipToVictim"`
	Alibi                string `json:"alibi"`
}

type Case struct {
	CaseID       int   `json:"caseId"`
	CrimeSceneID int   `json:"crimeSceneId"`
	SuspectsList []int `json:"suspectId"`
}

type PostgresMysteryStore struct { // handle Postgres database communication
	db *sql.DB
}

func NewPostgresMysteryStore(db *sql.DB) *PostgresMysteryStore {
	return &PostgresMysteryStore{db: db}
}

// decouple our database with an interface
type MysteryStore interface {
	CreateCase(*Case) (*Case, error)
	GetCaseByID(id int64) (*Case, error)
	DeleteCase(id int64) error
	// TO-DO Create crime scene
	// TO-DO Get crime scene
	// TO-DO Update crime scene
	// TO-DO Delete crime scene
	// TO-DO Create suspect
	// TO-DO Get crime scene
	// TO-DO Update crime scene
	// TO-DO Delete crime scene
}

func (pg *PostgresMysteryStore) CreateCase(mysteryCase *Case) (*Case, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // if anything happens, transaction can be rolled back to prevent partial data insertion

	// get a random crime_scene
	randomCrimeSceneQuery := `
		SELECT crimeSceneId FROM crime_scenes
		ORDER BY RANDOM()
		LIMIT 1
	`
	err = tx.QueryRow(randomCrimeSceneQuery).Scan(&mysteryCase.CrimeSceneID)
	if err != nil {
		return nil, err
	}

	// insert caseId and crimeSceneId into cases
	crimeSceneInsertQuery := `
		INSERT INTO cases (caseId, crimeSceneId)
		VALUES($1, $2)
		RETURNING caseId
	`

	err = tx.QueryRow(crimeSceneInsertQuery, mysteryCase.CaseID, mysteryCase.CrimeSceneID).Scan(&mysteryCase.CaseID)
	if err != nil {
		return nil, err
	}

	// get a random list of 3 suspects
	randomSuspectsQuery := `
		SELECT suspectId FROM suspects
		ORDER BY RANDOM()
		LIMIT 3
	`

	rows, err := tx.Query(randomSuspectsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var randomSuspectId int
		err = rows.Scan(&randomSuspectId)
		if err != nil {
			return nil, err
		}
		mysteryCase.SuspectsList = append(mysteryCase.SuspectsList, randomSuspectId)
	}

	// insert caseId and suspectId into case_suspects
	suspectInsertQuery := `
		INSERT INTO case_suspects (caseId, suspectId)
		VALUES($1, $2)
	`

	for _, suspectId := range mysteryCase.SuspectsList {
		_, err := tx.Exec(suspectInsertQuery, mysteryCase.CaseID, suspectId)
		if err != nil {
			return nil, err
		}
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return mysteryCase, nil
}

func (pg *PostgresMysteryStore) GetCaseByID(id int64) (*Case, error) {
	mysteryCase := &Case{}

	crimeSceneQuery := `
		SELECT caseId, crimeSceneId
		FROM cases
		WHERE caseId = $1
	`

	err := pg.db.QueryRow(crimeSceneQuery, id).Scan(&mysteryCase.CaseID, &mysteryCase.CrimeSceneID)

	if err == sql.ErrNoRows{
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	suspectEntryQuery := `
		SELECT suspectId
		FROM case_suspects
		WHERE caseId = $1
	`

	// iterate through the multiple suspects and append
	rows, err := pg.db.Query(suspectEntryQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var suspect int
		err = rows.Scan(&suspect)
		if err != nil {
			return nil, err
		}
		mysteryCase.SuspectsList = append(mysteryCase.SuspectsList, suspect)
	}	

	return mysteryCase, nil
}

func (pg *PostgresMysteryStore) DeleteCase(id int64) error {
	query := `
		DELETE from cases
		WHERE caseId = $1
	`

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected() // verify if delete was successful
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}