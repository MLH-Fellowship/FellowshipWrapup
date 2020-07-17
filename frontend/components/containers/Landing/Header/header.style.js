import styled from "styled-components";

export const StyledHeader = styled.header`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  min-height: 90vh;
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
  }

  .inputContainer {
    display: flex;
    align-items: center;
    max-width: 25rem;
    margin: 0 auto;
    padding-right: 1.5rem;

    border-radius: 5px;

    background-color: white;

    input {
      border: none;
      border-radius: 5px 0 0 5px;
      padding: 1.3rem 1rem;
      width: 100%;

      outline: none;

      font-family: Menlo, Monaco, Lucida Console, Liberation Mono,
        DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
    }

    button {
      height: 2rem;
      width: 6rem;
      padding: 0 1rem;
      border-radius: 5px;
      font-weight: 600;

      cursor: pointer;

      background-color: var(--color-main);
      border: none;
      color: white;
    }
  }
`;
