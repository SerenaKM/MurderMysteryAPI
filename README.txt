Title of the project: Using Go to create a Murder Mystery API

Name: Serena KM

Idea:
•	I want to create a Murder Mystery API which generates a case file for people to solve the case
•	I completed Frontend Masters' Complete Go for Professional Developers and want to put what I learned into practice

Requirements:
•	User requests a case number and program combines a case with randomised details in the database such as: crime scene and suspects
•	User sends guess and receives response about whether or not their guess was correct
•	Possible: user can "talk" to suspects, get clues

Entities:
•	Case -> {caseID, crimeSceneID, suspectIDs}
•	CrimeScene -> {crimeSceneID, location, description, keyObjects}
•	Suspects -> {suspectID, name, occupation, backstory, possibleMotive, relationshipToVictim, alibi}
•	Guess -> {caseID, whoDunIt, why}

API:
•	createCase() -> caseID
•	createCrimeScene() -> crimeSceneSummary
•	createSuspects() -> Suspects
•	guess() -> Guess

Endpoints:
•	PUT /NewCase
•	GET /CaseID
•	GET /CrimeScene
•	GET /Suspects {suspectID}
•	POST /Guess

Reflection (process, challenges, successes, learnings):
•	Learning about decoupling a database to future-proof against changing databases
•	Using Docker and a separate test database to test actual SQL queries and database interactions

List of dependencies for your project, including any libraries or packages the code relies on:
•	Chi router for HTTP services
•	pgx driver to connect to PostgreSQL
•	Goose for database migration