package flyweight

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type ID uint64

type Team struct {
	ID         ID
	Name       string
	Shield     []byte
	Players    []Player
	Historical []HistoricalData
}

type Player struct {
	Name         string
	Surname      string
	PreviousTeam ID
	Photo        []byte
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorID     ID
	LocalID       ID
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

func NewTeamFactory() teamFlyweightFactory {
	return teamFlyweightFactory{
		make(map[int]*Team, 0),
	}
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *teamFlyweightFactory) GetTeam(teamName int) *Team {
	if t.createdTeams[teamName] != nil {
		return t.createdTeams[teamName]
	}

	team := getTeamFactory(teamName)
	t.createdTeams[teamName] = &team

	return t.createdTeams[teamName]
}

func (t *teamFlyweightFactory) GetNumberofObjects() int {
	return len(t.createdTeams)
}
