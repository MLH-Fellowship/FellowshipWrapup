import { StyledHeader, Grid } from "./header.style";
import GithubProfile from "../GithubProfile";

export default function Header() {
  return (
    <StyledHeader>
      <Grid>
        <div className="textContainer">
          <h1>
            Hi, I'm <span>Sebastian Crossa</span>
          </h1>
          <h2>
            I'm a student from <span>Zapopan, Jalisco</span> and I'm a part of
            the{" "}
            <a
              href="https://fellowship.mlh.io/"
              target="_blank"
              rel="noopener noreferrer"
            >
              inaugural class of MLH Fellows of 2020
            </a>
            , where I worked alongside other fellows under the educational
            mentorship of professional software engineers.
          </h2>
        </div>
        <div style={{ margin: "0 auto" }}>
          <img src="/profile.jpg" alt="User profile pic" />
          <GithubProfile />
        </div>
      </Grid>
    </StyledHeader>
  );
}
