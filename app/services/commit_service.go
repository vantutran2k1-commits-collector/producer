package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/vantutran2k1-commits-collector/producer/app/constants"
	"github.com/vantutran2k1-commits-collector/producer/app/payloads"
	"github.com/vantutran2k1-commits-collector/producer/app/repositories"
	"github.com/vantutran2k1-commits-collector/producer/config"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CommitService interface {
	Collect() ([]*payloads.CommitPayload, error)
}

func NewCommitService(
	jobRepo repositories.JobRepository,
	kafkaProducer sarama.SyncProducer,
) CommitService {
	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}
	return &commitService{
		jobRepo:       jobRepo,
		kafkaProducer: kafkaProducer,
		httpClient:    httpClient,
	}
}

type commitService struct {
	jobRepo       repositories.JobRepository
	kafkaProducer sarama.SyncProducer
	httpClient    *http.Client
}

func (s *commitService) Collect() ([]*payloads.CommitPayload, error) {
	latestJob, err := s.jobRepo.GetLatestJob()
	if err != nil {
		return nil, err
	}

	var fromTime *time.Time
	if latestJob != nil {
		fromTime = &latestJob.CollectedFrom
	}
	commits, err := s.extractCommits(fromTime)
	if err != nil {
		return nil, err
	}

	return commits, nil
}

func (s *commitService) extractCommits(fromTime *time.Time) ([]*payloads.CommitPayload, error) {
	url := fmt.Sprintf("%s?per_page=20&page={pageNumber}", config.AppEnv.GithubCommitsApi)
	if fromTime != nil {
		url = fmt.Sprintf("%s&since=%s", url, fromTime.Format("2006-01-02T15:04:05Z"))
	}

	var commits []*payloads.CommitPayload
	var current []*payloads.CommitPayload
	page := 1
	for {
		currentUrl := strings.Replace(url, "{pageNumber}", strconv.Itoa(page), 1)
		fmt.Printf("URL: %s\n", currentUrl)
		req, err := http.NewRequest(http.MethodGet, currentUrl, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set(constants.GithubApiVersionHeader, "2022-11-28")
		if config.AppEnv.GithubApiKey != "" {
			req.Header.Set(constants.AuthorizationHeader, fmt.Sprintf("Bearer %s", config.AppEnv.GithubApiKey))
		}

		resp, err := s.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("unexpected response from location API")
		}

		if err := json.NewDecoder(resp.Body).Decode(&current); err != nil {
			return nil, err
		}

		if len(current) == 0 {
			break
		}
		commits = append(commits, current...)

		page++
	}

	return commits, nil
}

func (s *commitService) sendToTopic(payload payloads.CommitPayload) error {
	messageBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, _, err = s.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: config.AppEnv.KafkaGithubCommitsTopic,
		Value: sarama.ByteEncoder(messageBytes),
	})
	if err != nil {
		return err
	}

	return nil
}
