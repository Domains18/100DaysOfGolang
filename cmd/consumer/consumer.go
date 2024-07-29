package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"net/http"
	"github.com/Domains18/kafaka-notify/pkg/models"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

const (
	ConsumerGroup = "notifications-group"
	ConsumerTopic = "notifications"
	ConsumePort   = ":8081"
	KafkaServerAddress = "localhost:9092"
)


var ErrNoMessagesFound = errors.New("no messages found")

func getUserIDFromRequest(ctx *gin.Context)(string, error){
	userID := ctx.Param("userID")
	if userID == "" {
		return "", ErrNoMessagesFound
	}
	return userID, nil
}

type UserNotifications map[string][]models.Notification


type NotificationStore struct {
	data UserNotifications
	mu sync.RWMutex
}

func(ns *NotificationStore) Add(userID string, notification models.Notification){
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.data[userID] = append(ns.data[userID], notification)
}

func (ns NotificationStore) Get(userID string) []models.Notification{
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.data[userID]
}

// kafka functions

type Consumer struct {
	store *NotificationStore
} 


func (*Consumer) Setup(sarama.ConsumerGroupSession) error {return nil}
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil}

func (consumer *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages(){
		userID := string(msg.Key)
		var notification models.Notification
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notifications: %v", err)
			continue
		}
		consumer.store.Add(userID, notification)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error){
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup([]string{KafkaServerAddress}, ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("faled to initialize consumer group: %w", err)
	}
	return consumerGroup, nil
}


func setupConsumerGroup(ctx context.Context, store *NotificationStore){
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("failed to initialize: %v", err)
	}
	defer  consumerGroup.Close()

	consumer := &Consumer{ store: store}

	for {
		err = consumerGroup.Consume(ctx, []string{ConsumerTopic}, consumer)
		if err != nil{
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}


func handleNotifications(ctx *gin.Context, store *NotificationStore){
	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message":  err.Error()})
		return
	}

	notes := store.Get(userID)
	if len(notes) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No notificatiobs found for user",
			"notifications": []models.Notification{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"notifications": notes})
}