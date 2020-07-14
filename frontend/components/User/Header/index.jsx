import { StyledBackground } from "./header.style";

export default function Header() {
  return (
    <StyledBackground>
      <div className="textContainer">
        <img src="/profile.jpg" alt="[user] profile image" />
        <h1>
          Hi, I'm <span>Sebastian Crossa</span> and I'm a graduating student
          from the MLH Fellowship program. I contributed to{" "}
          <span className="underline">Amplify CLI</span> and{" "}
          <span className="underline">Amplify Docs</span>
        </h1>
      </div>
    </StyledBackground>
  );
}
