package authenticators

//Tests:
// token is not base64 encoded
// token does not contain a colon
// token does not exist in file
// file does not exist so cannot monitor
// file changed, pick up changes
// can successfully authenticate if hash is encoded correctly and exists in string.
