import React, { useState } from "react";
import Router from "next/router";
import { StyledHeader } from "./header.style";

const Header = () => {
  const [username, setUsername] = useState("");

  const handleInputChange = (e) => {
    const { name, value } = e.target;

    setUsername(value);
  };

  return (
    <StyledHeader>
      <div className="textContainer">
        <img src="/mlh-dark.svg" alt="MLH Logo" className="logo" />
        <h1>
          <span style={{ color: "var(--color-main)" }}>Congratulations</span>,
          graduating fellows of Class 0
        </h1>
        <h2>
          It's time to show off your hard work to the world. Enter your Github
          username and start generating your own, personal Fellowship portfolio
        </h2>

        <div className="inputContainer">
          <input
            type="text"
            onChange={(e) => handleInputChange(e)}
            value={username}
            placeholder="Username..."
          />
          <button onClick={() => Router.push(`/fellow/${username}`)}>GO</button>
        </div>
      </div>
    </StyledHeader>
  );
};

export default Header;
