package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"
	//passport "IKY-telegram-bot/artifacts/TGPassport"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var yesNoKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Yes"),
		tgbotapi.NewKeyboardButton("No")),
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Verify personal wallet"),
		tgbotapi.NewKeyboardButton("Verify collective wallet")),
)

//to operate the bot, put a text file containing key for your bot acquired from telegram "botfather" to the same directory with this file
var tgApiKey, err = os.ReadFile(".secret")
var bot, error1 = tgbotapi.NewBotAPI(string(tgApiKey))

//type containing all the info about user input
type user struct {
	tgid                int64
	dialog_status            int64
	//exportTokenName   string
	//exportTokenSymbol string
	//exportTokenSupply uint64
	//exportTokenType   uint64
	//tokenTypeString   string
}

//main database, key (int64) is telegram user id
var userDatabase = make(map[int64]user)

var msgTemplates = make (map[string] string)

var baseURL = "http://localhost:3000/"
var tg_id_query = "?user_tg_id="

var myenv map[string]string

const envLoc = ".env"




func main() {


	loadEnv()
	ctx := context.Background()
	pk := myenv["PK"] // load private key from env

	msgTemplates["hello"] = "Hey, this bot is attaching personal wallets to telegram user & collective wallets to chat id"
	msgTemplates["case0"] = "Go to link and attach your tg_id to your metamask wallet"
	msgTemplates["case1"] = "Awaiting for verification"

	bot, err = tgbotapi.NewBotAPI(string(tgApiKey))
	if err != nil {
		log.Panic(err)
	}

		// Connecting to network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY"]) // load from local .env file
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	// setting up private key in proper format
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	// Creating an auth transactor
	auth := bind.NewKeyedTransactor(privateKey)

	// check calls
	// check balance
	accountAddress := common.HexToAddress("0x16d97A46030C5D3D705bca45439e48529997D8b2")
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance of the validator bot: %d\n", balance)

	// Setting up Passport Contract
	passportCenter, err := passport.NewPassport(common.HexToAddress("0xef2DCEA186FA90FcD0c44C3caaE81E4586eE0685"), client)
		if err != nil {
			log.Fatalf("Failed to instantiate a TGPassport contract: %v", err)
		}

		// Wrap the Passport contract instance into a session
		session := &passport.PassportSession{
			Contract: passportCenter,
			CallOpts: bind.CallOpts{
				Pending: true,
				From:    auth.From,
				Context: context.Background(),
			},
			TransactOpts: bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasLimit: 0,   // 0 automatically estimates gas limit
				GasPrice: nil, // nil automatically suggests gas price
				Context:  context.Background(),
			},
		}
	
	log.Printf("session with passport centr initialized")

	/*
	   Events
	*/

	// Check retriving events of CashOut request

	var ch = make(chan *passport.PassportPassportApplied)
	subscription, err := SubscribeForApplications(session,ch)
	
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {

		if update.Message != nil {
			if _, ok := userDatabase[update.Message.From.ID]; !ok {

				userDatabase[update.Message.From.ID] = user{update.Message.Chat.ID, 0,}
				msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["hello"])
				msg.ReplyMarkup = mainKeyboard
				bot.Send(msg)
			} else {

				switch userDatabase[update.Message.From.ID].dialog_status {

				//first check for user status, (for a new user status 0 is set automatically), then user reply for the first bot message is logged to a database as name AND user status is updated
				case 0:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						//updateDb.exportTokenName = update.Message.Text
						updateDb.dialog_status = 1
						userDatabase[update.Message.From.ID] = updateDb

						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid,msgTemplates["case0"] )
						bot.Send(msg)

						tgid := userDatabase[update.Message.From.ID].tgid
						tgid_string := fmt.Sprint(tgid)
						link := baseURL + tg_id_query + tgid_string
						msg = tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid,link)
						bot.Send(msg)

					//	msg = tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid)

					}

				//logic is that 1 incoming message fro the user equals one status check in database, so each status check ends with the message asking the next question
				case 1:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						//updateDb.exportTokenSymbol = update.Message.Text
						updateDb.dialog_status = 2
						userDatabase[update.Message.From.ID] = updateDb
					}
					msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["case1"])
					bot.Send(msg)


				}
			}
		}
	}
}

func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}}



// subscribing for Applications events
func SubscribeForApplications(session *passport.PassportSession, listenChannel chan *passport.PassportPassportApplied) (event.Subscription, error)  {
	ApplicationsFilter := session.Contract.FilterPassportApplied
	
	subscription, err := ApplicationsFilter(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil,
	}, listenChannel)
	if err != nil {
		return nil, err
	}
	
	return subscription, err
}
