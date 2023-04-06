package ipfs

import (
	"fmt"
	"log"

	"github.com/IEatLemons/goUtils/request"
)

type IPFSV0 struct {
	Url       string
	BasicAuth string
	Files     *Files
}

func (IPFS *IPFSV0) CAT(path string) {
	fmt.Println("[IPFS CAT]", path)
	url := IPFS.Url + "/api/v0/cat?arg=" + path
	result, err := request.NewRequest(url).Post(nil).BasicAuth(IPFS.BasicAuth).Send()
	if err != nil {
		log.Println("[IPFS ADD]", err)
	} else {
		fmt.Println("[IPFS ADD]", "done", string(result))
	}
}
