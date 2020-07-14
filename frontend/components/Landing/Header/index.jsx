import React from "react";
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
          It's time to show off your hard work from the past 3 months to the
          world
        </h2>
        <button>Sign in with Github</button>
      </div>
    </StyledHeader>
  );
};

export default Header;
