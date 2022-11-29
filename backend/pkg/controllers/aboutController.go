/** @file aboutController.go
 * @brief This file contains all the functions for the about.json feature.
 * @author Timothee de Boynes
 */

// @cond

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type Service struct {
	Name    string `json:"name"`
	Actions []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		FieldNames  []string `json:"field_names"`
	} `json:"actions"`
	Reactions []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		FieldNames  []string `json:"field_names"`
	} `json:"reactions"`
}

type AboutJson struct {
	Client struct {
		Host string `json:"host"`
	} `json:"client"`
	Server struct {
		CurrentTime uint      `json:"current_time"`
		Services    []Service `json:"services"`
	} `json:"server"`
}

var Services []Service

// @endcond

/** @brief Gets the client IP adress from the request and returns it
 *
 * @param r *http.Request
 *
 * @return string, error
 */
func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

/** @brief Fills the global Servives if it is empty
 *
 */
func FillServices() {
	if Services != nil {
		return
	}

	data, err := os.ReadFile("pkg/controllers/services.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	JsonErr := json.Unmarshal([]byte(data), &Services)
	if JsonErr != nil {
		fmt.Fprintln(os.Stderr, JsonErr)
	}
}

/** @brief Writes all the information needed for the about.json in the response 
 *
 * @param w http.ResponseWriter, r *http.Request
 */
func GetAboutJson(w http.ResponseWriter, r *http.Request) {
	var aboutJson AboutJson

	ip, _ := getIP(r)
	aboutJson.Client.Host = ip

	now := time.Now()
	aboutJson.Server.CurrentTime = uint(now.Unix())

	FillServices()
	aboutJson.Server.Services = Services

	w.WriteHeader(http.StatusOK)
	js, _ := json.MarshalIndent(aboutJson, "", " ")
	w.Write(js)
}

/** @brief Returns all the services
 *
 * @param w http.ResponseWriter, r *http.Request
 */
func GetServices(w http.ResponseWriter, r *http.Request) {
	FillServices()
	w.WriteHeader(http.StatusOK)
	js, _ := json.MarshalIndent(Services, "", " ")
	w.Write(js)
}
