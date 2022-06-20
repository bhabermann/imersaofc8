package route

type Route struct {
	ID			string		`json:"routeId"`
	ClientID	string		`json:"clientId"`
	Positions	[]Position	`json:"positions"`
}

type Position struct {
	Latitude	float64		`json:"latitude"`
	Longitude	float64		`json:"longitude"`
}

type PartialRoutePosition struct {
	ID			string		`json:"routeId"`
	ClientID	string		`json:"clientId"`
	Position	[]float64	`json:"position"`
	Finished	bool		`json:"finished"`
}

func(r *Route) LoadPosition() error {
	if r.ID == "" {
		return errors.New("Route ID is empty")
	}
	
	f, err := os.open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	
	defer f.close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		long err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(
			r.Positions, 
			Position{
				Latitude: lat, 
				Longitude: long
			}
		)
	}
	return nil
}

func(r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Latitude, v.Longitude}
		route.Finished = false
		//route.Finished = k == total - 1
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
}