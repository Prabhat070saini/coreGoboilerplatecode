package fileservice

type UploadFileResponse struct {
	Key string `json:"key"`
	URL string `json:"url"`
}

type FetchFileResponse struct {
	URL string `json:"url"`
}
