package cv

import (
    "os"
    "io"
    "encoding/json"
    "github.com/go-yaml/yaml"
)

type CV struct {
    Name string `json:"name"`
    Title string `json:"title"`
    Contact Contact `json:"contact"`
    Skills []string `json:"skills"`
    Languages map[string]string `json:"languages"`
    Educations []Education `json:"educations"`
    About string `json:"about"`
    Experiences []Experience `json:"experiences"`  
}

type Contact struct {
    Phone string `json:"phone"`
    Email string `json:"email"`
    Website string `json:"website"`
    Location string `json:"location"`
}

type Education struct {
    Title string `json:"title"`
    Specialty string `json:"specialty"`
    Place string `json:"place"`
    StartYear uint `json:"start_year"`
    EndYear uint `json:"end_year"`
}

type Experience struct {
    Title string `json:"title"`
    Description string `json:"description"`
    StartYear uint `json:"start_year"`
    EndYear uint `json:"end_year"`
}

func ParseFromJson(jsonFilePath string) (CV, error) {
    var cv CV

    jsonFile, err := os.Open(jsonFilePath)
    if err != nil {
        return cv, err
    }
    defer jsonFile.Close()

    jsonParser := json.NewDecoder(jsonFile)
    if err = jsonParser.Decode(&cv); err != nil {
        return cv, err
    }

    return cv, nil
}

func ParseFromYaml(yamlFilePath string) (CV, error) {
    var cv CV

    yamlFile, err := os.Open(yamlFilePath)
    if err != nil {
        return cv, err
    }
    defer yamlFile.Close()

    data, err := io.ReadAll(yamlFile)
    if err != nil {
        return cv, err
    }
    err = yaml.Unmarshal([]byte(data), &cv)
    if err != nil {
        return cv, err
    }

    return cv, nil
}
