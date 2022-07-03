package security

import (
	"context"
	"encoding/json"
	"fiberapiv1/configs"
	"fiberapiv1/errs"
	"fiberapiv1/helper"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	//"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/api/idtoken"

	"github.com/MicahParks/keyfunc"
)

func getSecretKey() string {
	secretKey := configs.GetLineClientSecret() //or GetSecretKey
	if secretKey == "" {
		secretKey = "secretKey"
	}
	return secretKey
}

//use in service level
func GenerateToken(UserID string) (*string, error) {
	claims := jwt.StandardClaims{
		Issuer:    UserID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //time.Now().Add(time.Hour * 24).Unix() 1 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(getSecretKey()), nil
}

//validate v1
func validateStandardClaimsToken(tokenString string) (*jwt.StandardClaims, error) {
	claims := new(jwt.StandardClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, validateSignedMethod)
	if err != nil {
		return nil, err
	}
	var ok bool
	claims, ok = token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errs.NewInvalidAuthTokenError()
	}
	return claims, nil
}

func validateHMACToken(tokenHs string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenHs, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecretKey()), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errs.NewInvalidAuthTokenError()
	}
	return claims, nil
}
func validateEsToken(tokenES string) (map[string]interface{}, error) {
	//Get public key using kid link https://api.line.me/oauth2/v2.1/certs

	//step 1  Open our line_certs File
	certsFile, err := os.Open("security/line_certs.json")
	if err != nil {
		fmt.Println(err)
	}
	defer certsFile.Close()
	byteCerts, _ := ioutil.ReadAll(certsFile)
	var jwksJSON json.RawMessage = []byte(byteCerts)

	//step 2 Get the JWKS as JSON.
	jwks, err := keyfunc.NewJSON(jwksJSON)
	if err != nil {
		return nil, errs.NewError("Failed to create JWKS from resource at the given URL Error : " + fmt.Sprint(err.Error()))
	}

	//step 3 parse valid token
	token, err := jwt.Parse(tokenES, jwks.Keyfunc)
	claims, _ := token.Claims.(jwt.MapClaims)
	if err != nil {
		return claims, errs.NewError("Failed to parse the JWT Error : " + fmt.Sprint(err.Error()))
	}

	//step 4 Check if the token is valid.
	if !token.Valid { //The token is not valid.
		return claims, err
	}
	//The token is valid.
	return claims, nil
}

func validateDo(tokenString, alg string) (map[string]interface{}, error) {
	switch alg {
	case "ES256":
		return validateEsToken(tokenString)
	case "HS256":
		return validateHMACToken(tokenString)
	default:
		return nil, errs.NewError("func is not alg " + alg)
	}
}

//validate v2
func validateMapClaimsToken(tokenString string) (map[string]interface{}, error) {
	//step 1 decode header
	parts := strings.Split(tokenString, ".")
	jsonHeader, err := jwt.DecodeSegment(parts[0])
	if err != nil {
		panic(err)
	}
	var header map[string]interface{}
	err = json.Unmarshal([]byte(jsonHeader), &header)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded alg: %s\n", header["alg"])

	//step 2 validate func by type
	alg := header["alg"]
	claims, err := validateDo(tokenString, fmt.Sprint(alg))

	//step 3 format time ExpiresAt and show err
	var tm time.Time
	switch exp := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(exp), 0)
	case json.Number:
		v, _ := exp.Int64()
		tm = time.Unix(v, 0)
	}
	fmt.Println("claims : ", claims)
	fmt.Println("ExpiresAt : ", tm)

	if err != nil {
		fmt.Println("token err = ", err)
		return nil, err
	}
	return nil, nil
}

func oAuthLineRequest(c *fiber.Ctx) (interface{}, error) { //case get header idToken line provider
	headers := c.GetReqHeaders()
	token := headers["Authorization"]
	_, err := validateMapClaimsToken(token)
	if err != nil {
		return nil, errs.NewUnauthorizedError()
	}
	return nil, nil
}

func oAuthRequest(c *fiber.Ctx) error { //case get header idToken google provider
	headers := c.GetReqHeaders()
	token := headers["Authorization"]
	payload, err := idtoken.Validate(context.Background(), token, configs.GetGoogleClientID())
	if err != nil {
		return errs.NewUnauthorizedError()
	}
	email := payload.Claims["email"]
	fmt.Println("Email : ", email)
	return nil
}

func authRequest(c *fiber.Ctx) (*jwt.StandardClaims, error) { //case get header jwt token
	headers := c.GetReqHeaders()
	token := headers["Authorization"]
	claims, err := validateStandardClaimsToken(token)
	if err != nil {
		return nil, errs.NewUnauthorizedError()
	}
	return claims, nil
}
func authCookieRequest(c *fiber.Ctx) (*jwt.StandardClaims, error) { //case get cookie jwt token
	cookieName := helper.CookieTokenName
	cookie := c.Cookies(cookieName)
	claims, err := validateStandardClaimsToken(cookie)
	if err != nil {
		return nil, errs.NewUnauthorizedError()
	}
	return claims, nil
}

//use middleware in routes level
func AuthorizationRequired() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		_, err := oAuthLineRequest(c) //authRequest, authCookieRequest or oAuthRequest or oAuthLineRequest
		if err != nil {
			response := helper.BuildErrorResponse("Unauthorized!", err, nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}
		return c.Next()
	}
}
