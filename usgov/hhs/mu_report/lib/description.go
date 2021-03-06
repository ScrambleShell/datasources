package lib

import (
  "os"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.mu_report",
      Version: "2015-00-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://www.healthit.gov/sites/default/files/mu_report.csv")
  if err != nil {
    return nil, err
  }

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer fileReader.Close()

  csvReader := csv.NewReader(fileReader)
  if err != nil {
    return nil, err
  }

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch("http://www.healthit.gov/sites/default/files/mu_report.csv")
  if err != nil {
    return nil, err
  }

  fileReader, err := os.Open(path)
  if err != nil {
    return nil, err
  }

  csvReader := helpers.NewCsvReader(fileReader)

  return csvReader, nil
}