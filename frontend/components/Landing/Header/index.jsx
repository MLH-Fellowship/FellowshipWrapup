import React from "react";
import Router from "next/router";
import { StyledHeader } from "./header.style";

const Header = () => {
  return (
    <StyledHeader>
      <div className="textContainer">
        <img src="/mlh-dark.svg" alt="MLH Logo" className="logo" />
        <h1>
          <span style={{ color: "var(--color-main)" }}>Congratulations</span>,
          graduating fellows of Class 0
        </h1>
        <h2>
          It's time to show off your hard work to the world. Generate your own
          personal link by signing in with your Github account and start sharing
          with others.
        </h2>
        <button onClick={() => Router.push("/user/1")}>
          Sign in with Github
        </button>
      </div>
    </StyledHeader>
  );
};

export default Header;
