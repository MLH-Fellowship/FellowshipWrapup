import styled from "styled-components";

export const StyledBackground = styled.section`
  padding: 2rem;
  margin-bottom: 7.5rem;

  text-align: center;

  h1 {
    font-size: 2rem;
    max-width: 90%;
    margin: 0 auto;
  }

  h3 {
    margin-bottom: 0.5rem;
  }

  span {
    color: var(--color-main);
  }

  .progress-bars {
    display: grid;
    grid-template-columns: auto auto;
  }

  .progress-bar-item {
    padding: 1rem 3rem;
  }
`;
