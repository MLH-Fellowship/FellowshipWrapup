const fs = require("fs");
import path from "path";

const infoDir = path.join("./data");

// TODO: change it so it gets the actual user data from the server
export function getUserInfo() {
  fs.readFileSync("./data/accountInfo.json", function (err, data) {
    // Check for errors
    if (err) throw err;

    // Converting to JSON
    const users = JSON.parse(data);

    console.log(users); // Print users
  });
}
