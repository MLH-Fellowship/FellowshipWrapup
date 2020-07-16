const accountInfo = require("./data/accountInfo.json");

// TODO: change it so it gets the actual user data from the server
export function getUserInfo() {
  const stringify = JSON.stringify(accountInfo);

  return JSON.parse(stringify);
}
