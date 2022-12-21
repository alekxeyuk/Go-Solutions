package robot

// See defs.go for other definitions

// Step 1
// Define N, E, S, W here.
const (
	N Dir = iota
	E
	S
	W
)

func Left() {
	Step1Robot.Dir -= 1
	if Step1Robot.Dir < N {
		Step1Robot.Dir = W
	}
}

func Right() {
	Step1Robot.Dir += 1
	if Step1Robot.Dir > W {
		Step1Robot.Dir = N
	}
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
	case E:
		Step1Robot.X += 1
	case S:
		Step1Robot.Y -= 1
	case W:
		Step1Robot.X -= 1
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	}
	return "?"
}

// Step 2
// Define Action type here.
type Action Command

func (r *Step2Robot) Right() {
	r.Dir += 1
	if r.Dir > W {
		r.Dir = N
	}
}

func (r *Step2Robot) Left() {
	r.Dir -= 1
	if r.Dir < N {
		r.Dir = W
	}
}

func (r *Step2Robot) Advance(extent *Rect) {
	switch r.Dir {
	case N:
		if r.Northing < extent.Max.Northing {
			r.Northing += 1
		}
	case E:
		if r.Easting < extent.Max.Easting {
			r.Easting += 1
		}
	case S:
		if r.Northing > extent.Min.Northing {
			r.Northing -= 1
		}
	case W:
		if r.Easting > extent.Min.Easting {
			r.Easting -= 1
		}
	}
}

func StartRobot(command chan Command, action chan Action) {
	for cmd := range command {
		action <- Action(cmd)
	}
	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for act := range action {
		switch act {
		case 'R':
			robot.Right()
		case 'L':
			robot.Left()
		case 'A':
			robot.Advance(&extent)
		}
	}
	report <- robot
}

// Step 3
// Define Action3 type here.
type Action3 struct {
	Name string
	Command
}

type robotEntry struct {
	Ignored bool
	*Step2Robot
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "A robot without a name"
		action <- Action3{name, 'H'}
		return
	}
Out:
	for _, cmd := range script {
		switch cmd {
		case 'R', 'L', 'A':
			action <- Action3{name, Command(cmd)}
		default:
			log <- "An undefined command in a script"
			break Out
		}
	}
	action <- Action3{name, 'H'}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	halted := 0
	count := len(robots)
	robotsMap := make(map[string]*robotEntry)
	worldMap := make(map[Pos]bool)
	for i := 0; i < len(robots); i++ {
		if _, known := robotsMap[robots[i].Name]; known {
			log <- "Duplicate robot names"
			halted += 1
			continue
		}
		if exist := worldMap[robots[i].Step2Robot.Pos]; exist {
			log <- "Robots placed at the same place"
			halted += 1
			robotsMap[robots[i].Name] = &robotEntry{true, nil}
			continue
		}
		p := robots[i].Step2Robot.Pos
		if p.Northing > extent.Max.Northing || p.Easting > extent.Max.Easting || p.Northing < extent.Min.Northing || p.Easting < extent.Min.Easting {
			log <- "A robot placed outside of the room"
			halted += 1
			robotsMap[robots[i].Name] = &robotEntry{true, nil}
			continue
		}

		robotsMap[robots[i].Name] = &robotEntry{false, &robots[i].Step2Robot}
		worldMap[robots[i].Step2Robot.Pos] = true
	}

	for halted != count {
		act := <-action
		actRobot, known := robotsMap[act.Name]
		if !known {
			log <- "An action from an unknown robot"
			halted += 1
			continue
		}
		if actRobot.Ignored {
			continue
		}

		switch act.Command {
		case 'R':
			actRobot.Right()
		case 'L':
			actRobot.Left()
		case 'A':
			p := actRobot.Pos
			wb := false
			switch actRobot.Dir {
			case N:
				wb = worldMap[Pos{p.Easting, p.Northing + 1}]
			case E:
				wb = worldMap[Pos{p.Easting + 1, p.Northing}]
			case S:
				wb = worldMap[Pos{p.Easting, p.Northing - 1}]
			case W:
				wb = worldMap[Pos{p.Easting - 1, p.Northing}]
			}
			if !wb {
				delete(worldMap, p)
				actRobot.AdvanceWithLog(&extent, log)
				worldMap[actRobot.Pos] = true
			} else {
				log <- "Robot bump"
			}
		case 'H':
			actRobot.Ignored = true
			halted += 1
		}
	}
	rep <- robots
}

func (r *Step2Robot) AdvanceWithLog(extent *Rect, log chan string) {
	switch r.Dir {
	case N:
		if r.Northing < extent.Max.Northing {
			r.Northing += 1
		} else {
			log <- "Wall bump"
		}
	case E:
		if r.Easting < extent.Max.Easting {
			r.Easting += 1
		} else {
			log <- "Wall bump"
		}
	case S:
		if r.Northing > extent.Min.Northing {
			r.Northing -= 1
		} else {
			log <- "Wall bump"
		}
	case W:
		if r.Easting > extent.Min.Easting {
			r.Easting -= 1
		} else {
			log <- "Wall bump"
		}
	}
}
