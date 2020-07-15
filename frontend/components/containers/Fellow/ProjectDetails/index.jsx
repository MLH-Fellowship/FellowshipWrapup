import { ProgressBar } from "react-bootstrap";

import { StyledBackground } from "./projectDetails.style";

export default function ProjectDetails() {
  return (
    <StyledBackground>
      <h1>
        During the past 3 months I've contributed to projects like{" "}
        <span>[project]</span> &<span>[project]</span>. My most used languages
        were
      </h1>

      <div
        className="progress-bars"
        style={{ display: "grid", gridTemplateColumns: "auto auto" }}
      >
        <div className="progress-bar-item">
          <h3>JavaScript</h3>
          <ProgressBar now={60} label={"Used in 60% of commits"} />
        </div>

        <div className="progress-bar-item">
          <h3>TypeScript</h3>
          <ProgressBar now={60} label={"Used in 60% of commits"} />
        </div>
      </div>
    </StyledBackground>
  );
}
