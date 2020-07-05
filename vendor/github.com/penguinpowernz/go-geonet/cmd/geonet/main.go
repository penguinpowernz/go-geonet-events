package main

import (
	"flag"
	"fmt"

	geonet "github.com/penguinpowernz/go-geonet"
)

func main() {

	var printQuakes int
	var printVolcanes bool
	var printQuakeInfo string
	var printFeltReports string

	flag.IntVar(&printQuakes, "q", -2, "print latest quakes (with given MMI)")
	flag.BoolVar(&printVolcanes, "v", false, "print volcano alert status")
	flag.StringVar(&printQuakeInfo, "i", "", "get quake info for given publicID")
	flag.StringVar(&printFeltReports, "r", "", "get quake felt reports for given publicID")
	flag.Parse()

	cl := geonet.NewClient()

	switch {
	case printFeltReports != "":
		q, err := cl.Quake(printFeltReports)
		if err != nil {
			panic(err)
		}

		reports, err := q.FeltReports(cl)
		if err != nil {
			panic(err)
		}

		if len(reports) == 0 {
			fmt.Println("No felt reports for", q.PublicID)
		}

		fmt.Println("ID,MMI,Count,Lat,Lon")
		for _, rep := range reports {
			fmt.Printf("%s,%d,%d,%v,%v\n", printFeltReports, rep.MMI, rep.Count, rep.Coordinates[0], rep.Coordinates[1])
		}

	case printQuakeInfo != "":
		q, err := cl.Quake(printQuakeInfo)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%10s: %s\n", "ID", q.PublicID)
		fmt.Printf("%10s: https://www.geonet.org.nz/earthquake/%s\n", "Link", q.PublicID)
		fmt.Printf("%10s: %s\n", "Time", q.Time)
		fmt.Printf("%10s: %s\n", "Locality", q.Locality)
		fmt.Printf("%10s: %v, %v\n", "Coords", q.Coordinates[0], q.Coordinates[1])
		fmt.Printf("%10s: %0.1f\n", "Magnitude", q.Magnitude)
		fmt.Printf("%10s: %0.1f\n", "Depth", q.Depth)
		fmt.Printf("%10s: %d\n", "MMI", q.MMI)
		fmt.Printf("%10s: %s\n", "Quality", q.Quality)

	case printQuakes > -2:
		quakes, err := cl.Quakes(printQuakes)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID,Locality,Depth,Mag,MMI,Time,Link")
		for _, q := range quakes {
			fmt.Printf("%s,%s,%0.2fkm,%0.2f,%d,%s,https://www.geonet.org.nz/earthquake/%s\n", q.PublicID, q.Locality, q.Depth, q.Magnitude, q.MMI, q.Time, q.PublicID)
		}

	case printVolcanes:
		vs, err := cl.Volcanos()
		if err != nil {
			panic(err)
		}

		fmt.Println("Title,Level,Activity,Hazard,Link")
		for _, v := range vs {
			fmt.Printf("%s,%d,%s,%s,https://www.geonet.org.nz/volcano/%s\n", v.Title, v.Level, v.Activity, v.Hazards, v.ID)
		}

	}

}
