User story

A user will do:
Create an event
Join an event
Both: should select a date/time
Manage his event
Share an event

Models:
User - has -> * Event
Event - has -> * Participant
An User is a Participant

A MongoDB design should be
Event: {
	_id: hex,
	title: String,
	location: {
		name: String,
		geo: [Int, Int]
	},
	participants: [
		@User,
		...
	],
	admin: [
		@User,
		…
	],
	timeProprosal: [
		{
			begin: DateTime,
			end: DateTime,
			subscriber: [
				@User,
				...
			]
		},
		...
	]
}

User: {
	_id: hex, 
	name: String,
	email: String
}

Workflow:
Creation:
Guest creates an event , filling the form
Generate an unique id so as url for this event
Render the page
Join:
Navigate to url
Fill name and select several dates
Edit an Event:
Should have an Admin Page(url) for editing



