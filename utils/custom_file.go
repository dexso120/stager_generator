package utils

import (
    "path/filepath"
    "archive/zip"
    "io"
    "os"
    "log"
    "fmt"

)

func Zipfiles(zipFilePath string, files ...string){
    archive, err := os.Create(zipFilePath)
    if err != nil {
        log.Fatal(err)
    }

    defer archive.Close()
    zipWriter := zip.NewWriter(archive)

    for _, file := range files{
        f1, err := os.Open(file)
        if err != nil{
            log.Fatal(err)
        }

        defer f1.Close()

        baseFilename := filepath.Base(file)

        w1, err := zipWriter.Create(baseFilename)
        if err != nil{
            log.Fatal(err)
        }

        if _, err := io.Copy(w1, f1); err != nil {
            panic(err)
        }
    }

    zipWriter.Close()
}

func CleanUpTempOutfiles() error {
    dir := filepath.Join("./", "temp_outfile")
    entries, err := os.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("failed to read directory %s: %w", dir, err)
    }

    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }

        filePath := filepath.Join(dir, entry.Name())
        if err := os.Remove(filePath); err != nil {
            return fmt.Errorf("failed to delete file %s: %w", filePath, err)
        }
    }

    return nil
}

func CleanUpOutfiles() error {
    dir := filepath.Join("./", "outfile")
    entries, err := os.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("failed to read directory %s: %w", dir, err)
    }

    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }

        filePath := filepath.Join(dir, entry.Name())
        if err := os.Remove(filePath); err != nil {
            return fmt.Errorf("failed to delete file %s: %w", filePath, err)
        }
    }

    return nil
}

func EnsureDir(path string) error {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        if err := os.MkdirAll(path, 0755); err != nil {
            return fmt.Errorf("failed to create directory %s: %w", path, err)
        }
        return nil
    }
    if err != nil {
        return fmt.Errorf("failed to check directory %s: %w", path, err)
    }
    return nil
}