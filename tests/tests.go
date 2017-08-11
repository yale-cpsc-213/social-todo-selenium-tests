package tests

import (
	"fmt"
	"log"
	"net/url"
	"time"

	goselenium "github.com/bunsenapp/go-selenium"
	"github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests/selectors"
)

func RunForURL(seleniumURL string, testURL string, failFast bool, sleepDuration time.Duration) (int, int, error) {
	// Create capabilities, driver etc.
	capabilities := goselenium.Capabilities{}
	capabilities.SetBrowser(goselenium.ChromeBrowser())

	driver, err := goselenium.NewSeleniumWebDriver(seleniumURL, capabilities)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	_, err = driver.CreateSession()
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	// Delete the session once this function is completed.
	defer driver.DeleteSession()

	return Run(driver, testURL, true, failFast, sleepDuration)
}

// Run - run all tests
//
func Run(driver goselenium.WebDriver, testURL string, verbose bool, failFast bool, sleepDuration time.Duration) (int, int, error) {
	numPassed := 0
	numFailed := 0
	doLog := func(args ...interface{}) {
		if verbose {
			fmt.Println(args...)
		}
	}
	die := func(msg string) {
		driver.DeleteSession()
		log.Fatalln(msg)
	}
	logTestResult := func(passed bool, err error, testDesc string) {
		doLog(statusText(passed && (err == nil)), "-", testDesc)
		if passed && err == nil {
			numPassed++
		} else {
			numFailed++
			if failFast {
				time.Sleep(5000 * time.Millisecond)
				die("Found first failing test, quitting")
			}
		}
	}

	users := []User{
		randomUser(),
		randomUser(),
		randomUser(),
	}

	getEl := func(sel string) (goselenium.Element, error) {
		return driver.FindElement(goselenium.ByCSSSelector(sel))
	}
	countCSSSelector := func(sel string) int {
		elements, xerr := driver.FindElements(goselenium.ByCSSSelector(sel))
		if xerr == nil {
			return len(elements)
		}
		return 0
	}
	cssSelectorExists := func(sel string) bool {
		count := countCSSSelector(sel)
		return (count != 0)
	}
	cssSelectorsExists := func(sels ...string) bool {
		for _, sel := range sels {
			if cssSelectorExists(sel) == false {
				return false
			}
		}
		return true
	}

	// Navigate to the URL.
	_, err := driver.Go(testURL)
	logTestResult(true, err, "Site is up and running")

	time.Sleep(sleepDuration)

	doLog("Home page:")

	result := cssSelectorExists(selectors.BootstrapHref)
	logTestResult(result, nil, "looks 💯 ")

	result := cssSelectorExists(selectors.Header)
	logTestResult(result, nil, "has a header")
	result := cssSelectorExists(selectors.Footer)
	logTestResult(result, nil, "has a footer")

	result := cssSelectorExists(selectors.FooterHomeLink)
	logTestResult(result, nil, "footer links to home page")
	result := cssSelectorExists(selectors.FooterAboutLink)
	logTestResult(result, nil, "footer links to about page")

	result := cssSelectorExists(selectors.TeamLogo)
	logTestResult(result, nil, "has your team logo")

	result := cssSelectorExists(selectors.EventList)
	logTestResult(result, nil, "shows a list of events")

	// need to test event details

	result := cssSelectorExists(selectors.NewEventLink)
	logTestResult(result, nil, "has a link to the new event page")


	_, err := driver.Go(testURL + "/about")

	if err != nil {
		doLog("Couldn't load /about, please try again")
	}
	else {
		doLog("About page:")
		time.Sleep(sleepDuration)

		result := cssSelectorExists(selectors.BootstrapHref)
		logTestResult(result, nil, "looks 💯 ")

		result := cssSelectorExists(selectors.Header)
		logTestResult(result, nil, "has a header")
		result := cssSelectorExists(selectors.Footer)
		logTestResult(result, nil, "has a footer")

		result := cssSelectorExists(selectors.FooterHomeLink)
		logTestResult(result, nil, "footer links to home page")
		result := cssSelectorExists(selectors.FooterAboutLink)
		logTestResult(result, nil, "footer links to about page")

		result := cssSelectorExists(selectors.Names)
		logTestResult(result, nil, "has your names")

		result := cssSelectorExists(selectors.Headshots)
		logTestResult(result, nil, "shows your headshots")
	}

	_, err := driver.Go(testURL + "/events/new")

	if err != nil {
		doLog("Couldn't load /events/new, please try again")
	}
	else {
		doLog("New event page:")
		time.Sleep(sleepDuration)

		result := cssSelectorExists(selectors.BootstrapHref)
		logTestResult(result, nil, "looks 💯 ")

		result := cssSelectorExists(selectors.Header)
		logTestResult(result, nil, "has a header")
		result := cssSelectorExists(selectors.Footer)
		logTestResult(result, nil, "has a footer")

		result := cssSelectorExists(selectors.FooterHomeLink)
		logTestResult(result, nil, "footer links to home page")
		result := cssSelectorExists(selectors.FooterAboutLink)
		logTestResult(result, nil, "footer links to about page")

		result := cssSelectorExists(selectors.NewEventForm)
		logTestResult(result, nil, "has a form for event submission")

		titleResult := cssSelectorExists(selectors.NewEventTitle)
		titleLabelResult := cssSelectorExists(selectors.NewEventTitleLabel)
		logTestResult(titleResult && titleLabelResult, nil, "has a correctly labeled title field")

		imageResult := cssSelectorExists(selectors.NewEventImage)
		imageLabelResult := cssSelectorExists(selectors.NewEventImageLabel)
		logTestResult(imageResult && imageLabelResult, nil, "has a correctly labeled image field")

		locationResult := cssSelectorExists(selectors.NewEventLocation)
		locationLabelResult := cssSelectorExists(selectors.NewEventLocationLabel)
		logTestResult(locationResult && locationLabelResult, nil, "has a correctly labeled location field")

		yearResult := cssSelectorExists(selectors.NewEventYear)
		yearLabelResult := cssSelectorExists(selectors.NewEventYearLabel)
		// check for correct year options
		logTestResult(yearResult && yearLabelResult, nil, "has a correctly labeled year field")

		monthResult := cssSelectorExists(selectors.NewEventMonth)
		monthLabelResult := cssSelectorExists(selectors.NewEventMonthLabel)
		// check for correct month options
		logTestResult(monthResult && monthLabelResult, nil, "has a correctly labeled month field")

		dayResult := cssSelectorExists(selectors.NewEventDay)
		dayLabelResult := cssSelectorExists(selectors.NewEventDayLabel)
		// check for correct day options
		logTestResult(dayResult && dayLabelResult, nil, "has a correctly labeled day field")

		hourResult := cssSelectorExists(selectors.NewEventHour)
		hourLabelResult := cssSelectorExists(selectors.NewEventHourLabel)
		// check for correct hour options
		logTestResult(hourResult && hourLabelResult, nil, "has a correctly labeled hour field")

		minuteResult := cssSelectorExists(selectors.NewEventMinute)
		minuteLabelResult := cssSelectorExists(selectors.NewEventMinuteLabel)
		// check for correct minute options
		logTestResult(yearResult && yearLabelResult, nil, "has a correctly labeled minute field")

		// submit a bunch of bad form data


	}

	_, err := driver.Go(testURL + "/events/0")

	if err != nil {
		doLog("Couldn't load /events/0, please try again")
	}
	else {
		doLog("Event 0:")
		time.Sleep(sleepDuration)

		result := cssSelectorExists(selectors.BootstrapHref)
		logTestResult(result, nil, "looks 💯 ")

		result := cssSelectorExists(selectors.Header)
		logTestResult(result, nil, "has a header")
		result := cssSelectorExists(selectors.Footer)
		logTestResult(result, nil, "has a footer")

		result := cssSelectorExists(selectors.FooterHomeLink)
		logTestResult(result, nil, "footer links to home page")
		result := cssSelectorExists(selectors.FooterAboutLink)
		logTestResult(result, nil, "footer links to about page")

		result := cssSelectorExists(selectors.EventTitle)
		logTestResult(result, nil, "has a title")
		result := cssSelectorExists(selectors.EventDate)
		logTestResult(result, nil, "has a date")
		result := cssSelectorExists(selectors.EventLocation)
		logTestResult(result, nil, "has a location")
		result := cssSelectorExists(selectors.EventImage)
		logTestResult(result, nil, "has an image")
		result := cssSelectorExists(selectors.EventAttendees)
		logTestResult(result, nil, "has a list of attendees")

		// RSVP test
	}

	_, err := driver.Go(testURL + "/events/1")

	if err != nil {
		doLog("Couldn't load /events/1, please try again")
	}
	else {
		doLog("Event 1:")
		time.Sleep(sleepDuration)

		result := cssSelectorExists(selectors.BootstrapHref)
		logTestResult(result, nil, "looks 💯 ")

		result := cssSelectorExists(selectors.Header)
		logTestResult(result, nil, "has a header")
		result := cssSelectorExists(selectors.Footer)
		logTestResult(result, nil, "has a footer")

		result := cssSelectorExists(selectors.FooterHomeLink)
		logTestResult(result, nil, "footer links to home page")
		result := cssSelectorExists(selectors.FooterAboutLink)
		logTestResult(result, nil, "footer links to about page")

		result := cssSelectorExists(selectors.EventTitle)
		logTestResult(result, nil, "has a title")
		result := cssSelectorExists(selectors.EventDate)
		logTestResult(result, nil, "has a date")
		result := cssSelectorExists(selectors.EventLocation)
		logTestResult(result, nil, "has a location")
		result := cssSelectorExists(selectors.EventImage)
		logTestResult(result, nil, "has an image")
		result := cssSelectorExists(selectors.EventAttendees)
		logTestResult(result, nil, "has a list of attendees")

		// RSVP test
	}

	_, err := driver.Go(testURL + "/events/2")

	if err != nil {
		doLog("Couldn't load /events/2, please try again")
	}
	else {
		doLog("Event 2:")
		time.Sleep(sleepDuration)

		result := cssSelectorExists(selectors.BootstrapHref)
		logTestResult(result, nil, "looks 💯 ")

		result := cssSelectorExists(selectors.Header)
		logTestResult(result, nil, "has a header")
		result := cssSelectorExists(selectors.Footer)
		logTestResult(result, nil, "has a footer")

		result := cssSelectorExists(selectors.FooterHomeLink)
		logTestResult(result, nil, "footer links to home page")
		result := cssSelectorExists(selectors.FooterAboutLink)
		logTestResult(result, nil, "footer links to about page")

		result := cssSelectorExists(selectors.EventTitle)
		logTestResult(result, nil, "has a title")
		result := cssSelectorExists(selectors.EventDate)
		logTestResult(result, nil, "has a date")
		result := cssSelectorExists(selectors.EventLocation)
		logTestResult(result, nil, "has a location")
		result := cssSelectorExists(selectors.EventImage)
		logTestResult(result, nil, "has an image")
		result := cssSelectorExists(selectors.EventAttendees)
		logTestResult(result, nil, "has a list of attendees")

		// RSVP test
	}

	// need to test API


	// OLD CODE
	welcomeCount := countCSSSelector(selectors.Welcome)
	logTestResult(welcomeCount == 0, nil, "should not be welcoming anybody b/c nobody is logged in!")

	doLog("When trying to register, your site")

	err = submitForm(driver, selectors.LoginForm, users[0].loginFormData(), selectors.LoginFormSubmit)
	time.Sleep(sleepDuration)
	result = cssSelectorExists(selectors.Errors)
	logTestResult(result, err, "should not allow unrecognized users to log in")

	badUsers := getBadUsers()
	for _, user := range badUsers {
		msg := "should not allow registration of a user with " + user.flaw
		err2 := registerUser(driver, testURL, user)
		time.Sleep(sleepDuration)
		if err2 == nil {
			result = cssSelectorExists(selectors.Errors)
		}
		logTestResult(result, err2, msg)
	}

	err = registerUser(driver, testURL, users[0])
	if err == nil {
		time.Sleep(sleepDuration)
		result = cssSelectorExists(selectors.Welcome)
	}
	logTestResult(result, err, "should welcome users that register with valid credentials")

	el, err := getEl(".logout")
	result = false
	if err == nil {
		el.Click()
		var response *goselenium.CurrentURLResponse
		response, err = driver.CurrentURL()
		if err == nil {
			var parsedURL *url.URL
			parsedURL, err = url.Parse(response.URL)
			if err == nil {
				result = parsedURL.Path == "/"
				if result {
					time.Sleep(sleepDuration)
					result = cssSelectorsExists(selectors.LoginForm, selectors.RegisterForm)
				}
			}
		}
	}
	logTestResult(result, err, "should redirect users to '/' after logout")

	logout := func() {
		element, _ := getEl(".logout")
		result = false
		if err == nil {
			element.Click()
		}
	}

	// Register the other two users
	err = registerUser(driver, testURL, users[1])
	if err != nil {
		die("Error registering second user")
	}
	logout()
	err = registerUser(driver, testURL, users[2])
	if err != nil {
		die("Error registering third user")
	}
	logout()

	fmt.Println("A newly registered user")
	err = loginUser(driver, testURL, users[0])
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.Welcome) == 1, err, "should be able to log in again")

	numTasks := countCSSSelector(selectors.Task)
	logTestResult(numTasks == 0, nil, "should see no tasks at first")

	numTaskForms := countCSSSelector(selectors.TaskForm)
	logTestResult(numTaskForms == 1, nil, "should see a form for submitting tasks")

	badTasks := getBadTasks()
	for _, task := range badTasks {
		msg := "should not be able to create a task with " + task.flaw
		err2 := submitTaskForm(driver, testURL, task)
		var count int
		if err2 == nil {
			time.Sleep(sleepDuration)
			count = countCSSSelector(selectors.Errors)
		}
		logTestResult(count == 1, err2, msg)
	}

	task := randomTask(false)
	task.collaborator1 = users[1].email
	err = submitTaskForm(driver, testURL, task)
	time.Sleep(sleepDuration)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 1, err, "should see a task after a valid task is submitted")

	task = randomTask(false)
	err = submitTaskForm(driver, testURL, task)
	time.Sleep(sleepDuration)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 2, err, "should see two tasks after another is submitted")
	time.Sleep(3000 * time.Millisecond)

	logout()
	fmt.Println("User #2, after logging in")
	_ = loginUser(driver, testURL, users[1])
	time.Sleep(sleepDuration)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 1, err, "should be able to see the task that was shared with her by user #1")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskDelete) == 0, err, "should not be not prompted to delete that task (she's not the owner)")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskCompleted) == 0, err, "should see the task as incomplete")
	logTestResult(numTasks == 1 && countCSSSelector(selectors.TaskToggle) == 1, err, "should be able to mark the the task as complete")
	el, err = getEl(selectors.TaskToggle)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.TaskCompleted) == 1, err, "should see the task as complete after clicking the \"toggle\" action")
	logout()

	_ = loginUser(driver, testURL, users[0])
	fmt.Println("User #1, after logging in")
	time.Sleep(sleepDuration)
	numCompleted := countCSSSelector(selectors.TaskCompleted)
	numTasks = countCSSSelector(selectors.Task)
	logTestResult(numTasks == 2 && numCompleted == 1, err, "should see one of the two tasks marked as completed")
	el, err = getEl(selectors.TaskToggle)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.TaskCompleted) == 0, err, "should be able to mark that is incompleted when she clicks the \"toggle\" action")
	logTestResult(countCSSSelector(selectors.TaskDelete) == 2, err, "should be prompted to delete both tasks (she's the owner)")
	el, err = getEl(selectors.TaskDelete)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(countCSSSelector(selectors.Task) == 1, err, "should only see one after deleting a task")
	numTasks = countCSSSelector(selectors.Task)
	el, err = getEl(selectors.TaskDelete)
	el.Click()
	time.Sleep(sleepDuration)
	logTestResult(numTasks == 1 && countCSSSelector(selectors.Task) == 0, err, "should see none after deleting two tasks")

	return numPassed, numFailed, err
}
