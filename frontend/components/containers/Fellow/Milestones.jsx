import styled from "styled-components";
import Countup from "react-countup";

export default function Milestones() {
  return (
    <StyledBackground>
      <div className="text-container">
        <h1>Some of the milestones I hit during the fellowship</h1>
      </div>

      <div className="list">
        <div className="list-item">
          <span>
            <Countup end={30} duration={2} />
          </span>{" "}
          <br />
          lines of code
        </div>
        <div className="list-item">
          <span>
            <Countup end={30} duration={3} />
          </span>{" "}
          <br /> commits made
        </div>
        <div className="list-item">
          <span>
            <Countup end={30} duration={4} />
          </span>{" "}
          <br />
          issues participated
        </div>
      </div>
    </StyledBackground>
  );
}

const StyledBackground = styled.section`
  text-align: center;

  padding: 2rem;

  h1 {
    font-size: 2rem;
    margin-bottom: 1.5rem;
  }

  .text-container {
    max-width: 60rem;
    margin: 0 auto;
  }

  .list {
    display: flex;
    justify-content: center;
  }

  .list-item {
    color: var(--color-text);

    padding: 2.5rem 1.5rem 2rem 1.5rem;
    border-radius: 8px;
    width: 14rem;

    -webkit-box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 0.45);
    -moz-box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 0.45);
    box-shadow: 0px 1px 80px -16px rgba(130, 130, 130, 0.45);

    &:nth-child(2) {
      margin: 0 2rem;
    }

    span {
      font-size: 3.5rem;
      font-weight: 700;
      color: var(--color-secondary);
    }
  }
`;
