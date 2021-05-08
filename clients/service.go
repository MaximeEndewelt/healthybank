package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	GetAccountInformation = "https://api-sandbox.starlingbank.com/api/v2/accounts"
)

type service struct {
	httpClient *http.Client
}

func New() StarlingClient {
	return &service{
		httpClient: &http.Client{
			Timeout: 50 * time.Second,
		},
	}
}

func (s *service) GetAccountInformation() error {

	req, err := http.NewRequest("GET", GetAccountInformation, nil)
	if err != nil {
		log.Printf("Error initializing request: %s", err)
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJQUzI1NiIsInppcCI6IkdaSVAifQ.H4sIAAAAAAAAAH1Ty47cIBD8lZXP2yu_X7fc8gP5gDY0M2hssABPsory78EGj8eTVW6uqqapotu_E2lt0ic4S-A06Q_r0IxSXQZUtw-mp-Q9scvgKzqkCuumgrQbGigzXkPX5A0gttUwEEszUfhi-jUnfVbnadWWed29JxJdIMquaVcCGdOLct_1yMn8kNz3Toect0QFCF4glEXFoG3aCqjLWC1YkTVV43s7fSMVTmQ877KOEYiyq72bPIO2xRJKzuohZ2VZZIM_4WN9Y4ysDadyQV0lmgFEVggoK07Q4tBCU2KXp2mOKcc1MNMzrY8SnMJ1swoKJ-oNIX97Edzn_CJITspJIcmc-VFad2Ii4Nx4kz1x6R4gKM4hu070qDzwTyMdveHirtpI60cGUnF5l3zBMRQPOKJi0RpDw4Fp5Ywew0UrEzWthDQTOqkVaAFiUdw-JPu4fQfharZYp6c9Ik0oY-ORvBF16XGex88H2qomVBwd9ZxG8i12GDVzI7cGmQ0JMuS92_9JwUbQ5hH9Nkjl6GK2HM8H_xXjUTLsinu6iRx6N9gzDzc14i3UjJ9EuxRADBHAUQRywkvMFLTjE5xBZZEdDj0NwzLe-n2SdFDHbQEfFwa8N1j3we_WJN3Rc9TM78FTh40AvS7EKxtPGS3kuNsPeU7UVmWIkZzdCdizFB7X4t0PzMJFHz5OXExz4rY-z0x4MD_yr1oc4he9DjE0ZVfiy0gc4oNFmpzzAZc5whn3HyZ8xiwOtyUGbfiTlTO7ezizoZfz9duUmL2_UjMXkVoGy4x_wXU59m7P3Fb1vEHbkF5XKvnzF58NqELWBQAA.UAJRJt-qqG_ATqP9AEiz7Yiy0meygIilL1Gpl6nBKUxXBYuTCCRTLig_uUQJPhxbh7kgErCLlJk_3jcIWGjKPuISKNBkqPP2W5SVrDWCqslcmxptqBlwfaXhAspks3s2bEEZZWilGcxQldiuJKK7sgn2z5NzlcPHkk9cLENAwOMgRBfW4LbrPllA0CAAXznuFARzf31geMe546FucZ7Hst73wF9WMRATKu0cfXLyv6GFSzqlFbVBSZ8IHQW8ZQtY8C-3d8upbuKGcQtrIusjf58ixLhR2X3-v_It6WXg6KsQNwEoQYP4MVLT3zko0JQM833OSempqFLO7PzJIHvhvychinVWIzXXFDYs7J69AUUg1AYYZtIVaVYWJ5ZZI-zB73HlJ35bzTQwoOx0zAXFIvzZn2PR-1MyqsUFWiBA8y-v_7CQZpnU9MEed5GeZs5D24dCKVNF_jR5TZZbOiSXGctOPBLSFRqWzwpPwXBmgpxeMwAzUN_zm0rbdfAa4oFxrMHfxBY_lWZdoDt1zXQdTVzK8gFaFMF3ODaq5hQX-zrtg0CSe84EeCueNseQ9PjpeebFpW8NosaCyTW9aJ0NpxebpoXtfVzfkyUuZn4BEfZCPIVKUae9EifMwjqNEfbrz4RFpWT4Oo45cFLnW_qH8CHNxHjY9wrWtsshOHXpl2I")
	req.Header.Add("accept", "application/json")

	res, err := s.httpClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	if res.StatusCode > 400 {
		log.Printf("Bad request : %s", res.Body)
	} else {
		log.Printf("Good request : %s", res.Status)
	}
	target := &AccountInformation{}
	if err = json.NewDecoder(res.Body).Decode(target); err != nil {
		log.Printf("Decode fail %s", err)
	}

	fmt.Println(*target)

	defer res.Body.Close()
	return nil
}
