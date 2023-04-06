package ipfs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IEatLemons/goUtils/request"
)

const (
	APIV0FILES = "/api/v0/files"
)

type Files struct{}

type ResultLS struct {
	Entries []*V0Info `json:"Entries"`
}

type V0Info struct {
	Name           string `json:"Name"`
	Type           int    `json:"Type"`
	Size           int    `json:"Size"`
	Hash           string `json:"Hash"`
	CumulativeSize int    `json:"CumulativeSize"`
	Blocks         int    `json:"Blocks"`
}

func (IPFS *IPFSV0) LS(path string) (List *ResultLS, err error) {
	fmt.Println("[IPFS LS]", path)
	url := IPFS.Url + APIV0FILES + "/ls?arg=" + path
	result, err := request.NewRequest(url).BasicAuth(IPFS.BasicAuth).Post(nil).Send()
	if err != nil {
		return
	}
	List = &ResultLS{}
	err = json.Unmarshal(result, List)
	return
}

func (IPFS *IPFSV0) Stat(path string) (Info *V0Info, err error) {
	fmt.Println("[IPFS LS]", path)
	url := IPFS.Url + APIV0FILES + "/stat?arg=" + path
	result, err := request.NewRequest(url).BasicAuth(IPFS.BasicAuth).Post(nil).Send()
	if err != nil {
		return
	}
	Info = &V0Info{}
	err = json.Unmarshal(result, Info)
	return
}

func (IPFS *IPFSV0) CP(source, dest string) {
	fmt.Println("[IPFS CP]", source, "to", dest)
	url := IPFS.Url + APIV0FILES + "/cp?arg=" + source + "&arg=" + dest

	result, err := request.NewRequest(url).BasicAuth(IPFS.BasicAuth).Post(nil).Send()
	if err != nil {
		log.Println("[IPFS CP]", err)
	} else {
		fmt.Println("[IPFS CP]", "done", string(result))
	}
}

func (IPFS *IPFSV0) RM(name string) {
	fmt.Println("[IPFS RM]", name)
	url := IPFS.Url + APIV0FILES + "/rm?arg=" + name

	result, err := request.NewRequest(url).BasicAuth(IPFS.BasicAuth).Post(nil).Send()
	if err != nil {
		log.Println("[IPFS RM]", err)
	} else {
		fmt.Println("[IPFS RM]", "done", string(result))
	}
}

func (IPFS *IPFSV0) ADD(Payload *request.Payload, path string) {
	fmt.Println("[IPFS ADD]", path)
	url := IPFS.Url + APIV0FILES + "/write?create=true&arg=" + path
	result, err := request.NewRequest(url, []request.RequestOptions{
		request.SetContentType(request.MulFormData),
	}...).BasicAuth(IPFS.BasicAuth).File(Payload).Send()
	if err != nil {
		log.Println("[IPFS ADD]", err)
	} else {
		fmt.Println("[IPFS ADD]", "done", string(result))
	}
}

func (IPFS *IPFSV0) READ(path string) {
	fmt.Println("[IPFS FLUSH]", path)
	url := IPFS.Url + APIV0FILES + "/read?arg=" + path
	result, err := request.NewRequest(url).Post(nil).BasicAuth(IPFS.BasicAuth).Send()
	if err != nil {
		log.Println("[IPFS FLUSH]", err)
	} else {
		fmt.Println("[IPFS FLUSH]", "done", string(result))
	}
}

func (IPFS *IPFSV0) FLUSH(path string) {
	fmt.Println("[IPFS FLUSH]", path)
	url := IPFS.Url + APIV0FILES + "/flush?arg=" + path
	result, err := request.NewRequest(url).Post(nil).BasicAuth(IPFS.BasicAuth).Send()
	if err != nil {
		log.Println("[IPFS FLUSH]", err)
	} else {
		fmt.Println("[IPFS FLUSH]", "done", string(result))
	}
}
