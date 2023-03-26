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
	result, err := request.NewRequest(url).Post(nil).Send()
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
	result, err := request.NewRequest(url).Post(nil).Send()
	if err != nil {
		return
	}
	Info = &V0Info{}
	err = json.Unmarshal(result, Info)
	return
}

func (IPFS *IPFSV0) RM(name string) {
	fmt.Println("[IPFS RM]", name)
	url := IPFS.Url + APIV0FILES + "/rm?arg=" + name

	result, err := request.NewRequest(url).Post(nil).Send()
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
	}...).File(Payload).Send()
	if err != nil {
		log.Println("[IPFS ADD]", err)
	} else {
		fmt.Println("[IPFS ADD]", "done", string(result))
	}
}
