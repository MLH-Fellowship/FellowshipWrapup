const accountInfo = require("./data/accountInfo.json");

export const getUserInfo = () => {
  const stringify = JSON.stringify(accountInfo);

  return JSON.parse(stringify);
};

// export const testUserInfo = async () => {
//   const res = await axios({
//     method: "post",
//     url: `${process.env.BACKEND_URL}/accountinfo/sebastiancrossa`,
//     data: {
//       secret: process.env.BACKEND_SECRET,
//     },
//   }).then((res) => {
//     return res.data;
//   });
// };
