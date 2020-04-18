package download

import (
    "io"
    "net/http"
    "os"
)


// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func File(url string, filepath string) error {


	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla 5.0")
    // Get the data
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()



    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    return err
}
