
	if m.Content == "!steamsale" {
		s.ChannelMessageSend(m.ChannelID, "Скидки есть ✓ \nИгор нету ✓")
	}

	if strings.Contains(m.Content, "!cycle") {
		s.ChannelMessageSend(m.ChannelID, itter(m.Content))
	}

	if strings.Contains(m.Content, "!vanish") {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}

	if strings.Contains(m.Content, "!spoiler") {
		out := string("||" + m.Content + "||")
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, out)
	}

	if strings.Contains(m.Content, "!duck") {
		s.ChannelMessageSend(m.ChannelID, Duck(m.Content))
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}


##Token

id 646668550567755776
secret YGXASnubNry9RYDKn1dwgh-4mWPMSjQQ
token NjQ2NjY4NTUwNTY3NzU1Nzc2.XdUgTw.f9-gKkASsMpct00gSmUG5JmQUNA

Channel 1
Print channel id:  646675428123344909
Channel 2
Print channel id:  648412465557798922



	func discordmain() {
		// Create a new Discord session using the provided bot token.
		dg, err := discord.New("Bot " + Token)
		if err != nil {
			fmt.Println("error creating Discord session,", err)
			return
		}

		// Register the messageCreate func as a callback for MessageCreate events.
		dg.AddHandler(messageCreate)

		// Open a websocket connection to Discord and begin listening.
		err = dg.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
			return
		}
		// Wait here until CTRL-C or other term signal is received.
		fmt.Println("Bot is now running.  Press CTRL-C to exit.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		// Cleanly close down the Discord session.
		dg.Close()
	}