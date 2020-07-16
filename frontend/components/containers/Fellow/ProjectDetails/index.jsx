import { ProgressBar } from "react-bootstrap";

import { StyledBackground } from "./projectDetails.style";

export default function ProjectDetails() {
  return (
    <StyledBackground>
      <div className="text-container">
        <h1>
          During the past 3 months I've contributed to projects like{" "}
          <span>[project]</span> &<span>[project]</span>
        </h1>
      </div>
    </StyledBackground>
  );
}
