package server

import "chatApp/encryptdecrypt"

var password string = "C1C3C89A4E0DF49756259D5F3EAFB77D"
var authToken string = "90C3884100C373AE1E00EB7BECA1FDC3E4901F606D548CE7CFC40D4724BF5E28160E12C9582C689BD348D251E5432058058E4C450ABEF4F8AB7730EC29C7F8CF"
var delimiter string = "\n"
var finishedFalse string = "finished:false" + delimiter
var finishedTrue string = "finished:true" + delimiter
var authTokenStr string = "authToken:" + authToken + delimiter
var successResponse string = "wasSuccessful:true" + delimiter
var notSuccessResponse string = "wasSuccessful:false" + delimiter

//EndOfLife is the variable to determine when the server should be closed
var EndOfLife bool = false

//Encrypt encrypts a string
func Encrypt(text string) string {
	encrypted := encryptdecrypt.Encrypt(text, password)
	return encrypted
}

//Decrypt decrypts a string
func Decrypt(text string) string {
	decrypted := encryptdecrypt.Decrypt(text, password)
	return decrypted
}
