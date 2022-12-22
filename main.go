package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"os"

	"github.com/joho/godotenv"

	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"
	//passport "IKY-telegram-bot/artifacts/TGPassport"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var yesNoKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Yes"),
		tgbotapi.NewKeyboardButton("No")),
)

var optionKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("WhoIs")),
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Verify personal wallet")),
)

// to operate the bot, put a text file containing key for your bot acquired from telegram "botfather" to the same directory with this file
var tgApiKey, err = os.ReadFile(".secret")
var bot, error1 = tgbotapi.NewBotAPI(string(tgApiKey))

// type containing all the info about user input
type user struct {
	tgid          int64
	tg_username   string
	dialog_status int64
}

// event we got from blockchain
type event_bc = *passport.PassportPassportApplied

// channel to get this event from blockchain
var ch = make(chan *passport.PassportPassportApplied)
var ch_index = make(chan *passport.PassportPassportAppliedIndexed)

var ch_approved = make(chan *passport.PassportPassportApproved)

// main database for dialogs, key (int64) is telegram user id
var userDatabase = make(map[int64]user) // consider to change in persistend data storage?

var msgTemplates = make(map[string]string)

var tg_id_query = "?user_tg_id="
var tg_username_query = "&user_tg_name="

var myenv map[string]string

// file with settings for enviroment
const envLoc = ".env"

func main() {

	loadEnv()
	ctx := context.Background()
	pk := myenv["PK"] // load private key from env

	msgTemplates["hello"] = "Hey, this bot is attaching personal wallets to telegram user & collective wallets to chat id"
	msgTemplates["case0"] = "Open following link in metamask broswer"
	msgTemplates["await"] = "Awaiting for verification"
	msgTemplates["case1"] = "You have successfully authorized your wallet to your account. Now you can use additional functions"
	msgTemplates["who_is"] = "Input wallet address to know it's associated telegram nickname"

	//var baseURL = "http://localhost:3000/"
	//var baseURL = "https://ikytest-gw0gy01is-s0lidarnost.vercel.app/"
	var baseURL = myenv["BASEURL"]

	bot, err = tgbotapi.NewBotAPI(string(tgApiKey))
	if err != nil {
		log.Panic(err)
	}

	// Connecting to blockchain network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY_POLYGON_WS"]) // load from local .env file
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
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(137))

	// check calls
	// check balance
	accountAddress := common.HexToAddress(myenv["ACCOUNT_ADDRESS"])
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance of the validator bot: %d\n", balance)

	// Setting up Passport Contract
	passportCenter, err := passport.NewPassport(common.HexToAddress(myenv["PASSPORT_ADDRESS"]), client)
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

	log.Printf("session with passport center initialized")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {

		if update.Message != nil {
			if _, ok := userDatabase[update.Message.From.ID]; !ok {

				userDatabase[update.Message.From.ID] = user{update.Message.Chat.ID, update.Message.Chat.UserName, 0}
				msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["hello"])
				msg.ReplyMarkup = mainKeyboard
				bot.Send(msg)
				// check for registration
				registred := IsAlreadyRegistred(session, update.Message.From.ID)
				if registred {
					userDatabase[update.Message.From.ID] = user{update.Message.Chat.ID, update.Message.Chat.UserName, 1}
				}

			} else {

				switch userDatabase[update.Message.From.ID].dialog_status {

				//first check for user status, (for a new user status 0 is set automatically), then user reply for the first bot message is logged to a database as name AND user status is updated
				case 0:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {

						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["case0"])
						bot.Send(msg)

						tgid := userDatabase[update.Message.From.ID].tgid
						user_name := userDatabase[update.Message.From.ID].tg_username
						fmt.Println(user_name)
						tgid_string := fmt.Sprint(tgid)
						tgid_array := make([]int64, 1)
						tgid_array[0] = tgid
						link := baseURL + tg_id_query + tgid_string + tg_username_query + "@" + user_name
						msg = tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, link)
						bot.Send(msg)

						//subscription, err := SubscribeForApplications(session, ch)   //  this is ordinary subscription to NORMAL event
						subscription, err := SubscribeForApplicationsIndexed(session, ch_index, tgid_array) // this is subscription to INDEXED event. This mean we can pass what exactly value of argument we want to see
						if err != nil {
							log.Println(err)
						}

						go AsyncApproveChain(ctx, subscription, update.Message.From.ID, auth, passportCenter, session, userDatabase)

						updateDb.dialog_status = 4
						userDatabase[update.Message.From.ID] = updateDb

					}
				//	fallthrough // МЫ ЛЕД ПОД НОГАМИ МАЙОРА!
				case 1:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["case1"])
						msg.ReplyMarkup = optionKeyboard
						bot.Send(msg)
						updateDb.dialog_status = 2
						userDatabase[update.Message.From.ID] = updateDb

					}

				case 2:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						if update.Message.Text == "WhoIs" {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["who_is"])
							msg.ReplyMarkup = optionKeyboard
							bot.Send(msg)
							updateDb.dialog_status = 3
							userDatabase[update.Message.From.ID] = updateDb
						}
					}

				// whois
				case 3:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						var address_to_check_string = update.Message.Text
						var address_to_check = common.HexToAddress(address_to_check_string)
						tg_nickname, err := WhoIsAddress(session, address_to_check)
						if err != nil {
							log.Println("error with getting nickname associated with eth wallet, probably not registred yet")
						}
						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, tg_nickname)
						bot.Send(msg)
						updateDb.dialog_status = 2
						userDatabase[update.Message.From.ID] = updateDb

					}

				//
				case 4:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["await"])
						bot.Send(msg)
						updateDb.dialog_status = 4
						userDatabase[update.Message.From.ID] = updateDb
					}

				}
			}
		}
	}

} // end of main func

// load enviroment variables from .env file
func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

func AsyncApproveChain(ctx context.Context, subscription event.Subscription, tgid int64, auth *bind.TransactOpts, passportCenter *passport.Passport, session *passport.PassportSession, userDatabase map[int64]user) {
EventLoop:
	for {
		select {
		case <-ctx.Done():
			{
				subscription.Unsubscribe()
				break EventLoop
			}
		case eventResult := <-ch_index:
			{
				fmt.Println("User tg_id:", eventResult.ApplyerTg)
				fmt.Println("User wallet address:", eventResult.WalletAddress)
				applyer_tg_string := fmt.Sprint(eventResult.ApplyerTg)
				msg := tgbotapi.NewMessage(userDatabase[tgid].tgid, " your application have been recived "+applyer_tg_string)
				bot.Send(msg)
				ApprovePassport(auth, passportCenter, eventResult.WalletAddress)
				subscriptionApproved, err := SubscribeForApprovals(session, ch_approved)
				if err != nil {
					log.Fatal(err)
				}
			ApproveLoop:
				for {
					select {
					case <-ctx.Done():
						{
							subscriptionApproved.Unsubscribe()
							break ApproveLoop
						}
					case eventResult2 := <-ch_approved:
						{
							fmt.Println("User tg_id:", eventResult2.ApplyerTg)
							fmt.Println("User wallet address:", eventResult2.WalletAddress)
							approved_tg_string := fmt.Sprint(eventResult2.ApplyerTg)
							if approved_tg_string == applyer_tg_string {
								msg = tgbotapi.NewMessage(userDatabase[tgid].tgid, " your application have been APPROVED "+approved_tg_string)
								bot.Send(msg)
								subscriptionApproved.Unsubscribe()
								break ApproveLoop
							}

						}
					}
				}

				subscription.Unsubscribe()
				break EventLoop
			}
		}
	}
	updateDb := userDatabase[tgid]
	updateDb.dialog_status = 1
	userDatabase[tgid] = updateDb
}

// subscribing for Applications events. We use watchers without fast-forwarding past events
func SubscribeForApplications(session *passport.PassportSession, listenChannel chan<- *passport.PassportPassportApplied) (event.Subscription, error) {
	subscription, err := session.Contract.WatchPassportApplied(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

// subscribing for Applications events. We use watchers without fast-forwarding past events
func SubscribeForApplicationsIndexed(session *passport.PassportSession, listenChannel chan<- *passport.PassportPassportAppliedIndexed, applierTGID []int64) (event.Subscription, error) {
	subscription, err := session.Contract.WatchPassportAppliedIndexed(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
		applierTGID,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

// subscribing for APPROVAL events. We use watchers without fast-forwarding past events
func SubscribeForApprovals(session *passport.PassportSession, listenChannel chan<- *passport.PassportPassportApproved) (event.Subscription, error) {
	subscription, err := session.Contract.WatchPassportApproved(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

func ApprovePassport(auth *bind.TransactOpts, pc *passport.Passport, user_address common.Address) {

	tx_to_approve, err := pc.ApprovePassport(
		&bind.TransactOpts{
			From:      auth.From,
			Nonce:     nil,
			Signer:    auth.Signer,
			Value:     big.NewInt(0),
			GasPrice:  nil,
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   context.Background(),
		}, user_address,
	)

	if err != nil {
		log.Println("cant send approval reques to contract: ")
		log.Print(err)
	}

	fmt.Printf("transaction for APPROVAL passport sent! Please wait for tx %s to be confirmed. \n", tx_to_approve.Hash().Hex())

}

func DeclinePassport(auth *bind.TransactOpts, pc *passport.Passport, user_address common.Address) {

	tx_to_approve, err := pc.DeclinePassport(
		&bind.TransactOpts{
			From:      auth.From,
			Nonce:     nil,
			Signer:    auth.Signer,
			Value:     big.NewInt(0),
			GasPrice:  nil,
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   context.Background(),
		}, user_address,
	)
	if err != nil {
		log.Println("cant send DECLINING reques to contract: ")
		log.Print(err)
	}
	fmt.Printf("transaction for DECLINING passport sent! Please wait for tx %s to be confirmed. \n", tx_to_approve.Hash().Hex())
}

// allow bot to get tg nickname associated with this eth wallet
func WhoIsAddress(session *passport.PassportSession, address_to_check common.Address) (string, error) {
	passport, err := session.GetPassportByAddress(address_to_check)
	if err != nil {
		log.Println("cant get passport associated with this address, possible it's not registred yet: ")
		log.Print(err)
		return "error", err
	}
	nickname := passport.UserName
	return nickname, nil

}

func IsAlreadyRegistred(session *passport.PassportSession, user_id int64) bool {
	//GetPassportWalletByID
	passport_address, err := session.GetPassportWalletByID(user_id)
	if err != nil {
		return false
	}
	log.Println("check that user with this id:", user_id)
	log.Println("have associated wallet address:", passport_address)
	if passport_address == common.HexToAddress("0x0000000000000000000000000000000000000000") {
		log.Println("passport is null, user is not registred")
		return false
	} else {
		return true
	}
}

// generate link to trust page
func TrustUserLink(tgid_string string, friend_tg_id string, friend_username string) string {
	//http://localhost:3000/trust?user_tg_id=1337&friend_tg_id=1997&friend_user_name=sbekket
	loadEnv()
	baseURL_ := myenv["BASEURL"]
	trustURL := "/trust"
	tg_id_query_trust := "?user_tg_id="
	//	tg_username_query_trust := "&user_tg_name="
	to_id_query := "&friend_tg_id="
	to_username_query := "&friend_user_name="
	link := baseURL_ + trustURL + tg_id_query_trust + tgid_string + to_id_query + friend_tg_id + to_username_query + "@" + friend_username
	return link

}
