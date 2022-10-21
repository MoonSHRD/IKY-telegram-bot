# IKY-telegram-bot
Verifier bot for IKY (https://github.com/MoonSHRD/IKY)
Currently can be found in tg as @E_Passport_bot

This bot is comfortable telegram bot, which allow users to tether their personal wallets to their tg_id, using metamask

This bot is also protect from injections, so users can't attach wallet address to tgid they don't own.


Blockchain part is generated from IKY repo, see it for more details about contracts business logic

# How it works:
1. User create intent to attach ethereum wallet to their tgid
2. Bot generate link to IKY webapp, and start listining to events from it's contract with user tgid
3. User go to webapp page, make transaction, submitting it's tgid
4. Bot get event about it from blockchain
5. Bot verifies registration

## Deploy with Docker

1. Copy `.envExample` into `.env` and fill it with your values
2. Run with `docker-compose up -d`
3. Done
