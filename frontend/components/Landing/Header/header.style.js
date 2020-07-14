import styled from "styled-components";

export const StyledHeader = styled.header`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  min-height: 50vh;
  max-width: 95%;
  margin: 2rem auto;
  padding: 2.5rem;

  background-color: var(--color-gray);
  border-radius: 8px;

  text-align: center;

  .textContainer {
    max-width: 50rem;
    padding: 0 3rem;
    color: var(--color-text);

    img {
      height: 2rem;
      margin-bottom: 1rem;
    }

    h1 {
      font-size: 3rem;
      margin-bottom: 1.5rem;
    }

    h2 {
      font-size: 1.6rem;
      font-weight: 400;
      margin-bottom: 1.5rem;
    }

    button {
      background-color: var(--color-main);
      font-size: 0.9rem;
      font-weight: 600;
      font-family: Menlo, Monaco, Lucida Console, Liberation Mono,
        DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
      color: white;
      padding: 0.8rem 2.5rem;
      border: none;
      border-radius: 25rem;
      cursor: pointer;

      -webkit-touch-callout: none;
      -webkit-user-select: none;
      -khtml-user-select: none;
      -moz-user-select: none;
      -ms-user-select: none;
      user-select: none;
    }
  }
`;
