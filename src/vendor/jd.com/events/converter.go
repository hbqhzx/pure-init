package events

import (
	"encoding/json"
)

type Converter interface {
	ToEvent(any interface{}) Event
}

type Convert struct {
	ArkServiceTreeDict map[string]string
	DataWhiteList      map[string]struct{}
}

func NewConvert(dict map[string]string, whitelist []string) Convert {
	whiteSet := make(map[string]struct{})
	for _, word := range whitelist {
		if _, ok := whiteSet[word]; !ok {
			whiteSet[word] = struct{}{}
		}
	}

	return Convert{dict, whiteSet}
}

func (c Convert) Convert(any interface{}) (*Event, error) {
	anyMap := make(map[string]interface{})
	bytes, err := json.Marshal(any)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &anyMap)
	if err != nil {
		return nil, err
	}

	// move field in whitelist to "data" field of Event
	data := make(map[string]interface{})
	for field, _ := range c.DataWhiteList {
		if value, ok := anyMap[field]; ok {
			data[field] = value
			delete(anyMap, field)
		}
	}
	anyMap["data"] = data

	// rename field name of ark service tree node info
	for field, value := range anyMap {
		if newField, ok := c.ArkServiceTreeDict[field]; ok {
			anyMap[newField] = value
			delete(anyMap, field)
		}
	}

	// TODO: cleanup
	//fmt.Printf("%v", anyMap)
	bytes, err = json.Marshal(anyMap)
	if err != nil {
		return nil, err
	}
	var event Event
	err = json.Unmarshal(bytes, &event)
	//fmt.Println(string(bytes))
	return &event, nil
}
