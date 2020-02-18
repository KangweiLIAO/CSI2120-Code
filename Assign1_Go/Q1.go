// @Title  			Q1.go
// @Description 	A simple program to manage ticket sales for a theatre.
// @Author  		Kangwei Liao (8568800)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Play Struct: include basic information for Comedy and Tragedy
type Play struct {
	name      string
	purchased []Ticket
	showStart time.Time
	showEnd   time.Time
}

// Comedy Struct: embed Play struct and include additional information for Comedy
type Comedy struct {
	Play
	laughs float32
	deaths int32
}

// Tragedy Struct: embed Play struct and include additional information for Tragedy
type Tragedy struct {
	Play
	laughs float32
	deaths int32
}

// Show Interface: include methods for a Play
type Show interface {
	getName() string
	getShowStart() time.Time
	getShowEnd() time.Time
	addPurchase(*Ticket) bool
	isNotPurchased(*Ticket) bool
}

// Seat Struct: include information for a seat
type Seat struct {
	number int32
	row    int32
	cat    *Category
}

// Category Struct: include information for a category
type Category struct {
	name string
	base float32
}

// Ticket Struct: include information for a ticket
type Ticket struct {
	customer string
	s        *Seat
	show     *Show
}

// Theatre Struct: include information for a theatre
type Theatre struct {
	seats []Seat
	shows []Show
}

// @title    		getName()
// @description   	return the name of a play
// @return    		name string
func (p *Play) getName() string {
	return p.name
}

// @title    		getShowStart()
// @description   	return the start time of a play
// @return    		showStart time.Time
func (p *Play) getShowStart() time.Time {
	return p.showStart
}

// @title    		getShowEnd()
// @description   	return the end time of a play
// @return    		showEnd time.Time
func (p *Play) getShowEnd() time.Time {
	return p.showEnd
}

// @title    		addPurchase()
// @description   	add a ticket to a play
// @param 			*Ticket
// @return    		bool
func (p *Play) addPurchase(t *Ticket) bool {
	if p.isNotPurchased(t) {
		p.purchased = append(p.purchased, *t)
		return true
	}
	return false
}

// @title    		isNotPurchased()
// @description   	return true if the ticket is not purchase, otherwise false
// @param			*Ticket
// @return    		bool
func (p *Play) isNotPurchased(t *Ticket) bool {
	for _, i := range p.purchased {
		if (i.s.row == t.s.row) && (i.s.number == t.s.number) {
			return false
		}
	}
	return true
}

// @title    		NewSeat()
// @description   	return a seat info in a play
// @param			int32, int32, *Category
// @return    		time.Time
func NewSeat(num int32, row int32, category *Category) *Seat {
	return &Seat{num, row, category}
}

// @title    		NewTicket()
// @description   	return a ticket of a play
// @param			string, *Seat, *Show
// @return    		time.Time
func NewTicket(name string, seat *Seat, show *Show) *Ticket {
	return &Ticket{name, seat, show}
}

// @title    		NewTheatre()
// @description   	return a Theatre includes shows
// @param			int32, []Show
// @return    		*Theatre
func NewTheatre(numOfSeat int32, shows []Show) *Theatre {
	seats := make([]Seat, numOfSeat)
	return &Theatre{seats, shows}
}

// ---------- Additional Functions ----------

// @title    		NewComedy()
// @description   	return a default comedy show
// @return    		*Comedy
func NewComedy() *Comedy {
	comedy := new(Comedy)
	comedy.name = "Tartuffe"
	comedy.purchased = make([]Ticket, 0)
	comedy.showStart = time.Date(2020, 3, 3, 16, 0, 0, 0, time.UTC)
	comedy.showEnd = time.Date(2020, 3, 3, 17, 20, 0, 0, time.UTC)
	comedy.laughs = 0.2
	comedy.deaths = 0.0
	return comedy
}

// @title    		NewTheatre()
// @description   	return a default tragedy show
// @return    		*Tragedy
func NewTragedy() *Tragedy {
	tragedy := new(Tragedy)
	tragedy.name = "Macbeth"
	tragedy.purchased = make([]Ticket, 0)
	tragedy.showStart = time.Date(2020, 4, 16, 9, 30, 0, 0, time.UTC)
	tragedy.showEnd = time.Date(2020, 4, 16, 12, 30, 0, 0, time.UTC)
	tragedy.laughs = 0
	tragedy.deaths = 12
	return tragedy
}

// @title    		NewTheatre()
// @description   	return a default category of ticket
// @param			string, float32
// @return    		*Category
func NewCategory(name string, base float32) *Category {
	category := new(Category)
	category.name = name
	category.base = base
	return category
}

// @title    		ArrangeSeats()
// @description   	arrange a seats in a theatre by n*n
// @param			int
func (t *Theatre) arrangeSeats(numOfRow int) {
	seatNum, rowNum, perRow := 0, 0, cap(t.seats)/numOfRow
	for i, _ := range t.seats {
		if i%perRow == 0 {
			seatNum = 1
			rowNum += 1
		}
		t.seats[i].number = int32(seatNum)
		t.seats[i].row = int32(rowNum)
		t.seats[i].cat = NewCategory("Standard", 25.0)
		switch k := rowNum; k {
		case 1:
			t.seats[i].cat.name = "Prime"
			t.seats[i].cat.base = 35.0
		case 5:
			t.seats[i].cat.name = "Special"
			t.seats[i].cat.base = 15.0
		}
		seatNum++
	}
}

// @title    		GetTicket()
// @description   	ask for input for a new ticket in a theatre
// @return			*Ticket, int32, bool
func (t *Theatre) GetTicket() (*Ticket, int32, bool) {
	// initiate a seat and a ticket
	seat := new(Seat)
	ticket := new(Ticket)
	buffer := bufio.NewReader(os.Stdin)
	showChecked, seatChecked, ticketChecked := false, false, false
	seatNum, rowNum, showNum := int32(-1), int32(-1), int32(-1)
	totalSeats := int32(cap(t.seats))
	seatsPerRow := totalSeats / t.seats[totalSeats-1].row

	fmt.Print("Please enter customer's name: ")
	custName, _ := buffer.ReadString('\n')
	custName = strings.Replace(custName, "\n", "", -1)
	// check the input for show is valid or not
	for !showChecked {
		fmt.Print("Please enter the name of the show: ")
		showName, _ := buffer.ReadString('\n')
		showName = strings.Replace(showName, "\n", "", -1)
		for i, _ := range t.shows {
			if t.shows[i].getName() == showName {
				showNum = int32(i)
				showChecked = true
			}
		}
		if !showChecked {
			fmt.Print("Please check the name of existing show:\n")
			for n, s := range t.shows {
				fmt.Printf("    %d: %v\n", n+1, s.getName())
			}
		}
	}
	// check the input for seat is valid or not
	for !seatChecked {
		fmt.Print("Please enter the desired row number: ")
		fmt.Scanf("%d", &rowNum)
		fmt.Print("Please enter the desired seat number in row: ")
		fmt.Scanf("%d", &seatNum)
		for n, s := range t.seats {
			if (seatNum == s.number) && (rowNum == s.row) {
				seat = &t.seats[n]
				seatChecked = true
			}
		}
		if !seatChecked {
			fmt.Printf("The entered seat info is not valid, please try again. (%v rows, %v seats/row)\n",
				totalSeats/seatsPerRow, seatsPerRow)
		}
	}
	if showChecked && seatChecked {
		ticket = NewTicket(custName, seat, &t.shows[showNum])
		ticketChecked = true
	}
	return ticket, showNum, ticketChecked
}

// @title    		validateTicket()
// @description   	validate a ticket, return true if a seat is (re)assigned to ticket; false if all seats are taken.
// @return			*Ticket, int32, bool
func (t *Theatre) validateTicket(ticket *Ticket, showNum int32) (*Ticket, bool) {
	valid := t.shows[showNum].isNotPurchased(ticket) // check the seat is valid or not
	totalSeats := int32(cap(t.seats))
	seatsPerRow := totalSeats / t.seats[totalSeats-1].row
	var primeFull, standardFull, specialFull bool // if current category is full, reassign seat
	if !valid {                                   // reassign a seat for current ticket
		fmt.Printf("\nxxxxx The seat is already taken xxxxx")
		for {
			switch c := ticket.s.cat.name; c {
			case "Prime": // try reassign a prime seat for prime ticket
				for _, s := range t.seats[:seatsPerRow] {
					ticket.s = &s
					if t.shows[showNum].isNotPurchased(ticket) {
						fmt.Printf("\n----- Successfully reassigned a Prime seat -----\n")
						return ticket, true
					}
				}
				primeFull = true
			case "Standard": // try reassign a standard seat for standard ticket
				for _, s := range t.seats[seatsPerRow : totalSeats-seatsPerRow] {
					ticket.s = &s
					if t.shows[showNum].isNotPurchased(ticket) {
						fmt.Printf("\n----- Successfully reassigned a Standard seat -----\n")
						return ticket, true
					}
				}
				standardFull = true
			case "Special": // try reassign a special seat for special ticket
				for _, s := range t.seats[totalSeats-seatsPerRow:] {
					ticket.s = &s
					if t.shows[showNum].isNotPurchased(ticket) {
						fmt.Printf("\n----- Successfully reassigned a Special seat -----\n")
						return ticket, true
					}
				}
				specialFull = true
			}
			if !primeFull {
				ticket.s.cat.name = "Prime"
			} else if !standardFull {
				ticket.s.cat.name = "Standard"
			} else if !specialFull {
				ticket.s.cat.name = "Special"
			}
			if primeFull && standardFull && specialFull {
				return ticket, false // if all seats are taken, return false
			}
		}
	}
	return ticket, true // seat is valid, do not need to be reassigned
}

// ---------- Print Methods ----------

// @title    		printTicket()
// @description   	print ticket info
func (t *Ticket) printTicket() {
	show := *t.show
	fmt.Printf("%v Ticket - $%v:\n\tname: %v\n\tseat: row %d seat %d\n\tshow: %v\n",
		t.s.cat.name, t.s.cat.base, t.customer, t.s.row, t.s.number, show.getName())
}

// @title    		printAllSeatsInfo()
// @description   	print all seats info in a theatre
func (t *Theatre) printAllSeatsInfo() {
	for _, s := range t.seats {
		fmt.Printf("(%v) Seat %d (row: %d, price: %v): \n",
			s.cat.name, s.number, s.row, s.cat.base)
	}
}

// @title    		printAllShowsInfo()
// @description   	print all shows info in a theatre
func (t *Theatre) printAllShowsInfo() {
	for _, s := range t.shows {
		fmt.Printf("(Show) %v:\n\tStart Time: %v\n\tEnd Time: %v\n",
			s.getName(), s.getShowStart(), s.getShowEnd())
	}
}

// @title    		printTheatreInfo()
// @description   	print basic info of a theatre
func (t *Theatre) printTheatreInfo() {
	totalSeats := cap(t.seats)
	totalRows := t.seats[totalSeats-1].row
	totalShows := cap(t.shows)
	fmt.Printf("Theatre have %d seats in %d row(s), and %d show(s) in schedule.\n",
		totalSeats, totalRows, totalShows)
}

// ---------- Main Function ----------

func main() {

	// ----- Create Shows -----
	totalShows := 2
	shows := make([]Show, totalShows)
	comedy1 := NewComedy()
	comedy1.showStart = time.Date(2020, 3, 3, 19, 30, 0, 0, time.UTC)
	comedy1.showEnd = time.Date(2020, 3, 3, 22, 00, 0, 0, time.UTC)
	shows[0] = comedy1
	tragedy1 := NewTragedy()
	tragedy1.showStart = time.Date(2020, 4, 10, 20, 0, 0, 0, time.UTC)
	tragedy1.showEnd = time.Date(2020, 4, 10, 23, 0, 0, 0, time.UTC)
	shows[1] = tragedy1

	// ----- Create Theatre -----
	totalSeats := int32(25)
	ticketCount := make([]int32, totalShows)
	theatre := NewTheatre(totalSeats, shows)
	theatre.arrangeSeats(5)

	// ----- Check Theatre Info -----
	// theatre.printTheatreInfo()
	// theatre.printAllSeatsInfo()
	// theatre.printAllShowsInfo()

	for {
		var checked bool
		ticket, showNum, ticketCheck := theatre.GetTicket()
		if ticketCheck && !(totalSeats == ticketCount[showNum]) {
			// if ticket not sold out
			ticket, checked = theatre.validateTicket(ticket, int32(showNum))
			if checked {
				theatre.shows[showNum].addPurchase(ticket)
				ticketCount[showNum]++
				fmt.Printf("\n----- Purchase success! Following is the ticket info -----\n")
				ticket.printTicket()
			}
		}
		if !checked {
			fmt.Printf("\nxxxxx The play is full! Cannot assign more seats xxxxx\n")
		}
		// Purchased tickets info:
		// fmt.Printf("%s Purchased: %d\n%s Purchased: %d\n", theatre.shows[0].getName(), ticketCount[0],
		// 	theatre.shows[1].getName(), ticketCount[1])
		fmt.Print("\nPress any button to purchase next ticket...")
		fmt.Scanf("%v")
	}
}
