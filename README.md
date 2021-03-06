<center>

# Fellowship Wrapped

A summary showcase of all your time spent during the MLH Fellowship.

![fellowshipwrapped](https://user-images.githubusercontent.com/48270786/87963701-c1b68d80-cad6-11ea-9dc1-231675961677.gif)

#### Create a shareable link to display your contributions during the Fellowship on your resume or your website!

</center>


## Features

- A personalized page created for every fellow.

- Total lines of code written, commits, and issues participated in.

- Languages used in the respective fellow's projects

- Your performance ratio, like how many PRs were merged compared to total PRs opened by the fellow.

- and other awesome statistics!

## Technologies Used

- [Next.js](https://nextjs.org/)

- [Go Lang](https://golang.org/)

- [Bootstrap 4](https://getbootstrap.com/)

- [Github API](https://docs.github.com/en/graphql)

- [MLH-Fellow-Map](https://github.com/Korusuke/MLH-Fellow-Map)

## Built originally by

<table>
  <tr>
    <td align="center"><a href="https://gmcruz.me/"><img src="https://avatars1.githubusercontent.com/u/8129788?s=400&u=93725230cac5a1e8eb03f65e548e59d4cd14d70a&v=4" width="100px;" alt=""/><br /><sub><b>
Gabriel Cruz</b></sub></a></td>
    <td align="center"><a href="https://kartikcho.github.io"><img src="https://avatars1.githubusercontent.com/u/48270786?v=4" width="100px;" alt=""/><br /><sub><b>Kartik Choudhary</b></sub></a></td>
    <td align="center"><a href="https://iamcathal.github.io/"><img src="https://avatars0.githubusercontent.com/u/6561327?s=400&u=3746478b26e66ebe22eba9ba20097b477c455cc3&v=4" width="100px;" alt=""/><br /><sub><b>Cathal O'Callaghan</b></sub></a></td>
    <td align="center"><a href="http://sebastiancrossa.com/"><img src="https://avatars2.githubusercontent.com/u/20131547?s=400&u=90fe733b6d501490b786b039f3f9e9e19da042c2&v=4" width="100px;" alt=""/><br /><sub><b>Sebastian Crossa</b></sub></a></td>
  </tr>
</table>

## Installation

#### Frontend

To run the frontend locally, run the following commands in order
(Assumed you have NPM or Yarn installed on your machine)

`yarn` or `npm install`

`yarn dev` or `npm run dev` 


#### Backend

* To run the backend you will need to export your [GitHub personal access](https://docs.github.com/en/github/authenticating-to-github/creating-a-personal-access-token) token and your secret to be used in ensuring that bots aren't pinging your server and using up your credits. To do so export them like so:

`export GRAPHQL_TOKEN=[your github personal access token]`

`export secretKey=[your secret key]`

Then run the server with `cd backend/ && go run main.go`
