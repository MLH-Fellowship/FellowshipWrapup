import styled from "styled-components";

export const StyledBackground = styled.section`
  padding: 2rem;
  margin-bottom: 5.5rem;

  text-align: center;

  h1 {
    font-size: 2rem;
    max-width: 90%;
    margin: 0 auto 1.5rem auto;
  }

  h3 {
    margin-bottom: 0.5rem;
  }

  span {
    color: var(--color-main);
  }

  .text-container {
    max-width: 60rem;
    margin: 0 auto;
  }

  .progress-bars {
    display: grid;
    grid-template-columns: auto auto;
    grid-gap: 1rem;
    justify-content: center;
  }

  .progress-bar-item {
    background-color: var(--color-gray);
    border-radius: 5px;

    width: 28rem;

    padding: 1.5rem 3rem;
  }
`;
