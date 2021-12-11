package reader

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/reader/interfaces"
	chaos "liquide/re/popularity-leaderboard-builder/reader/src/chaosMonkey"
	chaosLib "liquide/re/popularity-leaderboard-builder/reader/src/chaosMonkey/lib"
)

func ReaderProvider() (interfaces.UserActionReader, error) {

	store := chaosLib.ItemStore{}
	itemBuilder := chaosLib.ItemBuilder{
		Store: &store,

		ItemTypes: []string{"Trade", "Advise", "Post"},
		Topics:    []string{"NIFTY", "BANKNIFTY", "OMICRON", "COVID", "BUY", "SELL", "OPTIONS"},
		Stocks:    []string{"RELIANCE", "INFY", "TCS", "LT", "AIRTEL", "GNA", "SOBHA", "KEI", "CIPLA"},
		Authors:   []string{"NinjaTrader", "Warlord", "BigBoss", "Ajay", "Ramesh", "Suresh", "Liquide", "Anand Rathi", "Anant Ladha"},
		Channels: []string{
			"ninjatrader.advise",
			"ninjatrader.trade",
			"anandrathi.trade.portfolio.open",
			"anandrathi.trade.portfolio.growth_stocks",
			"anandrathi.advise",
			"anantladha.trade.portfolio.open",
			"anantladha.trade.portfolio.super_2021",
			"anantladha.advise",
			"anantladha.post",
			"liquide.advise.open",
			"liquide.trade.portfolio.super_liquide",
			"others.advise",
			"others.post",
		},
	}
	actionBuilder := chaosLib.ActionBuilder{
		Store: &store,

		ActionTypes: []string{"Trade", "Comment", "Rate"},
		Users: []objects.User{
			{
				Name:   "RP",
				Cohert: "CH 1",
			},
			{
				Name:   "Shreyas",
				Cohert: "CH 1",
			}, {
				Name:   "Madhu",
				Cohert: "CH 1",
			}, {
				Name:   "Ashu",
				Cohert: "CH 2",
			}, {
				Name:   "ARB",
				Cohert: "CH 2",
			}, {
				Name:   "Ajay",
				Cohert: "CH 3",
			}, {
				Name:   "Suresh",
				Cohert: "CH 3",
			},
		},
	}

	return chaos.ChaosMonkeyReader{
		ItemBuilder:   itemBuilder,
		ActionBuilder: actionBuilder,
	}, nil
}
