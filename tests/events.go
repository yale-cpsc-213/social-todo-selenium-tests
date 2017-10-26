package tests

import (
	"net/url"
	"strconv"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests/selectors"
)

type Event struct {
	Title    string
	location string
	image    string
	year     string
	month    string
	day      string
	hour     string
	minute   string
	flaw     string
}

func randomEvent() Event {
	e := Event{
		Title:    randomdata.Adjective() + " " + randomdata.Noun(),
		location: randomdata.City() + ", " + randomdata.State(randomdata.Small),
		image:    "https://" + randomdata.Country(randomdata.FullCountry) + ".png", // ???
		year:     strconv.Itoa(randomdata.Number(2016, 2017)),
		month:    strconv.Itoa(randomdata.Number(11)),
		day:      strconv.Itoa(randomdata.Number(28)),
		hour:     strconv.Itoa(randomdata.Number(23)),
		minute:   strconv.Itoa(30),
	}
	return e
}

func (e Event) createFormData() map[string]string {
	data := map[string]string{
		selectors.NewEventTitle:    e.Title,
		selectors.NewEventImage:    e.image,
		selectors.NewEventLocation: e.location,
		selectors.NewEventYear:     e.year,
		selectors.NewEventMonth:    e.month,
		selectors.NewEventDay:      e.day,
		selectors.NewEventHour:     e.hour,
		selectors.NewEventMinute:   e.minute,
	}
	return data
}
func (e Event) getURLValues() url.Values {
	return url.Values{
		"title":    []string{e.Title},
		"location": []string{e.location},
		"image":    []string{e.image},
		"year":     []string{e.year},
		"month":    []string{e.month},
		"day":      []string{e.day},
		"minute":   []string{e.minute},
	}
}

func createFormDataAPITest() Event {
	e := Event{
		Title:    randomString(45),
		location: "New Haven, CT",
		image:    "https://upload.wikimedia.org/wikipedia/commons/thumb/1/1a/Branford_Court_spring_2.JPG/250px-Branford_Court_spring_2.JPG",
		year:     "2017",
		month:    "7",
		day:      "4",
		hour:     "8",
		minute:   "30",
	}
	return e
}

func getBadEvents() []Event {
	var e Event
	var events []Event

	e = randomEvent()
	e.flaw = "no title"
	e.Title = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "too-long title"
	e.Title = randomString(51)
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no image"
	e.image = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "image that is not a valid URL"
	e.image = "branford-friends-always.foo"
	events = append(events, e)

	e = randomEvent()
	e.flaw = "image that does not end with .png, .gif, or .jpg"
	e.image = "http://foo.com/branford-friends-always.foo"
	events = append(events, e)

	e = randomEvent()
	e.flaw = "no location"
	e.location = ""
	events = append(events, e)

	e = randomEvent()
	e.flaw = "too-long location"
	e.location = randomString(51)
	events = append(events, e)

	// all cases below aren't really possible with dropdown menus -- do we test anyway?

	// e = randomEvent()
	// e.flaw = "year too high"
	// e.year = strconv.Itoa(randomdata.Number(2018,2020))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "year too low"
	// e.year = strconv.Itoa(randomdata.Number(2010,2015))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "month too high"
	// e.month = strconv.Itoa(randomdata.Number(12,30))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "month too low"
	// e.month = strconv.Itoa(randomdata.Number(-11,-1))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "day too high"
	// e.day = strconv.Itoa(randomdata.Number(32,100))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "day too low"
	// e.day = strconv.Itoa(randomdata.Number(-33,-1))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "hour too high"
	// e.hour = strconv.Itoa(randomdata.Number(24,50))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "hour too low"
	// e.hour = strconv.Itoa(randomdata.Number(-20,-1))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "minute too high"
	// e.minute = strconv.Itoa(randomdata.Number(31,100))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "minute in middle"
	// e.minute = strconv.Itoa(randomdata.Number(1,29))
	// events = append(events, e)
	//
	// e = randomEvent()
	// e.flaw = "minute too low"
	// e.minute = strconv.Itoa(randomdata.Number(-27,-1))
	// events = append(events, e)

	return events
}
