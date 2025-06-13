package main

import (
	"errors"
	"fmt"
	"gowatcher_g3/internal/checker"
	"sync"
)

func main() {
	targets := []string{
		"https://www.google.com",
		"https://www.notarealwebsite.abc",
		"https://github.com",
		"https://www.movie.database/film/details",
		"https://www.gaming.news/release/new-game",
		"https://www.health.clinic/appointment/online",
		"https://www.car.manufacturer/model/electric",
		"https://www.home.decor/ideas/living-room",
		"https://www.environmental.org/project/clean-water",
		"https://www.space.agency/mission/mars",
		"https://www.fashion.magazine/trend/summer",
		"https://www.tech.conference/schedule/day1",
		"https://www.food.blog/recipe/dessert",
		"https://www.online.course/programming/python",
		"https://www.travel.guide/city/paris",
		"https://www.music.label/artist/new-album",
		"https://www.sports.club/events/match",
		"https://www.photography.tips/technique/lighting",
		"https://www.diy.tools/review/drill",
		"https://www.pet.vet/service/vaccination",
		"https://www.gardening.store/seeds/flower",
		"https://www.finance.advice/retirement/planning",
		"https://www.history.podcast/episode/ww2",
		"https://www.language.exchange/partner/find",
		"https://www.book.review/author/classic",
		"https://www.movie.review/genre/comedy",
		"https://www.gaming.forum/topic/strategy",
	}
	// waitgroup compteur de goroutines en attente
	var wg sync.WaitGroup
	//n lui indiqu
	wg.Add(len(targets))

	for _, url := range targets {

		go func(u string) {
			defer wg.Done()
			result := checker.CheckURL(u)

			if result.Err != nil {
				var unreacheable *checker.UnreachableURLError

				if errors.As(result.Err, &unreacheable) {
					fmt.Printf("%s inaccessible : %v\n", unreacheable.URL, unreacheable.Err)
				}
			} else {
				fmt.Printf("OK %s : %v\n", result.Target, result.Status)
			}
		}(url)
	}
	wg.Wait()

}
