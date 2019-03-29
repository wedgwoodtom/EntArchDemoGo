package main

import (
	"math"
	"math/rand"
	"strconv"
)

type MediaNotification struct {
	Id           string `json:"id"`
	Notification string `json:"notification"`
}

func MakeRandomMedia() MediaNotification {
	id := rand.Intn(math.MaxUint32)

	message := `{
        "$xmlns": {
            "media": "http://search.yahoo.com/mrss/",
            "pla": "http://xml.theplatform.com/data/object/admin",
            "plmedia": "http://xml.theplatform.com/media/data/Media"
        },
        "id": 6596199989,
        "method": "put",
        "type": "Media",
        "cid": "e157343f-c115-4713-baaa-e9e5493cba61",
        "entry": {
            "id": "http://data.media.theplatform.com/media/data/Media/1108606019587",
            "updated": 1512363274000,
            "ownerId": "http://access.auth.theplatform.com/data/Account/2649321885",
            "updatedByUserId": "https://identity.auth.theplatform.com/idm/data/User/mpx/2744660",
            "guid": "AMC_TWD_807_DAI",
            "title": "Time for After",
            "pla$adminTags": [
                "LongformIngested"
            ],
            "plmedia$approved": true,
            "plmedia$originalOwnerIds": [
                
            ],
            "plmedia$originalMediaIds": [
                
            ],
            "plmedia$programId": "",
            "plmedia$seriesId": "",
            "plmedia$availabilityState": "available"
        }
    }
`

	return MediaNotification{Id: strconv.Itoa(id), Notification: message}
}
