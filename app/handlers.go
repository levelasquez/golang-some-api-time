package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	location := "Etc/UTC"
	timezones := map[string]string{}
	tz := r.URL.Query().Get("tz")

	if tz != "" {
		sl := strings.Split(tz, ",")

		if len(sl) > 0 {
			for _, v := range sl {
				loc, err := time.LoadLocation(v)

				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)

					return
				}

				t := time.Now().In(loc).String()

				timezones[v] = t
			}

		} else {
			loc, err := time.LoadLocation(tz)

			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)

				return
			}

			t := time.Now().In(loc).String()

			timezones["current_time"] = t
		}

		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(timezones)

		return
	}

	loc, err := time.LoadLocation(location)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	t := time.Now().In(loc).String()
	timezones["current_time"] = t

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(timezones)
}
