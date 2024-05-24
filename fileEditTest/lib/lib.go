package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

// Define a struct to hold the overall YAML structure
type Config struct {
	SliderMapping map[int]interface{} `yaml:"slider_mapping"`
}

func AddWindowToSlider(windowTitle string, index int) {
	// Read the YAML file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// Unmarshal the YAML data into a generic map
	var genericData map[string]interface{}
	err = yaml.Unmarshal(data, &genericData)
	if err != nil {
		log.Fatalf("error unmarshaling file: %v", err)
	}

	// Extract the slider_mapping section
	var sliderMapping map[int]interface{}
	if sm, exists := genericData["slider_mapping"]; exists {
		// Convert to the expected type
		smBytes, err := yaml.Marshal(sm)
		if err != nil {
			log.Fatalf("error marshaling slider_mapping: %v", err)
		}
		err = yaml.Unmarshal(smBytes, &sliderMapping)
		if err != nil {
			log.Fatalf("error unmarshaling slider_mapping: %v", err)
		}
	} else {
		sliderMapping = make(map[int]interface{})
	}

	// Check the value type and append windowTitle accordingly
	if val, exists := sliderMapping[index]; exists {
		switch v := val.(type) {
		case string:
			// Convert string to list of strings
			sliderMapping[index] = []string{v, windowTitle}
		case []interface{}:
			// Convert interface{} slice to string slice and append the new value
			newVal := make([]string, len(v))
			for i, elem := range v {
				newVal[i] = fmt.Sprint(elem)
			}
			sliderMapping[index] = append(newVal, windowTitle)
		case []string:
			// If it's already a list of strings, simply append the new value
			sliderMapping[index] = append(v, windowTitle)
		default:
			log.Fatalf("unexpected type for slider_mapping: %v", reflect.TypeOf(val))
		}
	} else {
		sliderMapping[index] = []string{windowTitle}
	}

	// Update the genericData map with the modified slider_mapping
	genericData["slider_mapping"] = sliderMapping

	// Marshal the modified genericData back to YAML
	modifiedData, err := yaml.Marshal(&genericData)
	if err != nil {
		log.Fatalf("error marshaling modified data: %v", err)
	}

	// Write the modified YAML back to the file
	err = ioutil.WriteFile("config.yaml", modifiedData, 0644)
	if err != nil {
		log.Fatalf("error writing file: %v", err)
	}

	// Print the modified configuration
	fmt.Println("Modified configuration:")
	fmt.Println(string(modifiedData))
}
