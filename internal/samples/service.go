package samples

import (
	"database/sql"
	"log"
)

type SampleService struct {
	db *sql.DB
}

func NewSampleService(db *sql.DB) *SampleService {
    return &SampleService{db: db}
}

func (s *SampleService) GetSamplesAllSamples(userID int) ([]Sample, error) {
	rows, err := s.db.Query("SELECT id, sample_name, sample_body FROM samples WHERE user_id = $1", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    samples := make([]Sample, 0)
    for rows.Next() {
        var sample Sample
        err := rows.Scan(&sample.ID, &sample.SampleName, &sample.SampleBody)
		log.Println(err)
        if err != nil {
            return nil, err
        }
        samples = append(samples, sample)
    }

    return samples, nil
}

func (s *SampleService) GetSampleByID(userID, sampleID int) (*Sample, error) {
	var sample Sample
	err := s.db.QueryRow("SELECT id, sample_name, sample_body FROM samples WHERE user_id = $1 AND id = $2", userID, sampleID).Scan(&sample.ID, &sample.SampleName, &sample.SampleBody)
	if err != nil {
		return nil, err
	}
	return &sample, nil
}

func (s *SampleService) CreateSample(userID int, sample *Sample) error {
    _, err := s.db.Exec("INSERT INTO samples (sample_name, sample_body, user_id) VALUES ($1, $2, $3)", sample.SampleName, sample.SampleBody, userID)
    return err
}

func (s *SampleService) UpdateSample(userID, sampleID int, sample *Sample) error {
    _, err := s.db.Exec("UPDATE samples SET sample_name = $1, sample_body = $2 WHERE user_id = $3 AND id = $4", sample.SampleName, sample.SampleBody, userID, sampleID)
    return err
}

func (s *SampleService) DeleteSample(userID, sampleID int) error {
    _, err := s.db.Exec("DELETE FROM samples WHERE user_id = $1 AND id = $2", userID, sampleID)
    return err
}