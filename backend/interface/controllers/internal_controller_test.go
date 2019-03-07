package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
)

type MockFirebaseInternalRepository struct {
}

func (mockFirebase *MockFirebaseInternalRepository) FindPublicKey(client *http.Client) (*http.Response, error) {
	publicKey := publicKey{
		First:  "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIIGUK+PrXyps0wDQYJKoZIhvcNAQEFBQAwMTEvMC0GA1UE\nAxMmc2VjdXJldG9rZW4uc3lzdGVtLmdzZXJ2aWNlYWNjb3VudC5jb20wHhcNMTkw\nMjEzMjEyMDQ4WhcNMTkwMzAyMDkzNTQ4WjAxMS8wLQYDVQQDEyZzZWN1cmV0b2tl\nbi5zeXN0ZW0uZ3NlcnZpY2VhY2NvdW50LmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAL5vdrmmwDJs3c4NndKnGO3Dj/ZF90bhWvh77Rs8SrGEKyBf\nsGotzJUB4YlolYUL4umcaAPOLV9HNVRFoAmE8C+rtMQK5guwLsDFRT08DPIuyPRf\n3VlcZSXXUeS2dl6e/SymXT7RUeArdSnMvQ786wi5lx/waiaFUH1HVLLq8taWfhy1\nDrX9aR7GzMn7UUkYgF/I4zSZNwGChckap3JL9YlyN52TxnFRCamT65OMrtpApsxJ\nme8X+717/9hAo/DFxs1HbwW3T/vIQfdp4DWsIzgUzEI2R1tn00CLpy+t1zzM/4W5\nuxbxWAoJ3PTyZAExULOKKBBQ/XYQoo5qNCgwa0MCAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAOBgNVHQ8BAf8EBAMCB4AwFgYDVR0lAQH/BAwwCgYIKwYBBQUHAwIwDQYJ\nKoZIhvcNAQEFBQADggEBADRimV8iZJnK47d9j6Ssqq7wuJsb4K/X8pEM+wWtB9/b\nxjjffN8j5iCNtnz/PSmdRC9O80Mmi7F34SZRaKiMdDCx94UIavnyR1VqMbFirlvI\nIbZ4Xn3JxIP9anvK1hYuOvKbjhEP8dVRLNqxrUb1vpHxdrDN1cH7A0ZqxvLcPC4k\nkJjGzpIq7sYan9Lcn/qKNJijr1BZeiZWVzcu/j8IBdzCajrR0HVtpDZwB6YMu32L\n6SZvnDwmKs6Ycr+TIncl012jaEtrvQ5bT/YwKYajH0tYBkX3FwTz95vOHOKg3zSr\nePTkOFdR+JGTetQN/smlNsHuULN4TY5O7rypSxnbYpY=\n-----END CERTIFICATE-----\n",
		Second: "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIIKtPg86XjbF8wDQYJKoZIhvcNAQEFBQAwMTEvMC0GA1UE\nAxMmc2VjdXJldG9rZW4uc3lzdGVtLmdzZXJ2aWNlYWNjb3VudC5jb20wHhcNMTkw\nMjIxMjEyMDQ4WhcNMTkwMzEwMDkzNTQ4WjAxMS8wLQYDVQQDEyZzZWN1cmV0b2tl\nbi5zeXN0ZW0uZ3NlcnZpY2VhY2NvdW50LmNvbTCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAMkVZElCuClxdc6tFZuGSKZ3/8/AqpydCqXVknvhQAhC9TMQ\nhBecKKwCitoNahV24Ri7AJ0HAwkJuzm88RhS2ZsP3kt376VGnJAe30+5HxSu52Gg\n/PurVO9OCZMwZ1op6sBbEmgYQuS8bGJKXScgy3yN1C8MDtvDaO0N7XLu3/TUSjU5\nNWUt2dx/xd+q4PPOgBh6AuIcEj5nmk1QB5c9ir/RffOkNXOfxVdU5YnSLy3qDS0s\n3Wl1FO3VDvd6IlFi/Wy2vA8Imjth4bxmlH2u+QjIWwulUZA00onlO2IBLAgjCaPO\nvtTVRzp2zV++PZn2UHLjOGLrZ2gjhOXfEYQXlVkCAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAOBgNVHQ8BAf8EBAMCB4AwFgYDVR0lAQH/BAwwCgYIKwYBBQUHAwIwDQYJ\nKoZIhvcNAQEFBQADggEBAJa+HefrXAbpEi0amX1xMOjMFNBlaZLW+izxUHP1ayUC\ndFNIHBjCnR9Iy3zyDhblkotf2+9YEoTQcRFVBypBRrYsXPIPZLwZWaGSF8j2lQim\nacDyfBBbYJoBx7Agj3200YYXQZz9O+riDWfCkaOgD36iQRiuN/8ZGjPqzF81mDmG\ncyq9J1FEUFK7MUVt+hgjq73oEyTFMMLB7mqrxPKjg9tNms0t5X+9eXeG8LzySsC+\nMc0ACYnq6DzPU7IKDadRaNmHk/2UQO4BkYX2AOxTuwbAhAP7eVAPq49ImTcCDhmj\nqQUY3N/4hgkZV+On43NmPRigu/eJeWidEqBelX+CBhc=\n-----END CERTIFICATE-----\n",
	}

	body, err := json.Marshal(publicKey)
	if err != nil {
		return nil, err
	}

	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader(body)),
	}
	return res, nil
}

func (mockFirebase *MockFirebaseInternalRepository) ParseToken(idToken string, keys map[string]*json.RawMessage) (*jwt.Token, error) {
	return &jwt.Token{}, nil
}

func (mockFirebase *MockFirebaseInternalRepository) FindSub(parsedToken *jwt.Token) (string, bool) {
	return "matilda12345", true
}
