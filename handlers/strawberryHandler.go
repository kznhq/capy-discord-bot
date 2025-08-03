package handlers

import (
	"context"
	"os"
	"math/rand"

	"github.com/kznhq/utils"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/bwmarrin/discordgo"
)

// inside joke but basically when this one person sends a message, after a random number of messages this will send an image from an S3 bucket in reply to that person
func StrawberryHandler(s *discordgo.Session, message *discordgo.MessageCreate) {
	bucket := os.Getenv("BUCKET")
	strawberry := os.Getenv("STRAWBERRY")
	if message.Member != strawberry { // this only applies to one specific member
		return
	}

	utils.M.Lock()
	utils.StrawberryCounter += 1
	if utils.StrawberryCounter < utils.StrawberryLimit {
		utils.M.Unlock()
		return
	}
	utils.StrawberryLimit = rand.Intn(100) + 51 // decided to make it 51-150 messages that need to be sent before the image shows up
	utils.M.Unlock()

	// at this point, the person has sent the random number of messages since the last time this handler ran so now we send another
	cfg, _ := config.LoadDefaultConfig(context.TODO())

	client := s3.NewFromConfig(cfg)

	// we get all the images and pick a random one to use for the key as the image to reply to
	listOutput, _ := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucket,
	})
	key := *listOutput.Contents[rand.Intn(len(listOutput.Contents))].Key

	// get the object that we will send
	output, _ := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key: &key,
	})
	defer output.Body.Close()

	// package the file that we will send
	file := &discordgo.File{
		Name: key,
		Reader: output.Body,
	}
	
	// reply to the person's message with this image
	s.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
		Files: []*discordgo.File{file},
		Reference: &discordgo.MessageReference {
			MessageID: message.ID,
			ChannelID: message.ChannelID,
			GuildID: message.GuildID,
		},
	})

}
