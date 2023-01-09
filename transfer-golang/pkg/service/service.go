package service

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gitlab.com/goxp/cloud0/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transfer/pkg/handlers"
	"transfer/pkg/repo"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*service.BaseApp
	setting            *extraSetting
	kafkaConsumer      *repo.KafkaConsumer
	DrugHandler        *handlers.DrugHandlers
	UserHandler        *handlers.UserHandlers
	MedicationHandlers *handlers.MedicationHandlers
	doneC              chan struct{}
	closeC             chan struct{}
}

func NewService() *Service {
	s := &Service{
		BaseApp: service.NewApp("transfer", "v1.0"),
		setting: &extraSetting{},
		doneC:   make(chan struct{}),
		closeC:  make(chan struct{}),
	}
	_ = env.Parse(s.setting)

	//---------Kafka Producer START---------
	kafkaProducer, err := repo.NewKafkaProducer()
	if err != nil {
		return nil
	}

	//---------Kafka Producer END---------
	Drug := handlers.NewDrugHandlers(kafkaProducer)
	User := handlers.NewUserHandlers(kafkaProducer)
	Medication := handlers.NewMedicationHandlers(kafkaProducer)
	s.DrugHandler = Drug
	s.UserHandler = User
	s.MedicationHandlers = Medication

	//---------Kafka Consumer START---------

	s.kafkaConsumer, err = repo.NewKafkaConsumer("golang-handler")
	if err != nil {
		return nil
	}

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		log.Println("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer func() {
			s.kafkaConsumer.Consumer.Unsubscribe()
			stop()
			cancel()
			close(errC)
		}()

		if err := s.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		log.Println("Shutdown completed")
	}()

	go func() {
		log.Println("Listening and serving")

		if err := s.ListenAndServe(); err != nil {
			errC <- err
		}
	}()
	//---------Kafka Consumer END---------

	return s
}

// ListenAndServe ...
func (s *Service) ListenAndServe() error {
	commit := func(msg *kafka.Message) {
		if _, err := s.kafkaConsumer.Consumer.CommitMessage(msg); err != nil {
			log.Println("commit failed", err.Error())
		}
	}

	go func() {
		run := true

		for run {
			select {
			case <-s.closeC:
				run = false
				break
			default:

				msg, ok := s.kafkaConsumer.Consumer.Poll(150).(*kafka.Message)
				if !ok {
					continue
				}

				if *msg.TopicPartition.Topic == "mysql_server_name.subaid_01.Drug" {
					if err := s.DrugHandler.TransformData(msg.Key, msg.Value); err != nil {
						commit(msg)
					}
				} else if *msg.TopicPartition.Topic == "mysql_server_name.subaid_01.Patients" {
					if err := s.UserHandler.TransformData(msg.Key, msg.Value); err != nil {
						commit(msg)
					}
				// } else if *msg.TopicPartition.Topic == "mysql_server_name.subaid_01.Prescriptions" {
				// 	if err := s.MedicationHandlers.TransformData(msg.Key, msg.Value); err != nil {
				// 		commit(msg)
				// 	}
				}
			}
		}

		log.Println("No more messages to consume. Exiting.")

		s.doneC <- struct{}{}
	}()

	return nil
}

//// Shutdown ...
func (s *Service) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server")

	close(s.closeC)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context.Done: %w", ctx.Err())

		case <-s.doneC:
			return nil
		}
	}
}
